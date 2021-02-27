----
title: 代码中的监控日志
date: 2021-02-28
keywords: 业务日志,日志规范
description: 通过具体的例子来说明代码中何时写入监控日志，提供思考方法。应对不的情况。
tags:
- 后端
- 编程技巧
issues: 43
----

# 代码中的监控日志

日志有很多种类型，本文所说的代码中的监控日志指的是在某些环境记录日志，并在出现新日志时候及时监控日志内容判断业务是否稳定。


## 监控查询不存在的数据

看看下面的代码：

```js
// 业务逻辑层：获取商品信息
function BusinessLayer_GoodsDetail(goodsID) {
    goods = DataStorage_GoodsByID(goodsID)
    if (goods == null) {
        return {
            message: "商品不存在"
        }
    }
    return {
        title: goods.title,
        price: goods.price,
    }
}
// 数据存储层：获取商品数据
function DataStorage_GoodsByID(goodsID) {
    row = sql("SELCT * FROM goods WHERE id = ? LIMIT 1")
    if row == null {
        // 数据查询不到记录日志
        MonitorLog("remind", "database: goods not found, goodsID:" + goodsID )
        return null
    }
    return row
}
```

请求的 goodsID 对应的数据有可能不存在数据库中的。遇到这种情况需要思考是否是正常情况：

在商品详情的场景下少量的访问不存在的数据是正常的，如果短时间有大量的请求访问不存在的数据则一定是什么地方出现了问题。
例如：

1. 数据被意外删除
2. 有恶意攻击频繁请求不存在的id企图利用缓存穿透攻击服务器，并且缓存层的防御方案失效了
3. 代码中将 orderID 误传成了 goodsID 导致客户端拿到错误的 id 后发送请求。

如果不记录日志则会导致这些错误和攻击造成很明显的破坏时候才**被动发现**。

## 监控日志细节

来看看监控日志的调用代码：

```js
MonitorLog("remind", "database: goods not found, goodsID:" + goodsID )
```

第一个参数表明当前日志是提醒级别，少量的错误可以忽视，当出现大量错误时需分析排查。
可通过 sentry 等成熟的日志平台，在日志平台中管理各种错误的时间周期忽略次数。

第二个参数是详细的错误信息，错误必须带上关键的信息，便于后续分析。如果没有传递 goodsID 则会增加排查时间。

> 虽然一些日志体系可以做到记录当前客户端请求的ID，再通过ID查询相应的请求报文（比如HTTP报文），在报文中直接或间接查找 goodsID。但是这样依然会增加排查时间。

指的注意的是 `MonitorLog` 内部一定要将当前调用堆栈发送给日志系统。像直接输出到服务器日志文件，并且只记录错误消息内容也会增加排查时间。

> 各个语言都会有成熟的日志库，并且也有大量类似 sentry 这样的日志服务。千万不要简单的写入到服务器文件，因为业务出问题之前，你不会去检查不易于阅读日志文件。监控日志的作用不只是在出现问题时用于排查问题，还能提前预知问题。


当我们在一个日UV100万的项目中发现10分钟内容有1万次不存在的数据查询，则能通过监控系统快速发现问题。并分析问题解决问题。（谁都不想问题由客户反馈给客服，客服再反馈给自己）

## 监控的位置

刚才的监控代码是在 `DataStorage_GoodsByID` 中也就是数据逻辑层。而不是 `BusinessLayer_GoodsDetail`业务逻辑层。
原因是当前业务场景下要查询商品在大部分情况一定存在，因为商品 id 是客户端通过其他请求获取到的。

未完待续...
