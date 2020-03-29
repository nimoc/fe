# 《Gulp 入门指南》- 前言

[![blog.nimoc.io](http://blog.nimoc.io/notice/index.svg)](http://blog.nimoc.io/notice/index.html)

<!--_PAGEDATA
{
    "title": "《Gulp 入门指南》",
    "githubissuesid": 19,
    "createData": "2015-04-17",
    "keywords": "gulp,gulp入门指南,gulp教程,gulp编译sass,gulp压缩js",
    "description":"gulp 是基于 node 实现 Web 前端自动化开发的工具，利用它能够极大的提高开发效率。"
}
_PAGEDATA-->

gulp 是基于 node 实现 Web 前端自动化开发的工具，利用它能够极大的提高开发效率。

在 Web 前端开发工作中有很多“重复工作”，比如压缩CSS/JS文件。而这些工作都是有规律的。找到这些规律，并编写 gulp 配置代码,让 gulp 自动执行这些“重复工作”。

> 你可以直接访问 [GitHub: gulp-book](https://github.com/nimoc/gulp-book) 查看本书
## 目录
- [安装 Node 和 gulp](https://github.com/nimoc/gulp-book/blob/master/chapter1.md)
- [使用 gulp 压缩 JS](https://github.com/nimoc/gulp-book/blob/master/chapter2.md)
- [使用 gulp 压缩 CSS](https://github.com/nimoc/gulp-book/blob/master/chapter3.md)
- [使用 gulp 压缩图片](https://github.com/nimoc/gulp-book/blob/master/chapter4.md)
- [使用 gulp 编译 LESS](https://github.com/nimoc/gulp-book/blob/master/chapter5.md)
- [使用 gulp 编译 Sass](https://github.com/nimoc/gulp-book/blob/master/chapter6.md)
- [使用 gulp 构建一个项目](https://github.com/nimoc/gulp-book/blob/master/chapter7.md)
## 将规律转换为 gulp 代码

现有目录结构如下：

```
└── js/
    └── a.js
```
### 规律
1. 找到 js/目录下的所有 .js 文件
2. 压缩这些 js 文件
3. 将压缩后的代码另存在 dist/js/ 目录下
### 编写 gulp 代码

``` js
// 压缩 JavaScript 文件
gulp.task('script', function() {
    // 1. 找到
    gulp.src('js/*.js')
    // 2. 压缩
        .pipe(uglify())
    // 3. 另存
        .pipe(gulp.dest('dist/js'));
});
```
### 代码执行结果

代码执行后文件结构

```
└── js/
│   └── a.js
└── dist/
    └── js/
        └── a.js
```

a.js 压缩前

```
function demo (msg) {
    alert('--------\r\n' + msg + '\r\n--------')
}

demo('Hi')
```

a.js 压缩后

```
function demo(n){alert("--------\r\n"+n+"\r\n--------")}demo("Hi");
```

此时 `dist/js` 目录下的 `.js` 文件都是压缩后的版本。

你还可以监控 `js/` 目录下的 js 文件，当某个文件被修改时，自动压缩修改文件。启动 gulp 后就可以让它帮助你自动构建 Web 项目。

---

gulp 还可以做很多事，例如：
1. 压缩CSS
2. 压缩图片
3. 编译Sass/LESS
4. 编译CoffeeScript
5. markdown 转换为 html
- [开始阅读：安装 Node 和 gulp](https://github.com/nimoc/gulp-book/blob/master/chapter1.md)
- [目录](https://github.com/nimoc/gulp-book/blob/master/SUMMARY.md)
- [订阅本书](https://github.com/nimoc/gulp-book/subscription)
- [订阅博客](https://github.com/nimoc/blog/issues/15)

若作者显示不是Nimo（被转载了），请访问Github原文进行讨论：[https://github.com/nimoc/blog/issues/19](https://github.com/nimoc/blog/issues/19)

<script src="https://utteranc.es/client.js"
        repo="nimoc/blog"
        issue-number="19"
        theme="github-light"
        crossorigin="anonymous"
        async>
</script>