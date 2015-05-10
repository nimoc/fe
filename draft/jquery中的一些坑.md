jquery中的一些坑
==============================
<!--_PAGEDATA
{
    "title": "jquery中的一些坑",
    "githubissuesid": ,
    "createData": "2015-04-03",
    "keywords": "",
    "description":""
}
_PAGEDATA-->

    $.each(['mail@qq.com','demo@test.com'], function () {
        console.log(this)
    })

    /*
    String {0: "m", 1: "a", 2: "i", 3: "l", 4: "@", 5: "q", 6: "q", 7: ".", 8: "c", 9: "o", 10: "m", length: 11, [[PrimitiveValue]]: "mail@qq.com"}

    String {0: "d", 1: "e", 2: "m", 3: "o", 4: "@", 5: "t", 6: "e", 7: "s", 8: "t", 9: ".", 10: "c", 11: "o", 12: "m", length: 13, [[PrimitiveValue]]: "demo@test.com"}

    返回的不是字符串而是String对象，因为this无法指向字符串。

    需要进行如下操作
    console.log(this.toString() == 'mail@qq.com')  // false
    console.log(this.toString() === 'mail@qq.com') // true
    */

最佳使用方法是

    $.each(['mail@qq.com','demo@test.com'], function (index, value) {
        console.log(value)
    })


https://github.com/nimojs/RainUED/issues/31

2. jQuery 修改 checked 时使用 $(elem).prop("checked")


访问Github原文进行讨论：[https://github.com/nimojs/blog/issues/11](https://github.com/nimojs/blog/issues/)