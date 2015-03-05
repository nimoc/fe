<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title><%=title%></title>
    <link rel="stylesheet" href="../../static/bootstrap/3.3.2/bootstrap.min.css">
    <link rel="stylesheet" href="../../static/nimojs-blog/site.css">
    <meta name="keywords" content="<%=keywords%>" />
    <meta name="description" content="<%=description%>" />
</head>
<body>
<!-- Header S -->
<div class="navbar navbar-inverse navbar-fixed-top">
  <div class="container">
    <div class="navbar-header">
      <button class="navbar-toggle collapsed" type="button" data-toggle="collapse" data-target=".navbar-collapse">
      <span class="sr-only">Toggle navigation</span>
      <span class="icon-bar"></span>
      <span class="icon-bar"></span>
      <span class="icon-bar"></span>
      </button>
      <a class="navbar-brand" href="../../index.html">Home</a>
    </div>
    <div class="navbar-collapse collapse" role="navigation">
      <ul class="nav navbar-nav navbar-right hidden-sm">
        <li>
        <iframe class="githubstar" src="http://ghbtns.com/github-btn.html?user=nimojs&repo=blog&type=watch&count=true&size=little" allowtransparency="true" frameborder="0" scrolling="0" width="100" height="20">
        </iframe>
        </li>
      </ul>
    </div>
  </div>
</div>
<!-- Header E -->
<div class="container">
<%=content%>

<div id="ghComments" class="ui-comments" 
    data-issue-id="<%=githubissuesid%>"
    data-gh-name="nimojs/blog"
>
  <h3>评论</h3>
  <div class="alert alert-warning">
    你可以<a href="https://github.com/nimojs/blog/issues/<%=githubissuesid%>" target="_blank">点击此处</a> 添加评论。
  </div>
  <img class="loading" id="ghCommentsLoading" src="https://assets-cdn.github.com/images/spinners/octocat-spinner-64.gif" alt="加载评论中...">

</div>

</div>
<script src="../../static/jquery/1.11.1/jquery.min.js"></script>
<script src="../../static/nimojs-blog/site.js"></script>
</body>
</html>