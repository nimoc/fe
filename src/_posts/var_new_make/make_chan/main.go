package main

import "log"

func main() {
	{
		// 只是声明变量
		var nameCh chan string
		// 通过 make 初始化空间
		nameCh = make(chan string) // 注释这一行会因为 nameCh 没有分配内存空间导致死锁
		log.Print("nameCh ", nameCh) // 内存地址
		go func() {
			nameCh <- "nimoc"
		}()
		name := <-nameCh
		log.Print(name)
	}
	{
		{
			// 代码可以更简洁一点
			nameCh := make(chan string)
			log.Print("nameCh ", nameCh) // 内存地址
			go func() {
				nameCh <- "nimoc"
			}()
			name := <-nameCh
			log.Print(name)
		}
	}
}
