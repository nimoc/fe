# 将 TypeScript 中 松散的类型当做药品


> 本文代码有大量的 test 和  expect 函数,目的是替代注释,用 expect 说明变量和函数的返回值


## 动态语言不需要泛型

基于显而易见的原因如果你使用的是动态语言没有类型系统意味着一切都是泛型.

我通过列举一个 `filterZeroValue` 的例子来说明情况:

> 为了把重点放在类型系统上所以使用 `filterZeroValue` 这个简单的函数,实际情况中不大可能封装 `filterZeroValue` 而是直接写 `list.filter`.

比如在 JavaScript 中:

```ts
/*
 * 排除数组中的空值'
 * @param list
 * @return notZeroValueList
 * */
function jsFilterZeroValue(list) {
    return list.filter(function (item) {
        switch (typeof item) {
            case "string":
                return item != ""
            break
            case "number":
                return item != 0
                break
            default:
            throw new Error("filterZeroValue: list[] item must be string or number" )
        }
    })
}

test("jsFilterZeroValue", function () {
    expect(jsFilterZeroValue(["a","","c"])).toStrictEqual(["a","c"])
    expect(jsFilterZeroValue([1,0,3])).toStrictEqual([1,3])
})


```


你甚至可以3行代码搞定

```js
function jsFilterZeroValue() {
    return list.filter((item)=> {return !!item})
}
```


不这么做是因为需要在参数是 string number 之外的的类型时进行错误提示,和减少隐式类型转换.


## TypeScript 实现泛型

> 注意不要只看下面的代码后就结束,看完文章会发现下面的代码是不好的

```ts

function tsFilterZeroValue<T>(list: T[]): T[] {
    return list.filter(function (item) {
        switch (typeof item) {
            case "string":
                return item != ""
                break
            case "number":
                return item != 0
                break
            default:
                throw new Error("filterZeroValue: list[] item must be string or number")
        }
    })
}


test("tsFilterZeroValue", function () {
    expect(tsFilterZeroValue(["a","","c"])).toStrictEqual(["a","c"])
    expect(tsFilterZeroValue([1,0,3])).toStrictEqual([1,3])
})

```

虽然通过 `<T>(list: T[]): T[]` 约束了必须是个数组,并且输出的类型和输入的类型一致.但是还是不能明确只允许 `number[]` `string[]` .


### 联合类型

上面的列子可以用联合类型来解决,但是联合类型也不够好哦.

> 联合类型和泛型其实是一类方法,在现在的这个场景的目的就是偷懒.

```ts

function unionTypeFilterZeroValue(list: string[] | number[]) :string[] | number[] {
    let output = []
    for (let i= 0;i<list.length;i++ ) {
        const item = list[i]
        switch (typeof item) {
            case "string":
                if (item != "") {
                   output.push(item)
                }
                break
            case "number":
                if (item != 0) {
                    output.push(item)
                }
                break
            default:
                throw new Error("filterZeroValue: list[] item must be string or number")
        }
    }
    return output
}


test("unionTypeFilterZeroValue", function () {
    expect(unionTypeFilterZeroValue(["a","","c"])).toStrictEqual(["a","c"])
    expect(unionTypeFilterZeroValue([1,0,3])).toStrictEqual([1,3])
})


```


### TypeScript filterZeroValue 正确的实现

上面的 `tsFilterZeroValue` 和 `unionTypeFilterZeroValue` 都不够好,
反而 TypeScript 代码写的很复杂.虽然可能是我个人对 TypeScript 了解程度不够,
要注意团队中不是每个人都是 TypeScript 高手.

实际上在TypeScript中使用泛型绝大部分情况下是编码思维没有转换为静态类型思维.

请看下面的代码

```ts

function stringListFilterZeroValue(list: string[]) :string[] {
    return list.filter(function (v) {
        return v != ""
    })
}

function numberListFilterZeroValue(list: number[]) :number[] {
    return list.filter(function (v) {
        return v != 0
    })
}

test("stringAndNumberlistFilterZeroValue", function () {
    expect(stringListFilterZeroValue(["nimo","","nico"])).toStrictEqual(["nimo","nico"])
    expect(numberListFilterZeroValue([1,0,3])).toStrictEqual([1,3])
})

```


> 不要带入动态类型快猛糙的思维去写 TypeScript

该多写点"重复"的代码,这样反而实现会更简单,更易于阅读.

最重要的是有些情况下使用了泛型或联合类型加上编码时疏忽了会造成想不到的bug:

```ts

function updateSQL(id: string, names: string[]) :{sql:string, values:any[]} {
    const updateValue = stringListFilterZeroValue(names)
    // 如果 updateSQL 的函数参数 names 改成了  ages int[]
    // stringListFilterZeroValue 将会在编译期报错
    // 如果使用的是 unionTypeFilterZeroValue 则不会

    // names 修改后 要让此处编译期报错的目的是要
    // 提醒自己,在没有修改前的代码逻辑中期望 updateValue 是一个 string[]
    // 如果使用 unionTypeFilterZeroValue 则没有了这一层提醒
    // 而 JSON.stringify(string[]) 和 JSON.stringify(number[]) 的结果是不一样的
    // 而这个不一样类型系统是无法检查到的,因为返回值 的 values 属性因为 sql 的场景导致就是 any[]
    return {
        sql: `UPDATE tableName SET names = ? WHERE id = ?`,
        values: [JSON.stringify(updateValue), id],
    }
}



test("updateSQL", function (){
    expect(updateSQL("1", ["nimo", "nico"])).toStrictEqual(
        {
            sql: "UPDATE tableName SET names = ? WHERE id = ?",
            values: [
                '["nimo","nico"]',
                "1",
            ],
        }
    )
})


```


上面的例子不够完美,本文想表达的主要的观点是:

**控制参数数量和类型不可变**

在代码中明确函数参数固定且每个参数只能有一个类型能让代码更易于维护

**尽可能多的在编译期做类型检查发现问题**

即使单元测试和细心编码能检查出这种小概率的错误,但是编码要做悲观设计.不能总期望写代码的人状态在线

**将松散的类型当做药品使用**

泛型,联合类型这种应当当做药品去使用,不到万不得已不要使用.比如 Go 语言中就不支持 TypeScript 这种泛型,也照样构建了那么稳定的项目,
只要不是觉得业务代码中出现大量重复代码太麻烦,就要避免使用松散的类型.非业务逻辑的第三方封装代码,就必须让参数类型只能有一个.
除非你实现是 JSON.parse 这种必须用 any 的库.


> 有些人对于效率和质量的认知可能与作者有偏差,作者是绝对侧重质量,在要效率非常低下的情况下才通过深思熟虑的才写一些"偷懒的代码".
> 读者可以有自己的判断,但请注意: 如果因为类型不严谨导致项目中出现一个 bug,如果能后悔你会愿意花十倍的时间去弥补写出更多类型严谨的代码.


如果你觉得本文观点不错,请将本文推荐给你的朋友或同事


在Github发表评论: https://github.com/nimoc/fe/discussions/57

```ts;
```
