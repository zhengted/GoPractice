# 

[TOC]

# Crawler 技术总结

## 单任务版爬虫

### 总体算法&解析



![image-20201216195546138](https://i.loli.net/2020/12/16/LetE4G7NFgwR5kP.png)

#### 总体算法

- 总体算法类似**广度优先**算法走迷宫机制

#### 解析部分

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

  <img src="C:\Users\szgla\AppData\Roaming\Typora\typora-user-images\image-20201222113228391.png" alt="image-20201222113228391" style="zoom:50%;" />


#### 分析

- Seed 起始页面

- Engine 负责调度 每次从任务队列中获取一个任务

- Fetcher 引擎将取得的任务交给Fetcher ，Fetcher返回一个字节流

- Parser 引擎将Fetcher返回的字节流送给Parser从而获取下一个request并送入队列

## 并发版爬虫

### 核心架构

![image-20201222113840023](C:\Users\szgla\AppData\Roaming\Typora\typora-user-images\image-20201222113840023.png)

#### 并发并在哪里

- Worker是并发创建的，不会引起阻塞

  ```go
  func (e *ConcurrentEngine) Run(seeds ...Request) {
  	// simple调度的写法
  	in := make(chan Request)
  	out := make(chan ParseResult)
  	e.Scheduler.ConfigureMasterWorkerChan(in)
  	fmt.Println("Init scheduler")
  	
  	for i := 0; i < e.WorkerCount; i++ {
  		createWorker(in,out)
  	}
  	
  	for _,r := range seeds {
  		e.Scheduler.Submit(r)
  	}
  }
  //Simple调度的写法
  func createWorker(in chan Request, out chan ParseResult) {
  	go func() {
  		for  {
  			// tell scheduler i'm ready
  			request := <- in	// *1
  			result, err := worker(request)
  			if err != nil {
  				continue
  			}
  			out <- result		// *1和这里循环等待了
  		}
  	}()
  }
  func (s SimpleScheduler) Submit(request engine.Request) {
  	// send request down to worker chan
  	go func() {
  		s.workerChan <- request		// 原写法是将request塞入in中 这里是优化过后的写法 意会即可
  	}()
  }
  ```

  - 但是该调度仍然存在循环等待的问题，如果in取不到数据那么就无法往out里输入数据。out的操作依赖了in

#### 能不能让并发更多

- 调度器部分利用队列来维护，而不是来一个请求就往管道里输入

  - 调度器部分的Run代码

  ```go
  func (q *QueuedScheduler) Run() {
  	q.workerChan = make(chan chan engine.Request)
  	q.requestChan = make(chan engine.Request)
  	go func() {
  		var requestQ []engine.Request
  		var workerQ []chan engine.Request
  		for {
  			var activeRequest engine.Request
  			var activeWorker chan engine.Request
  			if len(requestQ) > 0 && len(workerQ) > 0 {
  				// 这里不执行发 会导致下面的r和w收不到东西
  				activeWorker = workerQ[0]
  				activeRequest = requestQ[0]
  			}
  			select {
  			// 这两个先后顺序不固定 用select控制
  			case r := <-q.requestChan:
  				// send r to a unknown worker
  				requestQ = append(requestQ, r)
  			case w := <-q.workerChan:
  				// send unknown next request to w
  				workerQ = append(workerQ, w)
  			case activeWorker <- activeRequest:
  				workerQ = workerQ[1:]
  				requestQ = requestQ[1:]
  			}
  
  		}
  	}()
  }
  func (q *QueuedScheduler) Submit(request engine.Request) {
  	q.requestChan <- request
  }
  // 告知Q 有一个channel（Worker）已经准备好工作了
  func (q *QueuedScheduler) WorkReady(w chan engine.Request) {
  	q.workerChan <- w
  }
  ```
  - 引擎类的Run

  ```go
  func (e *ConcurrentEngine) Run(seeds ...Request) {
  	// simple调度的写法
  	//in := make(chan Request)
  	//out := make(chan ParseResult)
  	//e.Scheduler.ConfigureMasterWorkerChan(in)
  	//fmt.Println("Init scheduler")
  	//
  	//for i := 0; i < e.WorkerCount; i++ {
  	//	createWorker(in,out)
  	//}
  	//
  	//for _,r := range seeds {
  	//	e.Scheduler.Submit(r)
  	//}
  
  	// 队列调度的写法
  	out := make(chan ParseResult)
  	e.Scheduler.Run()
  	for i := 0; i < e.WorkerCount; i++ {
  		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
  	}
  
  	for _, r := range seeds {
  		if IsDuplicate(r.Url) {
  			log.Printf("Duplicate request:"+"%s", r.Url)
  			continue
  		}
  		e.Scheduler.Submit(r)
  	}
  
  	for {
  		result := <-out
  		for _, item := range result.Items {
  			go func() { e.ItemChan <- item }()		// 这里存到itemsaver中
  		}
  
  		// URL dedup
  		for _, request := range result.Requests {
  			if IsDuplicate(request.Url) {
  				continue
  			}
  			e.Scheduler.Submit(request) 
              // 这里simple调度会有个问题  如果result里的request数量过大会循环等待 解决办法是开一个routine
              // 但是concurrent调度不会
  		}
  	}
  }
  ```

### 引入并发的ItemSaver

- 原有的存储是将数据在控制台打印出来。后续的持久化方式采用存储到ElasticSearch

- 期间学习了docker的使用

- ItemSaver代码如下，其中的channel是在进程开启的时候创建并成为了调度器的成员ItemChan

  ```go
  func ItemSaver(index string) (chan engine.Item, error) {
  	client, err := elastic.NewClient(elastic.SetSniff(false))
  	if err != nil {
  		return nil, err
  	}
  	out := make(chan engine.Item)
  	go func() {
  		itemCount := 0
  		for {
  			item := <-out
  			log.Printf("ItemSaver got item "+
  				"#%d:%v", itemCount, item)
  			itemCount++
  
  			err := Save(client, item, index)
  			if err != nil {
  				log.Printf("Item Saver: error saving item %s\n", err.Error())
  				continue
  			}
  
  		}
  	}()
  	return out, nil
  }
  
  func Save(client *elastic.Client, item engine.Item, index string) error {
  
  	if item.Type == "" {
  		return errors.New("must supply type")
  	}
  	IndexService := client.Index().Index(index).
  		Type(item.Type).
  		Id(item.Id).
  		BodyJson(item)
  	if item.Id != "" {
  		IndexService.Id(item.Id)
  	}
  
  	_, err := IndexService.Do(context.Background())
  
  	if err != nil {
  		return err
  	}
  	return nil
  }
  
  ```

  

## 分布式爬虫

### 继续分离