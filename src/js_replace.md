# 深入理解JavaScript-replace

[![blog.nimoc.io](http://blog.nimoc.io/notice/index.svg)](http://blog.nimoc.io/notice/index.html)

replace方法是属于String对象的，可用于替换字符串。
## 简单介绍:

`String.replace(searchValue,replaceValue)`
1.  String:字符串
2.  searchValue：字符串或正则表达式
3.  replaceValue:字符串或者函数
## 字符串替换字符串

``` javascript
'I am loser!'.replace('loser','hero')
//I am hero!
```

直接使用字符串能让自己从loser变成hero，但是如果有2个loser就不能一起变成hero了。

``` javascript
'I am loser,You are loser'.replace('loser','hero');
//I am hero,You are loser 
```
## 正则表达式替换为字符串

``` javascript
'I am loser,You are loser'.replace(/loser/g,'hero')
//I am hero,You are hero
```

使用正则表达式，并将正则的global属性改为true则可以让所有loser都变为hero
## 有趣的替换字符

`replaceValue` 可以是字符串。如果字符串中有几个特定字符的话，会被转换为特定字符串。

| 字符 | 替换文本 |
| :-- | :-- |
| $& | 与正则相匹配的字符串 |
| $` | 匹配字符串左边的字符 |
| $' | 匹配字符串右边的字符 |
| $1,$2,$3,…,$n | 匹配结果中对应的分组匹配结果 |
### 使用$&字符给匹配字符加大括号

``` javascript
var sStr='讨论一下正则表达式中的replace的用法';
sStr.replace(/正则表达式/,'{$&}');
//讨论一下{正则表达式}中的replace的用法
```
### 使用$`和$'字符替换内容

``` javascript
'abc'.replace(/b/,"$`");//aac
'abc'.replace(/b/,"$'");//acc
```
### 使用分组匹配组合新的字符串

``` javascript
'nimoc@126.com'.replace(/(.+)(@)(.*)/,"$2$1")//@nimoc
```
## replaceValue参数可以是一个函数

`String.replace(searchValue,replaceValue)` 中的**replaceValue**可以是一个函数.

如果replaceValue是一个函数的话那么，这个函数的arguments会有n+3个参数（n为正则匹配到的次数）

**先看例子帮助理解：**

``` javascript
function logArguments(){    
    console.log(arguments);//["nimoc@126.com", "nimoc", "@", "126.com", 0, "nimoc@126.com"] 
    return '返回值会替换掉匹配到的目标'
}
console.log(
    'nimoc@126.com'.replace(/(.+)(@)(.*)/,logArguments)
)
```

**参数分别为**
1.  匹配到的字符串（此例为nimoc@126.com,推荐修改上面代码的正则来查看匹配到的字符帮助理解)
2.  如果正则使用了分组匹配就为多个否则无此参数。（此例的参数就分别为`"nimoc", "@", "126.com"`。推荐修改正则为/nimo/查看控制台中返回的arguments值）
3.  匹配字符串的对应索引位置（此例为0）
4.  原始字符串(此例为nimoc@126.com)
### 使用自定义函数将A-G字符串改为小写

``` javascript
'JAVASCRIPT'.replace(/[A-G]/g,function(){
    return arguments[0].toLowerCase();
})//JaVaScRIPT 
```
### 使用自定义函数做回调式替换将行内样式中的单引号删除

``` javascript
'<span style="font-family:\'微软雅黑\';">;demo</span>'.replace(/\'[^']+\'/g,function(){      
    var sResult=arguments[0];
    console.log(sResult);//'微软雅黑'
    sResult=sResult.replace(/\'/g,'');
    console.log(sResult);//微软雅黑
    return sResult;
})//<span style="font-family:微软雅黑;">demo</span> 
```
## 最后的小试牛刀

这一节是交给阅读者发挥的内容：
### 洗扑克

需要将Thisnimoc-JavaScript使用正则替换成 `TJhaivsaNSicmroijpst`

[点此订阅博客](https://github.com/nimoc/blog/issues/15)

若作者显示不是Nimo（被转载了），请访问Github原文进行讨论：[https://github.com/nimoc/blog/issues/2](https://github.com/nimoc/blog/issues/2)

<script src="https://utteranc.es/client.js"
        repo="nimoc/blog"
        issue-number="15"
        theme="github-light"
        crossorigin="anonymous"
        async>
</script>