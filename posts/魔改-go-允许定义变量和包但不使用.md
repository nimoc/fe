# 魔改 go 允许定义变量和包但不使用


![](https://user-images.githubusercontent.com/3949015/59939027-3872f880-9489-11e9-82cd-ec4e9b8e1b7e.jpeg)

go语言特性要求变量声明后和包引入后必须使用,这是个很好的特性能够避免一些低级错误.

```go
func main () {
  url := "https://github.com/onface/blog/issues/32"
  // 报错 url declared and not used
}
```

但是开发调试的时候这个特性让我很不爽,非常非常不爽.

有些人会说可以用

```go
url := "https://github.com/onface/blog/issues/32"
_ = url
```

但是这就是个屁,是个天坑.

调试完了如果没有去掉 `_ = url` 那么 `url` 这个变量的`declared and not used` 就永远不会出现.

> 再说一遍 `_ = url` 就是个屁

官方不支持那就自己动手将未使用变量报错改为提示,干死你.

> go version go1.12.4 darwin/amd64

访问 GOROOT 目录 (window 路径在哪儿我不知道,自己找)

打开 `/usr/local/go/src/cmd/compile/main.go` 文件

搜索 `imported and not used` 修改

```diff
if name == "" || elem == name {
-   yyerrorl(lineno, "imported and not used: %q", path)
+   Warnl(lineno, "imported and not used: %q", path)
} else {
-   yyerrorl(lineno, "imported and not used: %q as %s", path, name)
+   Warnl(lineno, "imported and not used: %q as %s", path, name)
}
```

其实就是把 `yyerrorl` 改为了 `Warnl`



打开 `/usr/local/go/src/cmd/compile/walk.go` 文件

搜索 `declared and not used` 修改

```diff
if defn := ln.Name.Defn; defn != nil && defn.Op == OTYPESW {
  if defn.Left.Name.Used() {
    continue
  }
-  yyerrorl(defn.Left.Pos, "%v declared and not used", ln.Sym)
+  Warnl(defn.Left.Pos, "%v declared and not used", ln.Sym)
  defn.Left.Name.SetUsed(true) // suppress repeats
} yyerrorl {
-  yyerrorl(ln.Pos, "%v declared and not used", ln.Sym)
+  Warnl(ln.Pos, "%v declared and not used", ln.Sym)
}
```

也是把 `yyerrorl` 改为了 `Warnl`

修改完需要重新编译

```bash
/usr/local/go/src
sudo ./make.bash
```

但是 golang 1.4 版本以上编译需要安装 golang 1.4,所以先要安装好 golang 1.4

下载解压 https://github.com/golang/go/archive/release-branch.go1.4.zip

进入 go 1.4 的 `src` 目录运行

```shell
sudo ./make.bash
```

编译完成后将文件夹重命名为 `go1.4` 并将文件夹移动到 `/Users/用户名/go1.4`.

接着运行

```bash
cd /usr/local/go/src
sudo ./make.bash
```

完事搞定

经此一役,开发阶段通过 IDE 检查未使用变量和包.最终编译时通过 docker 环境编译确保正式代码不存在未使用变量和包.

[![](https://onface.github.io/blog/notice/index.svg)](https://onface.github.io/blog/notice/index.html)

若作者显示不是Nimo（被转载了），请访问Github原文进行讨论：[https://github.com/onface/blog/issues/32](https://github.com/nimojs/blog/issues/32)
