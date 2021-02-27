/*
# go 并发实战
*/
package blog_go_channel_test

import (
	"fmt"
	ge "github.com/og/x/error"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

var _ = `
本文假定你已经看过 goroutine channel 的一些教程,
但对于 goroutine channel 依然一知半解,在这个前提下阅读本系列文章.

---
## 写入多个文件 

定义测试用数据:
`

type Data struct {
	FileName string
	Content string
}
var dataList = []Data{
{FileName:"1.txt",Content:"a"},
{FileName:"2.txt",Content:"b"},
{FileName:"3.txt",Content:"c"},
{FileName:"4.txt",Content:"d"},
{FileName:"5.txt",Content:"e"},
{FileName:"6.txt",Content:"f"},
}

// 根据内容写入多个文件
func TestWriteFile(t *testing.T) {
	for _, data := range dataList {
		WriteFile(data)
	}
	/*
	顺序输出
	2020/05/03 19:43:21 1.txt: a
	2020/05/03 19:43:21 2.txt: b
	2020/05/03 19:43:21 3.txt: c
	*/
}
func WriteFile(data Data) {
	err := ioutil.WriteFile(data.FileName, []byte(data.Content), os.ModePerm) ; ge.Check(err)
	fmt.Println(data.FileName + ": " + data.Content)
}

var _ = `

TestWriteFile 方法有个缺陷:
写入操作是队列操作,需要等待前一个写入执行完才能继续写入下一个
而这个场景下的写入文件并不要求队列写入,我们期望能同时进行多个写入操作来提高写入速度

## Goroutine
`

func TestGoroutineWriteFile(t *testing.T) {
	// runtime.GOMAXPROCS(1)
	for _,data := range dataList {
		go WriteFile(data)
	}
	// 必须加上 time.Sleep
	// 因为 go WriteFile(data) 开启了新的 goroutine
	// 如果不加 time.Sleep 等待持续则会直接退出
	time.Sleep(time.Second)
}

var _= `
多运行几次会发现每次输出顺序都是不固定的,不一定是  a b c d e f ,而是无序的.
这说明 go WriteFile(data) 启动的 Goroutine 是并行的

并发和并行的区别用知乎上一个简单的答案来解释是:

> 你吃饭吃到一半，电话来了，你一直到吃完了以后才去接，这就说明你不支持并发也不支持并行。
> 你吃饭吃到一半，电话来了，你停了下来接了电话，接完后继续吃饭，这说明你支持并发。
> 你吃饭吃到一半，电话来了，你一边打电话一边吃饭，这说明你支持并行。

![并发与并行的区别？](https://www.zhihu.com/question/33515481)

你再你的电脑中运行 TestGoroutineWriteFile 那么结果一定是无序的
因为你的电脑CPU一般都是多核的

通过在函数总使用 runtime.GOMAXPROCS(1) 开启"单核模式",则可以模拟单核模式.
单核模式下无法做到并行则会变成队列写入

-----

## channel 

TestGoroutineWriteFile 实现了并行写入文件,提高了写入多个文件执行速度
但我们无法知道准确的完成时间,通过 time.Sleep(time.Second) 实现完成等待是不严谨的.

channel 可以用来在 goroutine 中传递信息

> 请简单看一遍代码然后再看代码后面的解释,再回过头看代码

`

func TestChannelWriteFileDeadlock(t *testing.T) {
	var ch = make(chan string) // 创建通道
	for _, data := range dataList {
		go ChannelWriteFile(data, ch) // 使用 go 关键字开启 goroutine
	}
	for { // 用死循环配合 <- ch 等待 goroutine 将数据写入通道
		writtenFileName := <- ch // <- ch 是堵塞操作,需要等待 goroutine 将数据写入通道
		fmt.Println("writtenFileName: " + writtenFileName)
	}
}

func ChannelWriteFile(data Data, ch chan string) {
	fmt.Println(data.FileName + ": " + data.Content)
	ch <- data.FileName // 将数据写入通道
}

var _= `
输出结果:
	4.txt: d
	writtenFileName: 4.txt
	6.txt: f
	writtenFileName: 6.txt
	5.txt: e
	writtenFileName: 5.txt
	1.txt: a
	2.txt: b
	3.txt: c
	writtenFileName: 1.txt
	writtenFileName: 2.txt
	writtenFileName: 3.txt
	fatal error: all goroutines are asleep - deadlock!


虽然不需要写死 time.Sleep 来等待完成,但是出现了

^ fatal error: all goroutines are asleep - deadlock! ^

这意味着出现了 goroutine 死锁

通过增加一个 count 计数器在准备执行 goroutine 时递增
在从通道接收到数据时候递减,直到递减为 0 时退出死循环防止死锁
要注意递增操作必须在 goroutine 执行前

`


func TestChannelWriteFile(t *testing.T) {
	var ch = make(chan string)
	count := 0 // 定义计数器,用于防止死锁
	for _,data := range dataList {
		count++ // 计数器递增
		go ChannelWriteFile(data, ch)
	}
	for {
		writtenFileName := <- ch
		fmt.Println("writtenFileName: " + writtenFileName)
		count-- // 计数器递减
		if count == 0 { break } // 递减到 0 则退出死循环
	}
}

var _ =`

再次执行会发现不会出现死锁提示了
^ fatal error: all goroutines are asleep - deadlock! ^

实际上只添加了这些代码

^^^go
count := 0 // 定义计数器,用于防止死锁
count++ // 计数器递增

count-- // 计数器递减
if count == 0 { break }
^^^

## 应用场景

比如基于最佳性能角度考虑我们要分两次查询数据库,而这两次查询不需要队列查询

此时则可以使用 goroutine 和 channel 实现并行查询

`

func TestDBQuery(t *testing.T) {
	log.Print("start database query")
	userNameCh := make(chan string)
	bookNameCh := make(chan string)

	go OneUserName("1",userNameCh)
	go OneBookName("1",bookNameCh)
	info := struct {
		UserName string
		BookName string
	}{}
	info.BookName = <- bookNameCh
	info.UserName = <- userNameCh
	log.Printf("%+v", info)
}

func OneUserName(userID string, userNameCh chan string) {
	time.Sleep(time.Second*2) // 模拟io
	log.Print("read userName")
	userNameCh <- "nimoc" // 模拟查询操作
}
func OneBookName(userID string, bookCh chan string) {
	time.Sleep(time.Second*2) // 模拟io
	log.Print("read bookName")
	bookCh <- "life is game" // 模拟查询操作
}

var _= `
因为 ch 传递的数据次数固定所以不需要使用 for 死循环 + break 中断循环
通过 goroutine 和 channel 能利用多核CPU去加速完成读取数据的任务

---

未完待续:

1. 超时设计
2. 单向通道
3. 缓冲通道
4. 死锁等各种坑
`

// 超时设计