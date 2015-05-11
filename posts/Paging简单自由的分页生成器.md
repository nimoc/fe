Paging 简单自由的分页生成器
===============
<!--_PAGEDATA
{
    "title": "Paging简单自由的分页生成器",
    "githubissuesid": 23,
    "createData": "2015-04-17",
    "keywords": "paging,分页生成器,ajax无刷新翻译",
    "description":"基于 Mustache 的分页生成器"
}
_PAGEDATA-->

> 传统的分页都是由后端输出 HTML 生成的，使用 AJAX 获取 JSON 显示内容时就需要由前端创建分页。

Paging 只需要知道当前页码和总页数就可以快速生成分页。

Paging 自带了一套默认的模板和样式：  
![](https://cloud.githubusercontent.com/assets/3949015/7386863/150a781c-ee8b-11e4-91a3-ec686b565e50.gif) 

Paging 的使用方法非常简单：  

```html
<link rel="stylesheet" href="http://static.nimojs.com/umd/alice-paging/1.1.0/index.css">
<script src="http://static.nimojs.com/umd/paging/0.0.1/paging.js"></script>

<div id="view"></div>

<script>
var html = Paging.render({
    // 当前页
    currentPage: 2,
    // 总页数
    pageCount: 10,
    // 链接前缀
    link: '?id='
})
document.getElementById('view').innerHTML = html
</script>
```
[预览效果](http://spmjs.io/docs/paging/latest/)

当你需要完全自定义风格时你可以通过[修改样式](http://spmjs.io/docs/paging/latest/examples/index.html) 和 [修改模板](http://spmjs.io/docs/paging/latest/examples/bootstrap.html) 完全控制外观。

- [GitHub](https://github.com/nimojs/paging)
- [示例：自定义界面](http://spmjs.io/docs/paging/latest/examples/index.html)
- [示例：AJAX无刷新分页](http://spmjs.io/docs/paging/latest/examples/ajax.html)
- [示例：控制前后几页显示数量](http://spmjs.io/docs/paging/latest/examples/before-page-count.html)
- [示例：bootstrap 分页](http://spmjs.io/docs/paging/latest/examples/bootstrap.html)
- [示例：handlebars & createData](http://spmjs.io/docs/paging/latest/examples/handlebars.html)

---

[订阅博客](https://github.com/nimojs/blog/issues/15)

若作者显示不是Nimo（被转载了），请访问Github原文进行讨论：[https://github.com/nimojs/blog/issues/23](https://github.com/nimojs/blog/issues/23)