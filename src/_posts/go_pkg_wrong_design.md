----
title: 错误的go包设计
date: 2021-05-15
keywords: go package, 函数签名设计
description: 
tags:
- go
issues: 53
----


# 错误的go包设计


## sql.ErrNoRows

```go
var name string
has := true
err := db.QueryRow("select name from users where id = ?", id).Scan(&name)
if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
        has = false
    } else {
        return err
    }
}
log.Print("name:", name)
return nil
```


scan 的函数签名是
```go
Scan(dest ...interface{}) error
```

应该是

```go
Scan(dest ...interface{}) (has bool, err error)
```


让使用者判断错误是不是 `sql.ErrNoRows` 会让代码变得绕来绕去.

因为这已经是标准库的实现,为了向前兼容函数签名不会变动.

所以我们自己实现 package 时不要这样设计恶心使用者.

但是我们可以自己封装一层,感兴趣的可以关注 https://github.com/goclub/sql

类似的设计还有 `http.ErrNoCookie`,用起来都很操蛋.

##  resp.Body.Close


 


