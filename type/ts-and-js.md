# 客观看待 TypeScript 与 JavaScript

## TypeScript 只是个类型注解/静态分析

TypeScript 赋予了 JavaScript 静态类型的能力但是由于 JavaScript 的语法向前兼容的历史包袱和 TypeScript 不想限制 JavaScript 本身的灵活性.导致 TypeScript 只是个类型注解,并不是 强类型语言.又因为浏览器运行环境导致在运行时还是以 JavaScript 运行,会出现命名定义了类型是 number 结果运行时候还是可能会是 undefined

## 使用 TypeScript 需要转换思维

母语是 JavaScript 的开发人员去学习 TypeScript 比学习一门新的静态类型语言还要难.因为始终会以 JavaScript 动态灵活的思维去写 TypeScript 代码,而 TypeScript 为了满足 JavaScript 的灵活性,有大量的高级类型语法.有些已经高级到复杂的程度.并且因为复杂度高经常出现新手看不懂的 TypeScript 编译报错.
TypeScript 很优秀.但是历史包袱和依附 JavaScript 本身会导致 TypeScript 很难写. 思维逻辑要转换,要 "自废武功" 去将一些以前 JavaScript 写起来非常简单的代码,用 TypeScript 使用很"啰嗦的"语法去实现.这种"啰嗦"真是静态语言严谨的代码风格. 所以使用 TypeScript 需要转换思维.最好学习一门跟 JavaScript 很像的强静态类型语言,比如我推荐 golang.

TypeScript 历史包袱和依附 JavaScript 本身会导致 TypeScript 很难写. 思维逻辑要转换,要 "自废武功" 去将一些以前 JavaScript 写起来非常简单的代码,用 TypeScript 使用很"啰嗦的"语法去实现.这种"啰嗦"真是静态语言严谨的代码风格. 所以使用 TypeScript 需要转换思维.最好学习一门出生就支持类型系统的语言,比如 java/golang/swift.


## node 中一定要使用 TypeScript

在 node 中一定要使用 TypeScript ,因为后端的各个函数各个数据之间的调用是关联性很强的.相比前端而言,我认为后端是要解决整个面的复杂度,前端是要解决单个点的复杂度,点与点之间的复杂度不是特别高.难度不分高低,根据场景决定.

## 在前端代码中不一定要使用 TypeScript

**在前端代码中不一定非要使用 TypeScript**, 在开发组件和开发核心高复用的模块时一定要用 TypeScript ,因为要确保稳定型和可维护性.但有些页面逻辑代码,将响应数据转换为渲染数据,并根据事件修改渲染数据的这些简单繁琐的逻辑是可以不用 TypeScript 的因为有时候时间不等人,项目 dealline 会逼得你 any 满天飞.

## 面对现实做出明确定义

我们要实事求是,如果你时间允许的情况下应该全部 TypeScript 加上类型,如果在页面琐碎的ui逻辑中时间来不及,是允许使用 any 的.因为现实会让你还是写 any.不如我们**面对现实做出明确定义**,而不是一刀切的不允许用 any 或者一刀切不用 TypeScript.

## 不要让编程语言限制住了自己

JavaScript 的类型系统并不只是只有 TypeScript ,但在前端领域目前类型系统只能用 TypeScript,因为其他类型系统的前端生态不成熟.

在后端领域,如果团队有足够的精力并且明显感觉到了本文说到的js历史包袱和ts后端生态参差不齐,可以考虑使用 [rescript](https://rescript-lang.org.cn/) 开发一套完整的后端工具链. rescript的类型系统更加严格,可减少类型体操的出现,减少类型代码复杂度.

> 后端也不必非要限制在 Node, 了解了解其他母胎就自带强类型系统的语言也是可以的.
