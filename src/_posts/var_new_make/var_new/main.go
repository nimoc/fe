package main

import (
	"log"
)

func main () {
	{
		var i int
		i++
		log.Print("var i int ",i)
	}
	{
		func() {
			defer func() {
				log.Print(recover())
			}()
			var i *int
			// panic: runtime error: invalid memory address or nil pointer dereference
			// 通过 var 声明的 *int 是空指针 递增会报错
			*i++
			log.Print("var i *int ",*i)
		}()
	}
	{
		var v int
		// 通过 = & 可以避免
		var i *int = &v
		*i++
		log.Print("var i *int = &v ",*i)
	}
	{
		// 或者使用 new （分配空间，并将指针指向零值）
		var i *int = new(int)
		*i++
		log.Print("var i *int = new(int)",*i)
	}
	{
		// 可以进一步简写
		i := new(int)
		*i++
		log.Print("i := new(int)",*i)
	}
	// interface 有可以使用 new 创建一个指向 nil 的指针
	type Closer interface {
		Close() error
	}
	newCloser := new(Closer)
	log.Print("newCloser:", newCloser) // 内存地址
	log.Print("*newCloser:", *newCloser)// nil
	// 不过目前我不知道有什么合适的场景使用 new(interface)
	// 欢迎留言补充 https://github.com/goclub/book/issues/new
}
