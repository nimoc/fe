# JavaScript异步编程

异步编程的代表是回调函数
<!--
{
    markrun_last_run: false
}
-->

````html
<!-- 加载 HTTP 请求库 -->
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
````

````js
// 在控制台加载 HTTP 请求库
var axiosScript = document.createElement('script')
axiosScript.src = "https://unpkg.com/axios/dist/axios.min.js"
document.body.appendChild(axiosScript)
````

````js
var url = 'https://api.github.com/repos/nimojs/github-comments/issues/7/comments?per_page=100'
// 获取 github issues 评论
axios.get(url).then(function (res) {
    console.log('data', res.data)
})
````

[axios](https://github.com/axios/axios)


## 获取评论后查询第一个回复的用户信息 (getUserInfo)

**使用回调函数实现**

````js
var url = 'https://api.github.com/repos/nimojs/github-comments/issues/7/comments?per_page=100'

axios.get(url).then(function (res) {
    var userName = res.data[0].user.login
    axios.get(`https://api.github.com/users/${userName}`).then(function (res) {
        console.log('callback getUserInfo', res.data)
    })
})
````

**使用 async 函数实现**

````js
var url = 'https://api.github.com/repos/nimojs/github-comments/issues/7/comments?per_page=100'
;(async function () {
    // axios.get(url).then(function(res){}) 中的 res 会赋值给 comments
    var comments = await axios.get(url)
    comments = comments.data
    var userName = comments[0].user.login
    var res = await axios.get(`https://api.github.com/users/${userName}`)
    console.log('async getUserInfo', res.data)
})()
````

> async 中不需要写任何回调函数且结构清晰


## 获取评论后逐个获取用户信息 (dataMap)

**使用回调函数实现**

````js
var url = 'https://api.github.com/repos/nimojs/github-comments/issues/7/comments?per_page=100'
var dataMap = {}
axios.get(url).then(function (res) {
    var users = []
    res.data.forEach(function (item) {
        var userName = item.user.login
        users.push(userName)
    })    
    var index = 0
    function delayEach (index, callback) {
        var userName = users[index]
        axios.get(`https://api.github.com/users/${userName}`).then(function (res) {
            dataMap[userName] = res.data
            if (index < users.length - 1) {
                delayEach(index + 1, callback)
            }
            else {
                callback()
            }
        })
    }
    delayEach(index, function callback () {
        // 返回所有信息
        console.log('callback dataMap', dataMap)
    })
})
````

**使用 async 函数实现**

````js
;(async function () {
    var dataMap = {}
    var url = 'https://api.github.com/repos/nimojs/github-comments/issues/7/comments?per_page=100'
    var comments = await axios.get(url)
    comments = comments.data

    for(var index = 0; index<comments.length; index++) {
        var userName = comments[index].user.login
        var user = await axios.get(`https://api.github.com/users/${userName}`)
        user = user.data
        dataMap[userName] = user
    }
    // 返回所有信息
    console.log('async dataMap', dataMap)
})()
````

> 异步操作特别多的情况下 async 函数的优势更能体现出来

## axios 中的 then 是什么

`then` 是 `Promise`对象的方法用于配置成功后的回调函数。

[Promise 简介](https://github.com/web-action/es-action/blob/master/es6/Promise.md)

`axios` 是基于 `Promise` 实现的 HTTP 请求。请求成功后触发 `.then(fn)` 中配置的函数。

如果你不了解 `Promise`。暂时只需要知道使用 `.then(fn)` 配置异步成功处理函数，使用 `.catch(fn)` 配置异步失败处理函数

````js
// 因为服务器返回 404，所以会执行 catch 配置的函数
axios.get('https://github.com/onface/null')
    .then(function (res) {
        console.log('then', res)
    })
    .catch(function (error) {
        console.log('catch', error)
    })
````

**async 函数与 Promise 配合使用最佳**

## Generator

`Generator` 也是 JavaScript 异步编程的一种方案，但是现在建议使用 async 函数。如果你想了解 `Generator` 可以阅读 [Generator 简介](https://github.com/web-action/es-action/blob/master/es6/Generator.md)。


## Promise Generator Async


`Promise` 统一了回调函数的写法，虽然还需要配置回调函数，但是不会陷入[回调地狱](http://callbackhell.com/)。`Generator` `Async` 搭配 `Promise` 才能用的更方便。


建议直接使用 `Async` ，`Generator` 可以在用熟练 `Async` 再去了解。自行决定如何使用它们。
