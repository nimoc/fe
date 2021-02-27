# JavaScript原型赋值陷阱
[![nimoc.io](http://nimoc.io/notice/index.svg)](https://nimoc.io/notice/)

不通过对象的 `constructor.prototype` 对原型中的属性进行递增时候会触发原型赋值陷阱。

本文通过一个 Popup 弹出框来解释陷阱的出现情况，并说明如何找到陷阱和解决问题。并且本文假设你至少简单了解 JavaScript 中的原型。
## 记录alert次数的弹出框

**单实例调用**

``` js
var Popup=function(){
}
Popup.prototype.alert=function(message){
    this.iMessageCount++;
    alert(message+'~alert过'+this.iMessageCount+'次');    
}
Popup.prototype.iMessageCount=0;
var oNimo=new Popup();
oNimo.alert('你好我是nimo!');//  alert 过1次
oNimo.alert('Nice to meet you,I am Nimo!');//  alert 过2次
```

代码解释
1. 创建构造函数 Popup
2. 给 Popup 添加 alert 方法。弹出内容是消息加弹出次数，每次弹出递增 iMessageCount 属性。
3. 添加公用属性 iMessageCount 用于记录弹出次数。
4. 创建oNimo实例，并用 oNimo 弹出2次内容。

弹出内容
1. 你好我是nimo!~alert过1次
2. Nice to meet you,I am Nimo!~alert过2次

**添加一个实例**

在上面的代码底部添加如下代码

```
var oDemo=new Popup();
oDemo.alert('我是demo!'); //alert过1次  
```

代码解释
1. 创建oDemo实例，并用oDemo弹出2次内容。

弹出内容
1. 我是demo!~alert过1次

oDemo 的弹出结果应该是 alert 过3次，结果却是 alert 过1次
## debug

遇到 bug 先将相关对象输出检查

``` js
console.log(oNimo);
console.log(oDemo);
```

![](https://cloud.githubusercontent.com/assets/3949015/7004697/0f1560d4-dca2-11e4-8c19-0668203b9000.png)
打印结果后发现原型中 iMessageCount 属性并没有递增，依然是0。而 oNimo 和 oDemo 自身属性中却存储着 iMessageCount 属性，分别是2和1。说明 `this.iMessageCount++` 递增的是对象自身属性并不是原型属
### 拆分 bug

既然问题出在 `this.iMessageCount++` 那么就对这行代码进行详细分析。

以下三行代码实际相等

``` js
this.iMessageCount++
this.iMessageCount=this.iMessageCount+1
this.iMessageCount=this.constructor.prototype.iMessageCount+1
```

解释
1. 递增操作
2. iMessageCount属性等于iMessageCount属性+1
3. 因为一开始对象自身并没有iMessageCount属性而原型中有，所有结果是将原型属性中的iMessageCount属性+1并赋值给对象自身属性中的iMessageCount属性。

当调用新的 `oDemo` 时并没有修改 `Popup` 的原型。和上面一样，只是获取了 `Popup.prototype.iMessage` 的值。

知识点：对象访问一个属性会首先查找自身属性如果找不到自身属性就查找对象的 constructor 中的 prototype 中的属性（对象构造函数的原型中的属性）。

跳过陷阱

如需对原型中的属性进行递增操作请直接对对象的 constructor 中的 protorype 中的属性进行递增。

修复后的代码：

关键代码： `this.constructor.prototype.iMessageCount++`

完整代码：

``` js
var Popup=function(){
}
Popup.prototype.alert=function(message){
    this.constructor.prototype.iMessageCount++;
    alert(message+'~alert过'+this.iMessageCount+'次');    
}
Popup.prototype.iMessageCount=0;
var oNimo=new Popup();
oNimo.alert('你好我是nimo!');//  alert 过1次
oNimo.alert('Nice to meet you,I am Nimo!');//  alert 过2次

var oDemo=new Popup();
oDemo.alert('我是demo!'); //alert过3次
```
## 小结

不通过对象的 constructor.prototype 对原型中的属性进行递增时候会触发原型递增陷阱。

如需对原型中的属性进行递增操作请直接对对象的constructor中的protorype中的属性进行递增。


[点此订阅博客](https://github.com/nimoc/blog/issues/15)

若作者显示不是Nimo（被转载了），请访问Github原文进行讨论：[https://github.com/nimoc/blog/issues/17](https://github.com/nimoc/blog/issues/17)
