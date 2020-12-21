# 

[TOC]

# Crawler 技术总结

## 单任务版爬虫

### 总体算法&解析



![image-20201216195546138](https://i.loli.net/2020/12/16/LetE4G7NFgwR5kP.png)

- 总体算法类似**广度优先**算法走迷宫机制

- 解析部分为三个部分（城市列表、城市、用户），三种架构通用

  - 解析的结果集

    ```go
    type ParseResult struct {
    	Requests []Request
    	Items    []Item
    }
    ```

  - 起点（种子链接）：本项目的种子链接为珍爱网的城市列表（**城市列表解析器的目的是找城市**）

    - 城市列表解析器：citylistParser

    ```go
    const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
    
    func ParseCityList(contents []byte, _ string) engine.ParseResult {
    	re := regexp.MustCompile(cityListRe)
    	matches := re.FindAllSubmatch(contents, -1)
    	result := engine.ParseResult{}
    
    	for _, m := range matches {
    		result.Requests = append(
    			result.Requests, engine.Request{
    				Url:    string(m[1]),
    				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
    			})
    	}
    	return result
    }
    ```

    - 筛选出城市，并将解析结果存到结果队列中

  - 结点一：城市，遍历页面（**来自城市解析器返回结果的链接**）内可以被**解析（regex库）为用户**的链接，将对应请求发给worker（**城市解析器的目的是找用户**）

    - 城市解析器：CityParser

      ```go
      var (
      	profileRe = regexp.MustCompile(
      		`<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
      	cityUrlRe = regexp.MustCompile(
      		`href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[^"]+)"`)
      )
      
      func ParseCity(contents []byte, _ string) engine.ParseResult {
      	matches := profileRe.FindAllSubmatch(contents, -1)
      	result := engine.ParseResult{}
      	for _, m := range matches {
      		result.Requests = append(
      			result.Requests, engine.Request{
      				Url:    string(m[1]),
      				Parser: NewProfileParser(string(m[2])),
      			})
      	}
      	// 这个部分是解析城市页面中的其他城市：比如看看XX城市
      	nmatches := cityUrlRe.FindAllSubmatch(contents, -1)
      	for _, m := range nmatches {
      		result.Requests = append(result.Requests,
      			engine.Request{
      				Url:    string(m[1]),
      				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
      			})
      	}
      
      	return result
      }
      ```

    - 筛选出**用户的链接**送至result

  - 结点二：用户

    - Item列表 只在这里起作用，因为对本爬虫来说用户才是最有价值的信息

    - 解析器代码

      ```go
      func parserProfile(
      	contents []byte, name string, url string) engine.ParseResult {
      	profile := model.Profile{}
      	profile.Name = name
      	age, err := strconv.Atoi(
      		extractString(contents, ageRe))
      	if err == nil {
      		profile.Age = age
      	}
      
      	height, err := strconv.Atoi(
      		extractString(contents, heightRe))
      	if err == nil {
      		profile.Height = height
      	}
      
      	weight, err := strconv.Atoi(
      		extractString(contents, weightRe))
      	if err == nil {
      		profile.Weight = weight
      	}
      
      	profile.Income = extractString(
      		contents, incomeRe)
      	profile.Gender = extractString(
      		contents, genderRe)
      	profile.Car = extractString(
      		contents, carRe)
      	profile.Education = extractString(
      		contents, educationRe)
      	profile.Hokou = extractString(
      		contents, hokouRe)
      	profile.House = extractString(
      		contents, houseRe)
      	profile.Marriage = extractString(
      		contents, marriageRe)
      	profile.Occupation = extractString(
      		contents, occupationRe)
      	profile.Xinzuo = extractString(
      		contents, xinzuoRe)
      
      	result := engine.ParseResult{
      		Items: []engine.Item{
      			{
      				url,
      				"zhenai",
      				extractString([]byte(url), idUrlRe),
      				profile,
      			},
      		},
      	}
      	matches := guessRe.FindAllSubmatch(contents, -1)
      	for _, m := range matches {
      		result.Requests = append(result.Requests,
      			engine.Request{
      				Url:    string(m[1]),
      				Parser: NewProfileParser(string(m[2])), // 函数调用本身就是拷贝 不需要重新拷贝一份注意！！！
      			})
      	}
      	return result
      }
      ```

### 结果怎么送下去

#### 引擎类 Engine

- 工作原理：维护一个请求队列，当队列不为空时，处理队首内容并弹出队首

- 如何处理队首内容：交给Worker类

  ```go
  func (e SimpleEngine) Run(seeds ...Request) {
  	var requests []Request
  	for _, r := range seeds {
  		requests = append(requests, r)
  	}
  
  	for len(requests) > 0 {
  		r := requests[0]
  		requests = requests[1:]
  
  		log.Printf("Run Fetching %s", r.Url)
  		parseResult, err := Worker(r)
  		if err != nil {
  			continue
  		}
  		requests = append(requests,
  			parseResult.Requests...)
  		for _, item := range parseResult.Items {
  			log.Printf("Got Item %v\n", item)
  		}
  
  	}
  }
  ```

  

#### 工作类 Worker

- Worker代码

  ```go
  func Worker(r Request) (ParseResult, error) {
  	//log.Printf("worker Fetching %v",r)
  	body, err := fetcher.Fetch(r.Url)	// Fetch部分
  	if err != nil {
  		log.Printf("Fetcher: error fetching url %s %v",
  			r.Url, err)
  		return ParseResult{}, err
  	}
  	return r.Parser.Parse(body, r.Url), nil	// 解析部分
  }
  ```

- Worker的工作内容：Fetch和Parse

  - Fetch部分：主要使用http的Get方法获取其html页面（resp.body 以字节流的方式送给对应的Parser）

    ```go
    var rateLimiter = time.Tick(time.Second / config.Qps)
    func Fetch(url string) ([]byte, error) {
    	<-rateLimiter			// 引入速率控制
    	resp, err := http.Get(url) // 这里耗时太长
    	if err != nil {
    		return nil, err
    	}
    	defer resp.Body.Close()
    
    	if resp.StatusCode != http.StatusOK {
    		return nil,
    			fmt.Errorf("wrong status code: %d", resp.StatusCode)
    	}
    	// 关于Reader和Writer部分会单独出一个文档来解析
    	bodyReader := bufio.NewReader(resp.Body)
    	e := determineEncodiong(bodyReader)
    	utf8Reader := transform.NewReader(bodyReader,
    		e.NewDecoder())
    	return ioutil.ReadAll(utf8Reader)
    }
    
    func determineEncodiong(
    	r *bufio.Reader) encoding.Encoding {
    	bytes, err := r.Peek(1024)
    	if err != nil {
    		log.Printf("Fetcher error :%v", err)
    		return unicode.UTF8
    	}
    	e, _, _ := charset.DetermineEncoding(
    		bytes, "")
    	return e
    }
    ```

    

  - Parse部分可参考上面的**解析Regex**（后续的爬虫重构会讲）

    - 每次的解析结果都会保留对于下一个URL的解析函数（被包装起来）

### 整体架构

- 上图

  ![image-20201216204908280](https://i.loli.net/2020/12/16/a3wQMSC1UDn2lIo.png)

- 分析

  - Seed 起始页面
  
  - Engine 负责调度 每次从任务队列中获取一个任务
  
  - Fetcher 引擎将取得的任务交给Fetcher ，Fetcher返回一个字节流
  
  - Parser 引擎将Fetcher返回的字节流送给Parser从而获取下一个request并送入队列
  
    

## 并发版爬虫



## 分布式爬虫