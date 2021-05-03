{#- Social Links -#}
{%- include "./_footer/social.swig" -%}

<div class="copyright">
  

  <span class="copyright-year">
    {%- set current = date(Date.now(), "YYYY") -%}
    &copy;
    {%- if theme.since and theme.since != current -%}
      {{ theme.since }} - {{ current }}
    {%- else -%}
      {{ current }}
    {%- endif -%}

    <span class="heart">
      <i class="iconfont icon-heart"></i>
    </span>
    <span class="author">{{ config.author }}</span>
  </span>
</div>
