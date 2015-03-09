arguments.callee caller 与 strict mode
==============================
<!--_PAGEDATA
{
    "title": "arguments.callee caller 与 strict mode",
    "githubissuesid": 12,
    "createData": "2015-03-08",
    "keywords": "arguments,arguments.callee,arguments.caller,use strict,strict",
    "description":"严格模式（use strict）为什么不支持 arguments.callee 和 arguments.caller"
}
_PAGEDATA-->

callee 和 caller 用法（熟悉用法可跳过）
--------------------------------

### arguments.callee
> callee 是 arguments 对象的属性。在该函数的函数体内，它可以指向当前正在执行的函数。当函数是匿名函数时，这是很有用的， 比如没有名字的函数表达式 (也被叫做"匿名函数")。

通过2个示例帮助理解：
**打印自身示例**：
```javascript
function testCallee () {
    console.log(arguments.callee);
}
testCallee();
// 控制台输出结果如下：
// function testCallee() {
//     console.log(arguments.callee);
// }
```
**匿名函数递归示例：**
```javascript
var iCalleeRunCount = 0;

setTimeout(function (){
	iCalleeRunCount++
    console.log(iCalleeRunCount);
    setTimeout(arguments.callee,100);
},100);
// 每隔 100毫秒输出递增1的数字：
// 1
// 2
// 3
// 4
// 5
```
运行流程：

一、 延迟100毫秒  
二、 iCalleeRunCount 递增1并输出（值为1）  
三、 延迟100毫秒  
四、 执行 `arguments.callee` 所指向的函数，当前正在执行的函数是
```javascript
function (){
	iCalleeRunCount++
    console.log(iCalleeRunCount);
    setTimeout(arguments.callee,100);
}
```
五、 跳到第二步
六、 无限次调用自身（无限次递归）

### caller
> 如果一个函数f是在全局作用域内被调用的,则f.caller为null,相反,如果一个函数是在另外一个函数作用域内被调用的,则f.caller指向调用它的那个函数.


```javascript
function son () {
    var fCaller = arguments.callee.caller;	// 等同于 var fCaller = son.caller
    if (fCaller === zhang) {
        console.log('我爹叫张三');
    }
    if (fCaller === wang) {
        console.log('我是隔壁老王的儿子');
    }
    if (fCaller === null) {
        console.log('我没爹')
    }
}

function zhang () {
    son();
}
function wang () {
    son()
}

zhang();
// 我爹叫张三
wang();
// 我是隔壁老王的儿子
son(); // 我没爹
```


strict mode 不支持 callee caller
------------------------------

如以上两个章节示例，可通过 callee 与 caller 快速定位到当前执行函数和执行当前函数的函数。（找指向自己和指向执行自己的函数）

但如果你开启了 strict mode 并使用 callee 或 caller 将会报错：  
`Uncaught TypeError: 'caller', 'callee', and 'arguments' properties may not be accessed on strict mode functions or the arguments objects for calls to them`  

如果你不了解什么是 strict mode 请[点击此处](http://www.ruanyifeng.com/blog/2013/01/javascript_strict_mode.html)。

使用 `use strict` 后 不允许使用`arguments.callee` 和 `arguments.caller`。

起初我不明白为什么严格模式会禁用这两个很好用的属性，查阅资料后发现：

> 严格模式下函数的 Arguments 对象定义的非可配置的访问器属性，"caller" 和 "callee"，在它们被访问时，将抛出一个 TypeError 的异常。在非严格模式下，"callee" 属性具有非常明确的意义，"caller" 属性有一个历史问题，它是否被提供，视为一个由实作环境决定的，在具体的 ECMAScript 实作进行扩展。在严格模式下对这些属性的定义的出现是为了确保它们俩谁也不能在规范的 ECMAScript 实作中以任何方式被定义。

得到信息1： **caller 不兼容**

> 因为arguments.caller根本不是ES3支持的东西,它仅仅是IE8-自己的实现.所以完全没必要特别提出来禁止arguments.caller.而悲剧的Chrome和IE10居然还实现了这个错误. 

得到信息2：**caller 是浏览器的私有实现，并非 ECMAScript 标准**

```javascript
[1,2,3,4,5].map(function (n) {
    return !(n > 1) ? 1 : arguments.callee(n - 1) * n;
});
```
> 这段代码性能不好，因为在通常情况下不可能实现内联和尾递归。

得到信息3：**性能不佳**



访问Github原文进行讨论：[https://github.com/nimojs/blog/issues/11](https://github.com/nimojs/blog/issues/12)