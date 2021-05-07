{#- Social Links -#}
{%- include "./_footer/social.swig" -%}

<div class="copyright">
  <a style="font-size:1em;" href="https://beian.miit.gov.cn/" taregt="_blank">沪ICP备2021012662号</a>
  <br />
  <span class="copyright-year" style="display:none;" >
    {%- set current = date(Date.now(), "YYYY") -%}
    &copy;
    {%- if theme.since and theme.since != current -%}
      {{ theme.since }} - {{ current }}
    {%- else -%}
      {{ current }}
    {%- endif -%}

    <span class="heart" hidden >
      <i class="iconfont icon-heart"></i>
    </span>
    <span class="author" hidden >{{ config.author }}</span>
  </span>
</div>
