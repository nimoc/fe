;`
----
title: 利用空值设计让 TypeScript 更稳定和易于维护
date: 2020-03-01
keywords: typescript空值,typescript规范,typescript技巧
description: 使用 ts 要带入静态类型的思维,虽然会写出很多在js角度看起来很麻烦的类型代码,但是这些代码会让你的项目稳定性更高.
而空值设计可以让编写 ts 更加轻松稳定.
tags:
- 前端
- TypeScript
issues: 33
----

# 利用空值设计让 TypeScript 更稳定和易于维护

[![nimo.fun](http://nimo.fun/notice/index.svg)](https://nimo.fun/notice/)

> 本文代码有大量的 test 和  expect 函数,目的是替代注释,用 expect 说明变量和函数的返回值

## 初始化缺省数据

举一个场景的例子, 前端ajax接收后端响应数据:
`

interface iUser {
    name:string
    age:number
}
test("响应数据", function () {
    let responseJSON :string = `{"name": "nimo"}`
    let res:iUser = JSON.parse(responseJSON)
    expect(res.name).toBe("nimo")
    expect(res.age).toBe(undefined);
    // iUser 中 age 不是 age?:number 但是结果值是 undefined
    // 具体为什么会是 undefined,和为什么后端响应的数据没有 age 我们就不深入讨论了
    // 但是这种情况会导致出现一些bug,比如:
    expect(res.age + 1).toBe(NaN);
    // 明明使用了ts,结果居然将 number 加上 number 得到了 NaN
    // 因为此时 res.age 是 undefined
})

;`

## 数据接口新增属性

接着来看对象数据的场景:

最初定义了一个数据结构,只有 url 属性
`;
interface iPost {
    url:string
}

test("然后有很多地方使用了 iData", function () {
    let a :iPost = {
        url: "https://github.com/nimoc/blog/issues/33"
    }
    console.log(a)
    let b :iPost = {
        url: "https://github.com/nimoc"
    }
    console.log(b)
})

;`
如果此时 iPost 新增了属性 title,则会导致 a b 声明时编译期报错

^^^ts
interface iPost {
    url:string
    title: string
}

test("然后有很多地方使用了 iData", function () {
    // TS2741: Property 'title' is missing in type '{ url: string; }' but required in type 'iPost'.
    let a :iPost = {
        url: "https://github.com/nimoc/blog/issues/33"
    }
    console.log(a)
    // TS2741: Property 'title' is missing in type '{ url: string; }' but required in type 'iPost'.
    let b :iPost = {
        url: "https://github.com/nimoc"
    }
    console.log(b)
})
^^^

想要解决这个问题就需要在 a b 两处声明的地方加上 title 属性,如果不只是 a b 两次,而是由几十处就会变得非常麻烦.

虽然你可以认为类型系统就应该这样严格,但是在这个场景下我更希望不需要改几十处代码

如果不想改几十处可能会导致我们写出不好的代码,例如修改 iPost 为

^^^ts
interface iPost {
    url:string
    title?: string
}
^^^

这种做法虽然不需要些几十处了,但是 [nimo](https://github.com/nimoc) 认为这种方式会引入不必要的 undefined .
导致明明用了类型系统,结果还要处理最繁琐的 undefined 问题.

## 空值设计 (zero values)

我们借鉴 [golang](http://golang.org/) 中zero values 的设计,来解决上述2个问题.


请看代码:
`


interface iPerson {
    name:string
    age:number
}
interface iMakePerson {
    name?:string
    age?:number
}

function Person(v: iMakePerson) :iPerson {
    return {
        name: v.name || "",
        age: v.age || 0
    }
}
test("响应数据空值填充", function () {
    let response = Person(JSON.parse(`{"name":"nimo"}`))
    // 不会出现 response.age 是 undefined 导致的 NaN 的情况
    expect(response.age + 1).toBe(1)
})

