项目迁移URL解决方案
-------------------------
<!--_PAGEDATA
{
    "title": "项目迁移URL解决方案",
    "githubissuesid": 11,
    "createData": "2015-03-08",
    "keywords": "Web开发,团队规范,链接失效,项目迁移，项目架构",
    "description":"介绍应对项目迁移的一种解决方案"
}
_PAGEDATA-->

此文章是 [Rain/doc/前后端开发约定](https://github.com/nimojs/rain/blob/master/doc/%E5%89%8D%E5%90%8E%E7%AB%AF%E5%BC%80%E5%8F%91%E7%BA%A6%E5%AE%9A.md#%E9%A1%B9%E7%9B%AE%E8%BF%81%E7%A7%BB%E8%A7%A3%E5%86%B3%E6%96%B9%E6%A1%88) 中 **-项目迁移解决方案-** 独立通用版本。

考虑如下场景：
```html
<!-- 首页代码 -->
您好，请<a href="/login/">登录</a>
```
项目是一个博客系统，域名是 `http://www.domain.com` 登录地址是 `http://www.domian.com/login/` 。

上线后需求方要求将博客迁移至 `http://www.domain.com/blog/` 。

迁移后访问首页，点击登录`(/login/)`。打开 `/login/` 页面后出现404。因为博客的登录页面变成了 `/blog/login/`，而页面中的链接没有修改。

此时需要将所有页面中的 URL 都加上 `/blog/` 前缀才可以确保所有 URL 正确，`/login/` 改为 `/blog/login/` 等。

------------

当项目迁移至子目录时，因为 URL 前缀固定导致所有页面需要同时修改。我们通过前缀变量的方式解决这个问题。

PHP代码修改如下
```html
define("APP_PATH","/");
您好，请<a href="<php echo APP_PATH ?>login/">登录</a>
```
渲染结果:您好，请`<a href="/login/">登录</a>`

此处是原生 PHP 渲染页面示例，不同后端框架渲染页面方式不同。大致都是定义一个常量，每个 URL 都加上此常量。

使用此方案后，可通过修改常量完成所有页面 URL 的迁移。

```html
define("APP_PATH","/blog/");
您好，请<a href="<php echo APP_PATH ?>login/">登录</a>
```
渲染结果:您好，请`<a href="/blog/login/">登录</a>`

**前端注意：**  
AJAX 路径也需要加上项目路径前缀，防止项目迁移 AJAX 路径错误。参考如下示例：
```javascript
<script>
var APP_PATH = "<php echo APP_PATH ?>";
</script>
<script>
$.get(APP_PATH + 'url/', function () {
	// ...
})
</script>
```

访问Github原文进行讨论：[https://github.com/nimojs/blog/issues/11](https://github.com/nimojs/blog/issues/11)