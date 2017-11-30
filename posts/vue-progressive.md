# Web管理后台解决方案


## 业务场景和需求

每个项目都需要管理后台。前端写好基础的头部、列表、表单样式，由后端复制代码组合出后台功能。


常见的例子就是后端自行复制 [Bootstrap](http://www.bootcss.com/) 代码，组成后台功能。

---


> 某些项目客户也会登录后台，需要更好看的界面。这样才能让客户买单。

**难看的界面**
![](http://pic19.nipic.com/20120217/2724388_110724230108_2.jpg)

> 有些管理操作复杂，需要交互更好的表单控件。比如树形选项管理用户权限。城市级联选择。

**树形选项**
![](https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1511946714735&di=3690be12c6ee15419411d319d51452b6&imgtype=0&src=http%3A%2F%2Fwww.gcpowertools.com.cn%2Fproducts%2Fimg%2Fc1%2Fasp%2FTreeViewControl_Checkbox.png)

面对这种情况我们可以这么做：

1. 由前端开发一套好看的界面和控件供后端调用
2. 找一个专门开发管理后台的开源框架

无论是自主研发还是使用第三方框架，这些代码最终都是由后端同学去使用。那么需要达到以下几点：

1. 调用简单，不需要太多复杂的参数。只在需要复杂功能时才需要很多配置
2. 功能强大，不能只是在外观上比原生的表单要好。而是有原生表单没有的功能。

第三方很多框架都是需要JS调用才能生效的比如：

```html
<div id="some"></div>
<script>
new Select({
    el: '#some',
    // name 用于生成隐藏的表单项，有些第三方框架都不支持生成隐藏的表单项
    name: 'user'
    options: [
        {label: '张三', value: '3tfw3gfwef'},
        {label: '李四', value: 'egf2efewff'},
        {label: '王五', value: '3gt23g23ft'}
    ]
})
</script>
```

后端同事使用时非常麻烦，因为很简单的表单控件都需要写很多JS。

如果是我们自己封装的就可以让后端同事还是写原生的 `<select>`，前端框架会自动替换成更漂亮的 select。

```html
<select name="city" class="v-select" >
    <option value="value1">text1</option>
    <option value="value2">text2</option>
</select>
```

> 所有 class属性包含 `v-select` 的都会被替换成更好看的 select

如果要使用级联选择，也可以封装出适合后端调用的组件。

**级联选择**

![](http://assets.jq22.com/plugin/pc-7730d36a-610a-11e4-b102-00163e001348.png)


```js
<script id="data" type="text/json" >
[
    {
        "label": "北京",
        "value": "1",
        "children": [
            {
                "label": "朝阳区",
                "value":"1-1",
                "children": [
                    {
                        "label": "黄泉路",
                        "value": "1-1-1"
                    }
                ]
            }
        ]
    },
    {
        "label": "上海",
        "value": "2",
        "children": [
            ...
        ]
    }
]
</script>
<span
    data-cascader-name="city"
    data-cascader-width="200"
    data-cascader-options='#data'
    data-cascader-placeholder="请选择城市"
></span>
```


即使调用如此简单可还是会遇到问题：

> 多个表单控件需要做关联，A变动后B也跟着变动。

有些团队遇到这种需求的时候可能就会选择让前端加入开发，承担主要的界面开发任务。

欧耶欧耶，界面也好看了，交互体验也更好了。但是开发成本增加了。

## Vue 闪亮登场

直接上代码

```html
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/iview/dist/iview.min.js"></script>
<link rel="stylesheet" href="https://unpkg.com/iview/dist/styles/iview.css">

<div id="app">
    <form>
        <date-picker name="date" type="date"></date-picker>
        <date-picker name="date2" type="date" value="2017-11-12" ></date-picker>
        <i-button html-type="submit" >提交</i-button>
    </form>
</div>
<script>new Vue({el: '#app'})</script>
```
[在线预览](https://codepen.io/nimojs/pen/WXaagG)
![](https://user-images.githubusercontent.com/3949015/33368402-9657923a-d52c-11e7-94dd-c50099a07553.png)


1. 引入所需 css js 文件  
2. 将 `<date-picker type="date"></date-picker>` 代码存放在 `<div id="app"></div>` 中  
3. 初始化 `new Vue({el: '#app'})`  

配置 `name` 后页面会生成 `<input name="date" />` 选择日期后会同步值到 `input`。通过 `<form />` 包裹 `<date-picker>` 表单提交后就能将日期提交给服务器。

---

虽然 `<span data-cascader-name="city" ... >` 这种data-api 的方式也能实现相同功能，但是难以实现多个控件关联修改。比如：

开启高级选项后，出现选择用户控件。

```html
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/iview/dist/iview.min.js"></script>
<link rel="stylesheet" href="https://unpkg.com/iview/dist/styles/iview.css">

<div id="app" style="padding:20px;" >
    <form>
        高级选项：
        <!-- v-model 可将组件操作绑定到 new Vue({}) 配置的 data 中 -->
        <i-switch name="options" v-model="options" ></i-switch>

        <i-button html-type="submit" >提交</i-button>

        <!-- v-if="options" 可以判断 options 为 true 时显示 -->
        <div v-if="options"style="padding-top:10px;" >
            <i-select name="options_admin" v-model="options_admin" style="width:200px;" >
                <i-option value="a" >
                    张三
                </i-option>
                <i-option value="b" >
                    李四
                </i-option>
            </i-select>
      </div>
    </form>
</div>
<script>
new Vue({
    el: '#app',
    data: function () {
        return {
            options: false,
            options_admin: ''
        }
    }
})
</script>
```
[在线预览](https://codepen.io/nimojs/pen/YEJRdG)

![](https://user-images.githubusercontent.com/3949015/33370036-4c5f1ec8-d531-11e7-8a54-f1574ac6640e.png)


这种关联修改的功能一般由前端开发，但后端也可以自己先尝试参考示例代码自行开发。

使用 Vue 以后，后端调用代码变得简单。在遇到后端复制代码做不了的页面功能时，由前端参与开发。

即使管理后台变得越来越复杂，也没关系。因为 Vue 完全能开发大型应用。

> Vue 的组件比 jQuery 功能更强大，不需要编译，拿来即用。前后端都可以参与开发。

## v-admin

> 团队可以根据业务场景和需求，选择一个 Vue 组件库写上示例文档给后端同事使用。降低管理后台的开发成本。

虽然直接 Vue 组件库能搭建交互非常强的界面，但是有些代码还是需要团队前端去开发的。比如：

1. 公用导航。
2. 表单提交，ajax 提交表单内容。
3. 点击按钮，ajax 提交指定信息。
4. 点击按钮，出现选项 ajax 提交选项。

有了这些功能后端就能更快速的搭建管理后台。

> 因为我们团队2年前就在开始做 data-api 式的后台框架，这些功能是使用频繁且能提高页面开发效率。
