# 封装 HTTP 接口

> 本文适合有一定前后端数据交互经验的开发人员阅读

### 为什么要封装 HTTP 接口

1. 过滤服务器端数据
2. 统一管理和语义化接口

#### 过滤服务器端数据

例如一个用户的状态分别有 `未激活` `已激活` `禁用`

接口返回的数据可能是

```js
// GET /user
[
    {
        type: 0, // 未激活
        name: 'nimo'

    },
    {
        type: 1, // 已激活
        name: 'tim'
    },
    {
        type: 2, // 禁用
        name: 'wisdom'
    }
]
```

如果直接使用此数据，前端逻辑代码会变得很糟糕。

```js
if (user.type === 2) {
    alert('用户被禁用')
}
```

#### 统一管理和语义化接口

有时候会遇到多个页面请求同一个接口的情况

```js
// 下面这段代码需要出现在 newsAdd.js 和 somePage.js 文件中
$.ajax({
    type: 'post',
    url: '/news',
    dataType: 'json'
}).done(function (res) {
    /** res 数据格式
        {
            status: 'pass',
            newsid: 'egvw423h35hy35hqgwfwsfw2'
        }
        {
            status: 'fail',
            code: 'sensitiveWords'
        }
     */
     var failDict = {
         'sensitiveWords': '包含敏感词',
         'limit': '每分钟只能发布一条新闻'
     }
     if (res.status === 'pass') {

         /* newsAdd.js 只需要弹窗 */
         alert('添加新闻成功')

         /* somePage.js 需要跳转到新闻页 */
         // location.href = '/news?id=' + newsid

     }
     else {
         let failMessage = failDict[res.code]
         if(typeof failMessage === 'undefined') {
             alert('添加失败，请联系管理员！错误代码：' + res.code)
         }
         else {
             alert(failMessage)
         }
     }
})
```

两个文件同时存在一样的代码，会难以维护。但是两个文件对于 `res.status === 'pass'` 情况下的处理方式又不同

### 如何封装

#### 数据过滤

我们期望后端返回的数据是

```js
[
    {
        name: 'nimo',
        // 连注释都不用写了
        type: 'inactive'
    },
    {
        name: 'tim',
        type: 'active'
    },
    {
        name: 'wisdom',
        type: 'disable'
    }
]
```


### 统一管理和语义化接口

封装后可以非常优雅的调用接口

```js
// newsAdd.js
var apiNews = require('./m/api/news')
var sendData = {
    title: 'abc',
    content: 'xxxxoxoxoxoxox'
}
apiNews.post(
    sendData,
    {
        pass: function () {
            alert('添加新闻成功')
        },
        fail: function (failMessage, code) {
            alert(failMessage)
        }
    }
)
```

```js
// somePage.js
var apiNews = require('./m/api/news')
var sendData = {
    title: 'abc',
    content: 'xxxxoxoxoxoxox'
}
apiNews.post(
    sendData,
    {
        pass: function (newsid) {
            alert('添加新闻成功，跳转至新闻页')
            location.href = '/news?id=' + newsid
        },
        fail: function (failMessage, code) {
            alert(failMessage)
        }
    }
)
```

#### proxy 和 dict

--------------------------------------

底层: XHR - ES6 fetch - wx.request
接口封装：jQuery.ajax
项目通用业务逻辑封装：proxy （网络错误 alert 弹窗，或者将错误发送到服务器端,服务器超过规定时间未响应，则弹出错误）
具体接口语义化封装： apiLogin(data, settings) settings.pass settings.error  settings.inactive


##### dict

```js
var dict = require('dict')
dict.addFamily('sms', {
    'login': 1,
    'register': 2
})
dict('sms', 'login') // 1
dict('sms', '1') // login
dict('sms', 1) // login
```
