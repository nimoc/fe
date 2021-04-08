----
title: 原子性
date: 2021-04-08
keywords: 原子性
description: 原子性
tags:
- 后端
- 编程技巧
issues: 44
----

# 原子性

## sql中插入数据

职业生涯中最先遇到的原子性问题一般是SQL INSERT，例如存在如下两张表

```
user
id,name

user_address
user_id,address
```

创建用户时提交的信息是

```js
{
    "name": "nimoc",
    "address": "A省B市C区D街道1号"
}
```

根据提交的数据需要执行2条SQL

```
userID = sql("INSERT INTO user (id, name) VALUES(uuid(), 'nimoc'")
sql("INSERT INTO user_address (id, name) VALUES(userID, 'A省B市C区D街道1号')")
```

如果 `INSERT INTO user` 执行成功，但 `INSERT INTO user_address` 执行失败会导致数据不完整（数据不一致）。

> 宕机，网络异常，2个sql之间有代码报错都会导致第2个sql没有执行或执行失败。

为了确保不出现这种情况，需要让2次插入是原子性操作。通过事务实现即可。

```
sql("BEGIN")
userID = sql("INSERT INTO user (id, name) VALUES(uuid(), 'nimoc'")
sql("INSERT INTO user_address (id, name) VALUES(userID, 'A省B市C区D街道1号')")
sql("COMMIT")
```

**sql保证了2个操作要么都执行完成，要么都不执行。符合了原子性，保障了数据一致**


## redis lua 脚本

新手常犯的错误是 redis 的 `get` `set` 命令一起用。

```js
function uv(userID) {
    visited = redis("GET", userID) != nil
    if (!visited) {
        // 标记用户访问过
        redis("SET", userID, "1")
        // 递增 uv 
        redis("INCR", "uv", 1)
    }
}
function queryUV() {
    return redis("GET", "uv")
}
```

这段代码有两个问题：
1. SET 之后 不一定能执行 INCR
2. A线程执行了 GET 之后，B线程也执行了 GET,他们获取到的结果都是 nil，都执行了 SET和INCR 操作。导致了数据不一致。多产生了一个UV.

可通过 redis lua 脚本让三个操作变成原子性操作。


```js
function uv(userID) {
    redisEval(`
        local visted = redis.call("GET", KEYS[1]) != nil
        if visted then 
            redis.call("SET", KEYS[1])
            redis.call("INCR", KEYS[2])
            return 1
        end
        return 0
    `, userID, "uv")
}
function queryUV() {
    return redis("GET", "uv")
}
```


注意 redis lua 脚本的原子性跟 sql 事务原子性不一样，redis lua 脚本内如果命令执行错误，是不会自动回滚的。
你需要确保命令语法不要出现错误，这样就能保证命令一定会执行。

**redis lua 脚本保证了3个命令一起执行，消除执行间隙。避免了数据竞争，达到了并发安全。**

> 这里使用lua脚本实现UV的统计只是为了说明原子性。
> 日常工作中 uv 这种场景用 HyperLogLog 或 Sets 实现能同时解决上述两个问题。

```js
function uv() {
    visted = redis("SADD", "visited", userID) == 0
}
function queryUV() {
    return redis("SCARD", "uv")
}
```

## 不是每个场景都需要达到原子性

考虑如下场景：

```
// 检查验证码
function checkCaptcha(captcha, sessionID) {
    key = "captcha:" + sessionID
    data = redis("GET", key)
    redis("DEL", key) // 读取 captcha 后立即删除，防止恶意穷举
    if (data == captcha) {
        return true
    }
    return false
}
```

GET 和 DEL 不是原子性操作，但是不会造成数据不一致。因为 如果 GET 执行了但是 DEL 没有执行，不会对数据造成任何改动。

分析不满足原子性时候要明确的写出或说出来如果不满足原子性会造成什么样的BUG。尝试明确的表述出会造成的 BUG 能减少一些非必要的原子性操作。

## 不同系统之间的原子性

考虑如下场景:

```js
function sendRedpack(accountID, openid, amount) {
    sql("begin")
    // CAS乐观锁扣除余额
    affected = sql("UPDATE account_finance SET balance = balance - $amount WHERE blance >= $amount AND account_id = $accountID")
    if (affected == 0) {
        sql("rollback")
        return
    }
    sql("INSERT INTO red_pack_record (account_id, openid, amount) VALUES($accountID, $openid, $amount)") 
    ok = httpRequest("https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", {...})
    if (ok == false) {
        sql("rollback")
        return
    } 
    sql("commit")
}
```

上面代码存在如下问题：


长事务，httpRequest的时间是不可控的，可能会导致事务长时间不结束（事务）。长事务会导致系统能支持的并发量下降。UPDATE 后 account_finance 中  accountID 这一条数据会被锁定。

虽然 UPDATE 和 INSERT是原子性，但是 httpRequest 与 SQL 操作不是原子性， sql commit 有可能因为网络原因失败。这就导致红包发出去了，但是钱没扣。

这就出现了**数据不一致**的问题。

---

通过本地任务/消息表可以解决不同系统之间的原子性。（分布式事务）

实现思路是

1. httpRequest 之前在事务中在任务表新增一个发放任务，httpRequest在事务 commit 之后执行。
2. httpRequest 执行完成后将发放任务标记完成
3. 如果 httpRequest 没有执行。由 cron 检查未完成的任务，重试发放。
4. 任务的执行需要是幂等性。微信的发放红包接口通过 order_no 实现了幂等，同一个 order_no 无论请求多少次接口，只会发放一次红包。

```js
function sendRedpack(accountID, openid, amount) {
    sql("begin")
    affected = sql("UPDATE account_finance SET balance = balance - $amount WHERE blance >= $amount AND account_id = $accountID")
    if (affected == 0) {
        sql("rollback")
        return
    }
    orderNo = randomString(28)
    sql("INSERT INTO red_pack_record (order_no, account_id, openid, amount) VALUES($orderNo, $accountID, $openid, $amount)")
    sendRedPackTaskID = sql("INSERT INTO send_red_pack_task (order_no, openid, amount, status) VALUES($orderNo, $openid, $amount, 'processing')") 
    sql("commit")
    ok = httpRequest("https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", {...})
    if (ok) {
        sql("UPDATE send_red_pack_task SET status= 'done' WHERE id = $sendRedPackTaskID")
    } else {
        // 处理金额退回，或将红包发放到用户的余额中。
        sql("UPDATE send_red_pack_task SET status= 'cancel' WHERE id = $sendRedPackTaskID")
    }
}
// 使用定时任务去对未发放的红包进行重试操作
function cronCompensationSendRedpack() {
    oneMinuteAgo = now() - minute()
    ownerKey = uuid()
    data = sql("SELECT * FROM send_red_pack_task WHERE status = 'processing' AND created_at < $oneMinuteAgo limit 100")
    for (item in data) {
        ok = httpRequest("https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", {item})
        if (ok) {
            sql("UPDATE send_red_pack_task SET status= 'done' WHERE id = $item.id")
        } else {
            // 处理金额退回，或将红包发放到用户的余额中。
        }
    } 
}
```

本地任务表的数据是在事务中插入的，确保了 扣钱->新增发放记录->新增发放任务 是原子性的。如果 ` sql("commit")` 失败了，一分钟后，`cronCompensationSendRedpack` 会重试未完成的任务。
不满足原子性会导致数据不一致，本地任务表配合补偿机制以此达到数据**最终一致性**。

未完待续...