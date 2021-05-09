----
title: go 中 var new make 的区别
date: 2021-02-11
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

[var_new|embed](./var_new_make/var_new/main.go)

## var make slice array

[var_make_slice_array|embed](./var_new_make/var_make_slice_array/doc_test.go)

## var make map

[var_make_map|embed](./var_new_make/var_make_map/doc_test.go)

## make chan

[make_chan|embed](./var_new_make/make_chan/main.go)

原文地址 https://nimo.fun/go_var_new_make/ (原文保持持续更新)