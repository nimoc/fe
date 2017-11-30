# React Component - 内部状态逻辑封装

> 内部状态逻辑封装只适合开发一些简单的页面，或页面中非常简单的模块。一旦涉及到状态要同步到顶层，就不应该使用这种方式。

> 深入讨论 声明式框架的内部状态和外部状态

普通实现

```jsx
<form onSubmit={function () {
    self.setState({loading: true})
    $.ajax({

    }).done(function () {
        self.setState({loading: hidden})
    })
}} >
    ...
    <Loading loading={self.state.loading} >
        <button></button>
    </Loading>
</form>
```

逻辑封装实现

> 适合 loading 状态不与其他操作关联的情况

```jsx
<LoadingLogic render={function (loading) {
    return (
        <form onSubmit={function () {
            loading.show()
            $.ajax({

            }).done(function () {
                loading.hide()
            })
        }} >
            ...
            <Loading loading={loading.value} >
                <button></button>
            </Loading>
        </form>
    )
}} />
```


```jsx
<Free  render={ (item)  => <input name="user" {...item('user')} /> } />
```
