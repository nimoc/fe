----
title: go 中 var new make 的区别
date: 2020-02-11
keywords: go,go new,go slice,go var,go make
description: 本文主要通过代码示例和原因来解释 var new make 之间的区别。
tags:
- 后端
- go
issues: 49
----


# go 中 var new make 的区别

`var` 用于声明变量。

`new` 分配内存空间，`func new(Type) *Type` 接收 一个类型，返回这个类型的指针，并将指针指向这个类型的零值（zero value）。

`make` 分配内存空间并根据参数初始化

> 本文主要通过代码示例和原因来解释 var new make 之间的区别。

## var new


通过代码记忆最为合适

[var_new](./var_new_make/var_new/main.go)
```.go
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

```

## var make slice array

[var_make_slice_array](./var_new_make/var_make_slice_array/doc_test.go)
```.go
package doc_test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestVarSlice(t *testing.T) {
	log.Print("声明并初始化切片")
	var numbers = []int{1,2,3}
	assert.Equal(t, len(numbers), 3)
}
func TestMakeSliceLen0(t *testing.T) {
			log.Print("通过 make 声明并初始化切片")
			var numbers = make([]int, 0)
			// 等同于 numbers := []int{}
			assert.Equal(t, numbers, []int{})
			assert.Equal(t, len(numbers), 0)
			numbers = append(numbers, 1,2,3)
			assert.Equal(t, len(numbers), 3)
}

func TestMakeSliceLen2(t *testing.T) {
	log.Print("通过 make 声明并初始化切片长度为2， 初始化元素为类型的 zero value (0)")
	var numbers = make([]int, 2)
	// 等同于 numbers := []int{0, 0}
	assert.Equal(t, numbers, []int{0, 0})
	assert.Equal(t, len(numbers), 2)
	numbers = append(numbers, 1,2,3)
	assert.Equal(t, numbers, []int{0, 0, 1, 2, 3})
	assert.Equal(t, len(numbers), 5)
}

func TestArray(t *testing.T) {
	log.Print("创建数组（固定长度的切片）数组元素为 zero value")
	// 此处的 2 是 cap
	var numbers = [2]int{}
	assert.Equal(t, numbers[0], 0)
	assert.Equal(t, numbers[1], 0)
	// // 如果 index(10) 超出了 cap(2) 则会发生**编译期**错误
	// log.Print(numbers[10])
	assert.Equal(t, len(numbers), 2)
}

func TestArrayLen2Cap2(t *testing.T) {
	log.Print("通过 make 创建长度2 容量2的数字，初始化元素为类型的 zero value (0)")
	var numbers = make([]int, 2, 2)
	// 等同于 numbers := []int{0,0}
	assert.Equal(t, numbers, []int{0,0})
	assert.Equal(t, len(numbers), 2)
	assert.Equal(t, cap(numbers), 2)
	numbers[0] = 9
	assert.Equal(t, numbers, []int{9,0})
	// 如果 index(10) 超出了 cap(2) 则会发生**运行时**错误
	// numbers[10] = 9
}

func TestArrayLen0Cap2(t *testing.T) {
	log.Print("通过 make 创建长度0 容量2的数组")
	var numbers = make([]int, 0, 2)
	assert.Equal(t, numbers, []int{})
	assert.Equal(t, len(numbers), 0)
	assert.Equal(t, cap(numbers), 2)

	numbers = append(numbers, 1)
	assert.Equal(t, numbers, []int{1})
	assert.Equal(t, len(numbers), 1)
	assert.Equal(t, cap(numbers), 2)

	numbers = append(numbers, 2)
	assert.Equal(t, numbers, []int{1,2})
	assert.Equal(t, len(numbers), 2)
	assert.Equal(t, cap(numbers), 2)

	numbers = append(numbers, 3)
	assert.Equal(t, numbers, []int{1,2,3})
	assert.Equal(t, len(numbers), 3)
	assert.Equal(t, cap(numbers), 4)
}

func TestMakeArrayLen0Cap2Panic(t *testing.T) {
	log.Print("大部分场景下 numbers := make([]string, len, cap) len 和 cap 设置的不一致是没有意义的")
	func() {
		defer func() {
			log.Print(recover())
		}()
		numbers := make([]int, 0, 2)
		log.Print(`因为 numbers[0] = 1 会panic`)
		numbers[0] = 1
	}()
}
```

## var make map

[var_make_map](./var_new_make/var_make_map/doc_test.go)
```.go
package doc_test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestVarMap(t *testing.T) {
	log.Print("声明变量，但没有分配空间")
	var data map[string]int
	assert.Equal(t, data, map[string]int(nil))
}
func TestDeclareEmptyMap(t *testing.T) {
	defer func() {
		log.Print(recover())
	}()
	log.Print("只声明变量，不分配空间，赋值时会panic")
	var data map[string]int
	data["age"] = 1
}
func TestMap(t *testing.T) {
	log.Print("声明变量，并初始化赋值")
	var data = map[string]int{}
	assert.Equal(t, data, map[string]int{})

	data["name"] = 1
	assert.Equal(t, data, map[string]int{"name":1})
}
func TestMakeMap(t *testing.T) {
	log.Print("通过 make 分配空间，也可以避免panic")
	var data = make(map[string]int)
	assert.Equal(t, data, map[string]int{})

	data["name"] = 1
	assert.Equal(t, data, map[string]int{"name":1})
}

```

## make chan

[make_chan](./var_new_make/make_chan/main.go)
```.go
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

```
