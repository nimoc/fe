# JavaScript初学者建议：不要去管浏览器兼容

[![blog.nimoc.io](http://blog.nimoc.io/notice/index.svg)](http://blog.nimoc.io/notice/index.html)


如果可以回到过去的话，我会告诉自己这句话：**"初学 JavaScript 的时候无视 DOM 和 BOM 的兼容性"**
## 我初学时的处境

在我初学 JavaScript 的时候最头痛的就是浏览器兼容问题。在 Firefox 下面好好的代码到 IE 就不能显示了，又或者是在 IE 能正常显示的代码在 Firefox 又报错了。

前端开发工程师的职责就包括跨浏览器开发。所以我就**在还不了解 JS 这门语言本身的时候去花时间学习浏览器兼容知识**，**这样会让JS学习难度增加**。**但是不能兼容主流浏览器的代码不能用在实际项目中。**

**DOM 和 BOM 的兼容性问题一度让我的 JavaScript 学习停滞不前**。语言理解不够，代码又只能在特定浏览器运行。
## 我的建议

如果你正初学 JavaScript 并有着和我一样的处境的话我建议你：**初学 JavaScript 时无视 DOM 和 BOM 的兼容性，将更多的时间花在了解语言本身（ECMAScript）**。只在特定浏览器编写代码（Chrome/Firefox/Safari），实际工作中使用成熟的 JavaScript 框架（jQuery等）。放心，很少有公司会让 JS 新手用原生 JS 做前端开发。
### 学习 JS 初期无视兼容问题有什么好处
1. 降低学习难度
2. 减少挫败感
3. 花更多的时间学习 ECMAScript
## 什么时候学习JS跨浏览器开发知识

而浏览器兼容问题留到什么时候解决呢？  
当你能熟练使用 JavaScript 框架编写可复用的代码时（jQuery插件或前端控件），或当你准备自己开发一个 JavaScript 框架时。
## 其他一些 JavaScript 初学者建议
1. 无编程经验千万不要拿JavaScript权威指南当入门书籍
2. 应该用JavaScript高级程序设计最新版本（目前是第三版）作为入门书籍
3. 传值和传址、作用域知识必须理解
4. 调试工具必须懂并多用，学会自己捕捉错误。（chrome developer tool/Firebug）
5. 耐心再耐心，对每一个知识点深挖能学的更轻松。

以上就是我的一些分享希望若能帮助到初学 JavaScript 的你，如果觉得有误导的地方敬请立即指出。

[点此订阅博客](https://github.com/nimoc/blog/issues/15)

若作者显示不是Nimo（被转载了），请访问Github原文进行讨论：[https://github.com/nimoc/blog/issues/1](https://github.com/nimoc/blog/issues/1)

<script src="https://utteranc.es/client.js"
        repo="nimoc/blog"
        issue-number="15"
        theme="github-light"
        crossorigin="anonymous"
        async>
</script>