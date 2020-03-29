<h1>Web前端写给Web设计师的注意事项</h1>

[![blog.nimoc.io](http://blog.nimoc.io/notice/index.svg)](http://blog.nimoc.io/notice/index.html)

<blockquote>
<p>Web 设计和 Web 前端都应该仔细阅读此文档，会减少因为设计不合理导致的返工。</p>
</blockquote>


<p>Web 设计因为要在浏览器中实现，有时还需要『动』起来，在设计时有一定的限制。</p>


<p><strong>前端同行应该以此文档作为审核设计稿的依据，不应该拿到设计稿直接开发。</strong></p>


<p>有任何问题请 <a href="https://github.com/nimoc/web-desgin-notes/issues/new">参与讨论</a> <a href="https://github.com/nimoc/web-desgin-notes/issues">讨论列表</a></p>


<blockquote>
<p> <strong><a href="https://github.com/nimoc/web-desgin-notes/subscription">Watch</a></strong> 订阅本文档更新</p>
</blockquote>


<hr>

<p><a name="user-content-hash_top" href="https://github.com/nimoc/web-design-notes#hash_top"></a></p>


<p><strong>索引</strong></p>


<p><a href="https://github.com/nimoc/web-design-notes#hash_collect">资源</a></p>


<ol>
<li><a href="https://github.com/nimoc/web-design-notes#hash_size">页面尺寸</a>

<ol>
<li><a href="https://github.com/nimoc/web-design-notes#hash_size_min-width">最小宽度</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_responsive">响应式设计</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_retina">移动设备 Retina</a></li>
</ol></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_font">字体</a>

<ol>
<li><a href="https://github.com/nimoc/web-design-notes#hash_font-size">大小</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_font-special">特殊字体</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_font-icon">字体图标</a></li>
</ol></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_text-overflow">内容溢出</a>

<ol>
<li><a href="https://github.com/nimoc/web-design-notes#hash_text-overflow-ddd">...</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_text-overflow-clip">裁剪</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_text-overflow-tip">提示</a></li>
</ol></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_psd">PSD</a>

<ol>
<li><a href="https://github.com/nimoc/web-design-notes#hash_psd-layer-name">图层命名</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_psd-retina">Retina</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_psd-marker">标注</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_psd-font">字体</a></li>
</ol></li>
<li>栅格化</li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_status">状态</a>

<ol>
<li><a href="https://github.com/nimoc/web-design-notes#hash_status-loading">Loading</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_status-hover">hover</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_status-error">error</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_status-paging">分页</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_status-logout">用户超时登出</a></li>
</ol></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_ui">UI组件化</a>

<ol>
<li><a href="https://github.com/nimoc/web-design-notes#hash_ui-charts">图表</a></li>
</ol></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_typo">typo 内容排版样式</a>

<ol>
<li><a href="https://github.com/nimoc/web-design-notes#hash_typo-rich-text-editor">富文本编辑</a></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_typo-markdown">markdown</a>)</li>
</ol></li>
<li><a href="https://github.com/nimoc/web-design-notes#hash_everyone_checkout">技术团队审核设计稿</a></li>
</ol>


[订阅博客](https://github.com/nimoc/blog/issues/15)

若作者显示不是Nimo（被转载了），请访问Github原文进行讨论：[https://github.com/nimoc/blog/issues/26](https://github.com/nimoc/blog/issues/26)

<script src="https://utteranc.es/client.js"
        repo="nimoc/blog"
        issue-number="15"
        theme="github-light"
        crossorigin="anonymous"
        async>
</script>