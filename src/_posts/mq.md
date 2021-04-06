# 从 sql 出发认识消息队列

文章思路：


```sql
INSERT INTO task (status,body) VALUES('ready', '...')
```
 发布消息（发布任务）

---

```sql
UPDATE task SET status = unacked, processor_id = $processor_id WHERE status = ready
```

配合 `affected` 并发安全的标记数据

---

```sql
SELECT id, body FROM task WHERE status = unacked processor_id = $processor_id
````

读取标记后的数据

---

```sql
UPDATE task SET status = done, processor_id = $processor_id WHERE status = ready
```

 配合 `affected` 标记数据已处理完毕

---

```go
UPDATE task SET status = ready, processor_id = '' WHERE status = unacked AND created_at < $tenMinutesAgo
```  

将十分钟前标记正在处理但未完成的消息（任务）重新标记为待处理。(这里的十分钟可以根据业务情况修改)


