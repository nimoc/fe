使用淘宝NPM镜像
===============
<!--_PAGEDATA
{
    "title": "使用淘宝NPM镜像",
    "githubissuesid": 14,
    "createData": "2015-04-19",
    "keywords": "安装 node_modules 失败,安装 node 模块很慢",
    "description":"介绍使用淘宝NPM镜像安装nodejs模块的的方法"
}
_PAGEDATA-->

本文是 [《Gulp 入门指南》 - 安装 Node 和 gulp](https://github.com/nimojs/gulp-book/blob/master/chapter1.md) 的附属教程

> 在使用 node 模块时可以使用 `npm install` 查找 [package.json](https://github.com/nimojs/gulp-demo/blob/master/package.json) 中的声明的依赖模块并安装。

但因为国内网络原因，使用 npm 安装某些模块会非常慢。

例如 [gulp-demo](https://github.com/nimojs/gulp-demo/blob/master/package.json) 依赖了十多个模块，使用 `npm install` 时经常因为网络原因无法安装所有的包。

---------

好在淘宝提供了 npm 镜像供我们使用。
## 安装 cnpm
在命令行输入：
```
npm install -g cnpm --registry=https://registry.npm.taobao.org
```

## 使用 cnpm 安装模块

```
cnpm install [name]
```

> 这真的是个大坑，有时候一个 基于 gulp 的自动化开发环境的依赖包文件会有几十MB。使用 npm 安装很久也不一定能安装完成。

**相关链接：**

- [gulp 入门指南](https://github.com/nimojs/gulp-book)

[点此订阅博客](https://github.com/nimojs/blog/issues/15)

若作者显示不是Nimo（被转载了），请访问Github原文进行讨论：[https://github.com/nimojs/blog/issues/20](https://github.com/nimojs/blog/issues/20)