## 微信小程序广告对接指导手册

[![nimoc.io](http://nimoc.io/notice/index.svg)](https://nimoc.io/notice/index.html)

友情链接: [fucking-weapp](https://github.com/onface/fucking-weapp)

## 业务场景

因公司业务需求，需要**在微信小程序中插入广告**。点击广告后跳转至广告客户的小程序中，或展示广告客户的信息。

或者需要推广自己公司的微信小程序，需要支付其他公司广告费用，其他公司通过各种方式**引导普通用户进入微信小程序**。

## 支持的对接方式

### 流量方

1. 跳转其他小程序
2. 跳转APP
3. 打开网页

### 广告主

1. 普通链接
2. 公众号推文
3. APP直接打开小程序
4. 公众/服务号菜单


> 流量方指的是在自己的小程序中加入客户的广告收取广告费用的公司

> 广告主指的是支付广告费用给其他公司，其他公司通过各种方式引导普通用户进入自己的小程序

## 流量方

### 跳转其他小程序

#### 商务

跳转其他小程序需要客户提供他的小程序 `app-id` 小程序路径和 `extra-data`, `extra-data` 一般作为统计不同流量方的标识。（推广ID）

并在自己的小程序管理后台关联客户的 `app-id`

跳转有两种方式：

1. 进入自己的小程序后直接跳转到其他小程序
2. 点击按钮跳转到其他小程序

> 目前可以打开自己的小程序直接跳转到客户的小程序，但是这个功能后续会被微信禁用。


#### 技术实现

##### 直接跳转

> 此接口即将废弃，请使用 `<navigator>` 组件来使用此功能

```js
wx.navigateToMiniProgram({
  appId: '',
  path: '',
  extraData: {
    channel_id: '客户提供的流量方标识'
  },
  success(res) {
    // 打开成功
  }
})
```

[wx.navigateToMiniProgram](https://developers.weixin.qq.com/miniprogram/dev/api/navigateToMiniProgram.html)

##### 点击按钮跳转

```html
<navigator
    target="miniProgram"
    open-type="navigate"
    app-id="小程序APPID"
    path="小程序路径"
    extra-data=""
    version="release"
    >跳转其他小程序</navigator>
```
[navigator](https://developers.weixin.qq.com/miniprogram/dev/component/navigator.html)


### 跳转 APP

#### 商务

跳转 APP 必须通过用户主动点击按钮才能跳转。

与客户的技术对接时请将这个链接发送给客户 https://developers.weixin.qq.com/miniprogram/dev/api/launchApp.html 需要客户的技术人员接入微信 OpenSDK。

#### 技术实现

可尝试在 `app-parameter` 中传递用于统计的流量方标识

```html
<button open-type="launchApp" app-parameter="wechat" binderror="launchAppError">打开APP</button>
```

```js
Page({
    launchAppError: function(e) {
        console.log(e.detail.errMsg)
    }
})
```

[launchApp](https://developers.weixin.qq.com/miniprogram/dev/api/launchApp.html)

### 打开网页

微信小程序官方提供了打开网页的功能，但是只能打开绑定此小程序的域名。


每个小程序只能绑定20个业务域名，业务域名也就是客户的网址。绑定业务域名时需要在微信小程序管理后台下载验证文件，交给广告客户上传到客户的服务器。（客户在安全性和便利性上不一定会愿意），上传成功后即可在小程序中打开客户的网站。但客户的网站中不可以跳转到其他不在业务域名的网站。

> 注意：业务域名一年只能修改50次。所以建议寻找稳定的广告客户。小客户经常更换的不建议合作。

[web-view](https://developers.weixin.qq.com/miniprogram/dev/component/web-view.html)


## 广告主

### 普通链接

找设计人员设计页面，比如：

![](http://effect.admpv.com/turntable/index.png)

然后由技术人员上传到服务器，给商务人员一个在线访问地址。商务将地址给客户。客户引导用户进入页面识别二维码进入微信小程序。

#### 技术实现

通过微信小程序后台 **设置>开发设置>扫普通链接二维码打开小程序** 可以生成二维码，二维码可以加上唯一标识，统计不同流量的访问情况。


### 公众号推文

#### 阅读原文

与普通链接一致，向技术索取链接地址

#### 小程序

微信推文中可插入小程序卡片，点击卡片可直接打开微信小程序。技术需要向商务提供 AppID,路径(path) ，路径上应该带上统计id。

##### 流量方操作流程

1. 流量方时需要在微信小程序管理后台管理关联我们的小程序（我们需要提供小程序 AppID）
2. 关联后需要我们自己在微信小程序管理后台同意关联
3. 创建推文时选择小程序>插入小程序卡片

### APP直接打开小程序


移动应用拉起小程序是指用户可以通过接入该功能的第三方移动应用（APP）跳转至某一微信小程序的指定页面，完成服务后跳回至原移动应用（APP）。

##### 流量方操作流程

[移动应用拉起小程序功能](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=21526646385rK1Bs&token=&lang=zh_CN)

**需要在微信开放平台创建应用后才能实现APP打开小程序 [创建应用](https://open.weixin.qq.com/cgi-bin/frame?t=home/app_tmpl&lang=zh_CN)**

[Android开发示例](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=21526646437Y6nEC&token=&lang=zh_CN)

[iOS开发示例](https://open.weixin.qq.com/cgi-bin/showdocument?action=dir_list&t=resource/res_list&verify=1&id=21526646447MMfXU&token=&lang=zh_CN)

### 公众/服务号菜单

#### 流量方操作流程

登录微信公众号/服务号后台，选择左侧导航自定义菜单。选择对应菜单，菜单内容选择跳转小程序。

未出现我们小程序的，需要关联我们的小程序。我们需要提供 AppID 给流量方关联后需要我们在后台确定绑定。


若作者显示不是Nimo （被转载了），请访问Github原文进行讨论：https://github.com/nimoc/blog/issues/31


<script src="https://utteranc.es/client.js"
        repo="nimoc/blog"
        issue-number="31"
        theme="github-light"
        crossorigin="anonymous"
        async>
</script>