----
title: 深入浅出 go int 类型
date: 2020-02-11
keywords: go,go int,
description: 列举一些 int 使用不当的坑
tags:
- go
issues: 52
----

# go中的int

> 这是一篇计算机基础介绍文章，可帮助不了解计算机基础的读者初探计算机基础的同时了解go中的 int
> 文中还列举出一些 int 使用不当的坑

go 中有 `int` 来表示数字类型，并且还有 `uint` 类型，还细分出:

```go
type uint8
type uint16
type uint32
type uint64
type int8
type int16
type int32
type int64
```

## bit

```go
// uint8 is the set of all unsigned 8-bit integers.
// Range: 0 through 255.
type uint8 uint8
```
> 类型 uint8 的数字的范围是 0 到 255

只所以取名 uint8 而不是 uint7 uint9 是因为uint8 使用 8个bit存储数据。

> bit/比特：  binary digit/二进制数字

bit 是计算机表示信息的最小单位，你可以让计算机中的数据非常多的灯泡组成的。灯泡**明**亮时为`1`，灯泡灰**暗**时候为`0`。

> 实际上计算机电路中高电平为1 ，低电平为0

当我们有8个灯泡时，灯泡的明暗可以有好很多中组合，比如:

- 0000 0000
- 0000 0001
- 0000 0010
- ...
- 1111 1110
- 1111 1110

> 为了方便阅读，每四个灯泡增加一个空格。

通过简单的数学算一下 2的8次方

`pow(2,8) = 256`

给每个状态加一个编号，则有256个编号。

- 0: 0000 0000
- 1: 0000 0001
- 2: 0000 0010
- ...
- 254: 1111 1110
- 255: 1111 1110

> 因为从 0 开始编号，所以最后一个状态是255

此时回到计算机的世界，将灯泡视为bit。我们知道了8个bit可以表达256种状态，如果用来表示数字可以 从0表达到255。
故此 uint8 的范围是 0 ~ 256。

如果数据只有一个bit,只能表达2种状态，表达能力有限。所以计算机行业约定8个bit为一个byte(字节)。编程语言中大部分场景以 byte 为最小单位。

```go
// int8 is the set of all signed 8-bit integers.
// Range: -128 through 127.
type int8 int8
```

> 类型 int8 的数字的范围是 -128 到 127

uint 中的 `u` 的意思是 `unsigned` 无符号。没有 `-` 没有 `.` 等数学符号，那就只能是0和正整数。

int 则没有了 `u` 则允许出现负整数，例如 `-1` `-2`。

> int = integer = 整数

而 8个bit只能表达 256 个状态，那需要考虑正数和负数分别占多少。

1. 负整数占 128 个： `-128` 至 `-1`
2. 零 占 1 个
4. 正整数占127个

`128 + 1 + 127 = 256`

所以类型 int8 的数字的范围是 -128 到 127。

弄清楚了 uint8 和 int8 ，再来看 uint16 int16 uint32 int32 uint64 int64 就好理解了。

类型名中int和 uint 后面的数字表示的是在计算机中存储多少bit。

计算机中有很多与 byte(8bit) 相关的数据格式例如 RGB。程序员为了节省空间选择了byte来存储信息，
所以RGB的范围是 `rgb(0,0,0)` 到 `rgb(255,255,255)`。

在 go 中可能出现强制类型转换的代码，例如:

```go
package main
import "log"
func some(n int64) {
	log.Print(n)
}
func main () {
	var a int32
	a = 1
	some(int64(a))
}
```

类似上面的低bit转换为高bit的代码不会发生编译期错误和运行时错误。

但是高bit转换为低bit就会发生意想不到的错误。例如：


```go
package main
import "log"
func some(n uint8) {
	log.Print(n) // 255
}
func main () {
	var a uint16
	a = 65535
	some(uint8(a))
}
```

uint8 强制转换 uint16 在编译期和运行期都不会报错，但大部分情况下会产生bug，因为程序员期望打印的值是 65535。

之所以出现 255 是因为 uint8 最高只能存储255，所以大容量类型转换为小容量类型时需判断容量是否足够。

> TODO: 加上 goclub/conv 实现转换的代码

## byte

## bitmap

## unicode
