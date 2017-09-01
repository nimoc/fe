# redux-action

## file

```shell
view = state + props + jsx
state = getStore()
store = Server data + UI status
# Server data
    /store/sender/index.js
# UI status
    /view/sender/store/index.js
```

## action

```shell
Server action: "ADD_SENDER"
${action_name}
UI action:  "sender_START_LOADING" | "sender_END_LOADING"
${view_dir}_${action_name}
```
伪代码
```js
"sender_START_LOADING"
$.ajax({
    ...
    done: function (res) {
        if (res.status === 'success') {
            "ADD_SENDER"
        }
    },
    always: function () {
        "sender_END_LOADING"
    }
})
```