test("多处使用 Person ", function () {
    let a  = Person({
        name: "nimo",
    })
    expect(a).toStrictEqual({name:"nimo",age:0})
    let b = Person({
        age: 18,
    });
    expect(b).toStrictEqual({name:"",age:18})
})

;`
如果要新增属性则只需在  iPerson 和 iPerson中分别增加新属性

比如新增了 nikename

^^^ts
interface iPerson {
    name:string
    age:number
    nikename:string
}
interface iMakePerson {
    name?:string
    age?:number
    nikename?:string
}
function Person(v: iMakePerson) :iPerson {
    return {
        name: v.name || "",
        age: v.age || 0,
        nikename: v.nikename || "",
    }
}
^^^

使用所有 Person不会报错,因为接口定义了 nikename?:string

^^^ts
let a  = Person({
    name: "nimo",
})
expect(a).toStrictEqual({name:"nimo",age:0,nikename:""})
^^^


如果新增了 gender ,并且要求 gender 是必填的那么可以这样修改 iPerson

^^^ts
interface iPerson {
    name:string
    age:number
    nikename:string
    gender:string
}
interface iMakePerson {
    name?:string
    age?:number
    nikename?:string
    gender:string
}
function Person(v: iMakePerson) :iPerson {
    return {
        name: v.name || "",
        age: v.age || 0,
        nikename: v.nikename || "",
        gender: v.gender,
    }
}
^^^

注意此时在 iPerson 中 ^gender^ 不是 ^gender?^,没有通过 ? 定义可以为undefined. 这样在所有调用 Person 的地方都需要定义 gender

^^^ts
// 编译期报错
// TS2345: Argument of type '{ name: string; }' is not assignable to parameter of type 'iPerson'.
Person({
    name: "nimo",
})

// 不报错
Person({
    name: "nimo",
    gender: "male",
})
^^^


**通过空值设计可以消除代码中的 undefined , 提高开发效率,增加项目稳定性**

基于空值make函数你可以略过部分属性的声明,不必要写大量的重复代码,但请切记一点在 make 函数中空值只能有

^^^js
""
0
false
[]
另外一个 make 函数
^^^

这是因为如果你在 make 中定义了以上其他的值,会让调用 make 函数的人不明白到底make后属性默认值是什么.

> 不能用 {} 是因为 另外一个make函数替代了空值对象.

另外一个make 函数请看下面的例子

`
interface iSon {
    name:string
}
interface iMakeSon {
    name?:string
}
function Son (v :iMakeSon):iSon {
    return {
        name: v.name || ""
    }
}
interface iFamily {
    unity:boolean
    son: iSon
}
interface iMakeFamily {
    unity?: boolean
    son?: iSon
}
function Family(v :iMakeFamily):iFamily {
    return {
        unity: v.unity || false,
        son: v.son || Son({})
    }
}

test("多层mark",function () {
    let data = Family({
        unity: true,
        // son: Son({}) // 此行可有可无,根据实际场景决定
    })
    expect(data).toStrictEqual({
        "son": {
            "name": ""
        },
        "unity": true
    })
})

;`

----

多说一句,在 ts 中还可以将 son 直接包括在 family 中
`
interface iSome {
    unity:boolean
    son: {
        name:string
    }
}
;`
---

使用 ts 要带入静态类型的思维,虽然会写出很多在js角度看起来很麻烦的类型代码,但是这些代码会让你的项目稳定性更高.
而空值设计可以让编写 ts 更加轻松稳定.

> 如果你学习 typescript 发现怎么用都不顺手,我建议先学习一门纯粹的强静态类型语言(建议 golang ).去掌握强静态类型语言编程思维.
> 因为 TypeScript 是 对 JavaScript进行 类型批注,而不是真正意义上的静态类型语言.


如果你觉得空值函数的设计不错,请将本文推荐给你的朋友或同事

原文地址 https://github.com/nimoc/blog/issues/33 (原文保持持续更新)

这样能让更多人提供更安全的make函数.


`;
