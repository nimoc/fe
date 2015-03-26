使用ruby.taobao安装 Sass
===============
<!--_PAGEDATA
{
    "title": "使用ruby.taobao安装 Sass",
    "githubissuesid": 14,
    "createData": "2015-03-26",
    "keywords": "安装 Sass,Could not find a valid gem 'Sass',安装Sass失败",
    "description":"介绍使用淘宝ruby镜像安装Sass的方法"
}
_PAGEDATA-->


> Sass 是一种 CSS 的开发工具，提供了许多便利的写法，大大节省了设计者的时间，使得 CSS 的开发，变得简单和可维护。

Sass 需使用rubygems 安装，而因为国内网络原因 rubygems 会连接失败，我们可以使用 [ruby.taobao.org](http://ruby.taobao.org/) 提供的镜像安装 Sass。


本文是 [Gulp 入门指南：使用 Gulp 编译 Sass](https://github.com/nimojs/gulp-book) 的附属教程

**目录:**  


- [安装 Ruby](#hash_ruby1)
- [切换 gem 源至 taobao 并安装 Sass](#hash_gem2)


安装 Ruby
-----

[下载安装ruby](http://rubyinstaller.org/)


**检测 Ruby 安装成功**

在 Windows 中可按 徽标键（alt键左边）+ R 打开输入 cmd + 回车打开命令行。

在命令行中输入 `ruby -v` 查看版本号

```
ruby -v
    ruby 2.1.5p273 (2014-11-13 revision 48405) [x64-mingw32]
```


切换 gem 源至 taobao 并安装 Sass
------------------

在命令行依次输入

```
gem sources --remove https://rubygems.org/
    https://rubygems.org/ removed from sources

gem sources -a https://ruby.taobao.org/
    https://ruby.taobao.org/ added to sources

gem source -l
    *** CURRENT SOURCES ***

    https://ruby.taobao.org/

gem install sass
    Fetching: sass-3.4.13.gem (100%)
    Successfully installed sass-3.4.13
    Parsing documentation for sass-3.4.13
    Installing ri documentation for sass-3.4.13
    Done installing documentation for sass after 12 seconds
    1 gem installed
```

如果你在安装时出现如下提示，则表明网络不佳或源没有切换到 ruby.taobao.org 

```
ERROR:  Could not find a valid gem 'sass' (>= 0), here is why:
          Unable to download data from https://rubygems.org/ - SSL_connect retur
ned=1 errno=0 state=SSLv3 read server certificate B: certificate verify failed (
https://api.rubygems.org/latest_specs.4.8.gz)
```

> Mac 下如果一直没有响应则也是网络不佳或没有切换成功

**检测 Sass 安装成功**

```
sass -v
    Sass 3.4.13 (Selective Steve)
```



相关链接：

- [Sass用法指南](http://www.ruanyifeng.com/blog/2012/06/Sass.html)
- [下载安装ruby](http://rubyinstaller.org/)
- [taobao ruby 镜像](http://ruby.taobao.org/)
- [Gulp 入门指南：使用 Gulp 编译 Sass](https://github.com/nimojs/gulp-book)

若作者显示不是Nimo（被转载了），请访问Github原文进行讨论：[https://github.com/nimojs/blog/issues/14](https://github.com/nimojs/blog/issues/14)