<link rel="stylesheet" type="text/css" href="https://nimo.fun/css/style.css?v=2.11.0" />
{#- Website mata -#}
{%- include "./_head/meta.swig" -%}

{#- Website rss -#}
{%- if theme.rss === 'default' and config.feed and config.feed.path -%}
  {{ feed_tag(config.feed.path) }}
{%- elif theme.rss -%}
  {{ feed_tag(theme.rss) }}
{%- else -%}
  {{ feed_tag('atom.xml') }}
{%- endif -%}

{#- Website favicon -#}
{%- if theme.favicon -%}
  <link rel="shortcut icon" type="image/x-icon" href="{{ url_for(theme.favicon) }}?v={{ theme.version }}" />
{%- endif -%}

{#- Canonical, good for google search engine -#}
<link rel="canonical" href="{{ config.url }}/{{ page.canonical_path.replace('index.html', '') }}"/>

{#- Fancybox styling -#}
{%- if theme.fancybox -%}
  <link rel="stylesheet" type="text/css" href="{{ url_for('lib/fancybox/jquery.fancybox.css') }}" />
{%- endif -%}

{#- MathJax (LaTeX) support -#}
{%- if theme.latex -%}
  <script type="text/x-mathjax-config">
    MathJax.Hub.Config({ tex2jax: { inlineMath: [['$','$'], ['\\(','\\)']] } });
  </script>
  <script type="text/javascript" async src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.5/MathJax.js?config=TeX-MML-AM_CHTML"></script>
{%- endif -%}

{#- Theme styling -#}
<link rel="stylesheet" type="text/css" href="{{ url_for('css/style.css') }}?v={{ theme.version }}" />

{#- Analytics and push -#}
{%- if not env.debug -%}
  {%- include "../_script/analytics.swig" -%}
  {%- include "../_script/push.swig" -%}
{%- endif -%}

{#- LeanCloud Counter -#}
{%- include "../_script/counter.swig"%}

{%- set themeConfig = { leancloud: theme.leancloud, toc: theme.toc, fancybox: theme.fancybox, pjax: theme.pjax, latex: theme.latex } -%}
<script>
  window.config = {{ JSON.stringify(themeConfig) }};
</script>
<style>
.header .logo-wrapper .logo {
font-family: Optima-Regular, Optima, PingFangSC-light, PingFangTC-light, "PingFang SC", Cambria, Cochin, Georgia, Times, "Times New Roman", serif;
letter-spacing:-.02em;
}
html .post .post-content img {
  border:.1em rgba(11,11,11,.2) dotted;
  border-radius:.3em;
  margin-left:auto;
  margin-right:auto;
  display:block;
}
.post .post-header .post-title ,
.header .site-navbar .menu .menu-item-link,
.post .post-content h1,
.mobile-navbar .mobile-header-logo .logo,
html body {
 font-family: Optima-Regular, Optima, PingFangSC-light, PingFangTC-light, "PingFang SC", Cambria, Cochin, Georgia, Times, "Times New Roman", serif;
}
.post .post-content h1 {
display:none;
}
::selection {
    background: #95beea;
    color: #111;
}
html body {
  background:white;
}
.home-tags {
    font-size:12px;
    background:rgba(111,111,111,.1);
    border-radius:.3em;
    padding:.2em .3em;
}
</style>
