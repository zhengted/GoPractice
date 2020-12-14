
# Golang基础
[TOC]
## 基本语法
- 基础语法包括以下内容
    - 内建类型
    - 条件语句
    - 循环
    - 内建容器
### 内建类型
- bool string
- (u)int,(u)int8,(u)int16,(u)int32,(u)int64,
- float32,float64
- complex64,complex128 复数类型 1i表示虚部
- byte rune(支持中文 32位)
- 常量和枚举（枚举有个iota的自增技巧）
### 条件语句
- if condition {}
- 延申写法 
```golang
// 分号分割赋值语句
if val, ok = next.at(steps);!ok || val != 0 {     
	continue
}
```
- switch
```golang
switch type {
case: "string"
    // do something
case: "int8"
    // do something
default:
    // do something
}
```
### 循环语句
- golang中没有while，所有的循环都是靠for
- 格式
    - for 初始化;条件;每次循环会调用{}
    - for k,v := range 容器 {}
```golang
// 1.
for i := 1; i < 10; i++ {
    fmt.Println(i)
}
// 2.
a := []int{1,2,3,4,5,6}
for _,v := range a {
    fmt.Println(v)
}
```
### 内建容器
- 容器可以直接用 var name type 来定义，但是此类容器为nil。nil在使用中如果不进行append，会出现各种问题
```golang
var slice []int
slice[1] = 0    // panic
// correct
slice := make([]int,0）
slice := []int{}
```
- array
    - 数组，创建时确定大小
    - 不好用，用的少，一般用slice
- slice
    - 可以理解成数组的视图，所作的修改会改动到数组
    - 删除和增加可以使用append灵活操作
- map
    - 字典，键值对
    - 遍历时是无序的
    - 可以通过外部包强制有序
```golang
import "sort"

var names []string
for name := range ages {
    names = append(names, name)
}
sort.Strings(names)
for _, name := range names {
    fmt.Printf("%s\t%d\n", name, ages[name])
}
```

## 面向对象与面向接口
- struct 结构体
    - 大小写敏感
    - 可以比较，前提是结构体内的是内置类型
    - 可以实现成员函数
- interface 接口
    - 不能定义接口内的成员变量
    - 接口内只能是函数
    - 理解成没有成员变量的虚基类
## 函数式编程
- 函数是一等公民
- 可以作为参数，可以作为结构体
- 也可以作为结构体的成员变量
## 错误处理
- defer和recover
    - defer可以在程序遭遇panic或者return之前进行一次调用
        - 适用于服务器连接，数据库连接，文件操作在open或者connect之后调用一次关闭
    - recover通常和defer一同使用
        - recover可以让程序在遭遇错误时取出当前的error
- 错误的显示
    - 通常在web开发中我们会隐藏错误信息，需要自定义error，服务器显示error，web页面则显示对应的message
```golang
type userError interface {
	error
	Message() string
}
```
## 测试
- 测试的go文件名必须以“_test”结尾
- 函数的开头需要大写的Test
- 表格驱动测试
    - 将条件写入表格构造测试数据
```golang
func TestSubStr(t *testing.T)  {
    // 测试表格
	testData := []struct{
		str string
		res int
	}{
		// Normal Case
		{"abcabcbb",3},
		{"pwwkew",3},

		// Edge Cases
		{"",0},
		{"bbbb",1},
		{"b",1},
		{"abcabcabcd",4},

		// Chinese support
		{"一二三二一",3},
	}
	for _,tt := range testData {
		actual := lengthOfLongestSubstring(tt.str)
		if actual != tt.res {
			t.Errorf("lengthOfLongestSubstring %s expected:%d actual:%d",
				tt.str,tt.res,actual)
		}
	}
}
// 最长不重复子串 长度
func lengthOfLongestSubstring(s string) int {

	for i := range lastOccured {
		lastOccured[i] = -1
	}
	start := 0
	maxLength := 0

	for i,ch := range []rune(s) {
		if lastI := lastOccured[ch];lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
```
- 查看代码覆盖率
    - 可以在点击的时候选择Coverage
    - 命令行
        - go test -coverprofile=c.out
        - go tool cover -html=c.out
- 性能测试
    - 以Benchmark开头，循环测试中的b.N由编译器指定
    - 性能测试一般选择较大的文件做测试
- 查看性能测试结果
    - go test -bench . -cpuprofile cpu.out
    - go tool pprof cpu.out
- 针对服务器整体测试VS针对代码测试 
    - 服务器：速度慢但是覆盖全面  代码：更偏向于单元测试，速度快
## goroutine 和 channel
### routine
- 启动一个协程的方法是使用go关键字
- 协程与启动它的函数是独立开的
- 但是如果主函数返回了，协程也就停止工作了
```golang
// example 
func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
```
### channel
- 协程之间的通信是通过channel进行的
- chan的方向和类型
    - 当为类型时有三种 chan int，chan<- int，<-chan int
    - 使用时:
        - ch<- 表示输入至channel
        - <-ch 表示从channel中取出
#### channel阻塞
- 注意channel是会阻塞导致死锁的
```golang
type worker struct {
	in chan int
	done chan bool
}

func doWork(id int,
	w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c number %d\n",
			id, n, n)
		done <- true
	}
}
func chanDemo() {
	var (
		workers [10]worker
	)

	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i // 阻塞式 第一个循环发了 正在等待done <- true
	}
	for i, worker := range workers {
		worker.in <- 'A' + i
	}
}
```
- 解决方法：
    - 对done这个channel的输入修改为一个新的协程
    - 使用sync中的waitgroup（参考Routines/channel/done/done.go）
### 非阻塞式实现
- select关键字可以实现非阻塞式
- 用法比较复杂
```golang
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			// 生成数据的速度 1.5s
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		// 消耗数据 1s
		fmt.Printf("Worker %d received %d\n",
			id, n)
	}
}

func CreateWorker(id int) chan int { // <-chan：只发不收  chan<- ：只收不发
	c := make(chan int)
	go worker(id, c)
	return c
}

// select 实现非阻塞
func main() {
	var c1, c2 = generator(), generator()
	var worker = CreateWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tk := time.Tick(time.Second)
	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		case <-tk:
			fmt.Println("queue length = ", len(values))
		case <-tm:
			fmt.Println("Bye")
			return
		}
	}
}
```