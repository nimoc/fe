package drafts_test

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"testing"
	assert "github.com/stretchr/testify/assert"
)

var _= `
----

title: 在Go中自定义枚举类型
date: 2020-08-01
tags: 后端,golang
issues: 42

----
`
var _= `
# 在Go中自定义枚举类型

`

var _ = `
## Go语言特性
使用枚举的场景一般是函数参数或用来表示（状态 种类 模式） 等信息。
直接使用字符串或数字会让代码不易于维护
`
type Event struct {
	Title string
	Type string // 只允许2中活动类型 "normal" "hot"
}
func (event Event) IsNormal() bool {
	return event.Type == "normal"
}
func NewEvent(eventType string, title string) Event {
	return Event{
		Title: title,
		Type: eventType,
	}
}
func ExampleEvent() {
	// 普通的读书会
	NewEvent("normal", "Reading party")
	// 热门的二手交易市场
	NewEvent("hot", "Flea market")
	// 普通的相亲活动（写错了nromal，应该是 normal）
	blindDate := NewEvent("nromal", "Blind date")
	// 因为拼错了 normal 导致 IsNormal 的判断错误
	blindDate.IsNormal() // false
	// 如果使用数字，那会更糟糕
	/*
		type Event struct {
			Title string
			Type int
		}
		NewEvent(0, "Reading party")
		NewEvent(1, "Flea market")
		NewEvent(1, "Blind date")
	*/
}

var _ = `
使用原始类型（字符串 数字）来传递数据容易出错，代码可维护性不高。

而 go 中没有直接提供枚举类型，在很多go代码中可以看到包装int类型来模拟enum的例子。
`
type NewsType uint8
const (
	NewsTypeNormal NewsType = iota
	NewsTypeHot
)
var _ = `
使用 在 const 中使用 iota  后等同于
	const (
		Normal NewsType = 0
		Hot    NewsType = 1
	)
`

type News struct {
	Title string
	Type NewsType
}
func (news News) IsNormal() bool {
	return news.Type == NewsTypeNormal
}
func NewNews(newsType NewsType, title string) News {
	return News{
		Title: title,
		Type: newsType,
	}
}
func ExampleNews() {
	NewNews(NewsTypeNormal, "Thinking in go  has been published")
	NewNews(NewsTypeHot, "Dogs bury shit")

	// 拼写错 Normal 会编译失败 Unresolve reference 'NewsTypeNromal'
	// NewNews(NewsTypeNromal, "Men give birth to children")

	// 注意硬编码 NewsType 的原始类型 int 时候能通过编译，
	// 我们需要避免这种硬编码传参。使用 NewNews(NewsTypeHot,"") 而不是 NewNews(1, "")
	NewNews(1, "")

	// 而通过传递变量不会编译通过，这是go的语言类型系统的特性。
	// hot := 1
	// NewNews(hot, "")
}

var _= `
1. 只在程序内部使用的enum可使用 uint8
2. 逻辑层代码可以使用 string,因为API接口中使用数字对客户端不友好。
3. 数据库中依然可以使用 tinyint 存储 enum （减少存储空间，便于提高查询速度。）

`
type GoodsType string
const (
	GoodsTypeNormal GoodsType = "normal"
	GoodsTypeHot GoodsType = "hot"
)
type Goods struct {
	Type GoodsType
}
type RequestGoods struct {
	Type GoodsType
}

func TestJSON(t *testing.T) {
	{
		goods := Goods{}
		err := json.Unmarshal([]byte(`{"type":"normal"}`), &goods) ; if err != nil {panic(err)}

		assert.True(t, goods.Type == GoodsType("normal"))
		assert.True(t, goods.Type == GoodsTypeNormal)
		assert.False(t, goods.Type != "normal")
	}
	// json.Unmarshal 能解析是因为 json 库通过反射识别了 Kind 。
	assert.Equal(t, reflect.ValueOf(GoodsTypeNormal).Kind(), reflect.String)
}

var _ = `
## 集中管理 

enum 类型还可以做一个关联性更强的设计
`

type MessageType string
func (MessageType) Enum() (enum struct {
	Normal MessageType
	Hot MessageType
}) {
	enum.Normal = "normal"
	enum.Hot = "hot"
	return
}
func (t MessageType) String() string { return string(t) }
type Message struct {
	Type MessageType
	Content string
}

func ExampleMessage() {
	msg := Message{
		Type: Message{}.Type.Enum().Hot,
		Content: "Project 1.0 release",
	}
	msg.Type = msg.Type.Enum().Normal
}

var _ = `
初看可能觉得 ^Message{}.Type.Enum().Hot^ 和 ^msg.Type.Enum().Normal^ 略显啰嗦，不用担心这一点。
^Type.Enum()^ 方法加强了枚举类型与值的关联性，长远看能提高开发效率。

随着需求的增加我们可能会需要将 ^msg.Type^ 转换为 string 类型，这种情况下使用

	var msgType string	
	msgType = string(msg.Type)

这不是最好的选择，因为不了解 MessageType 的情况下使用强制类型转换会出错。
稳妥的做法是给 MessageType 增加 String() 方法

	func (t MessageType) String() string { return string(t) }

	var msgType string	
	msgType = msg.Type.String()

使用 msg.Type 时候会自觉地查看 msg.Type 下面的方法，看见 String() 时就能明白这是 MessageType 实现者提供的类型转换方法。
此处的 String() 是"实现代替约定"的一种实践。 
`


var  _ = `
在 switch 章节说明过，要警惕不完整的 switch。而 Enum 与 switch 一定是要搭配使用的。故此需要给 MessageType 增加 Switch 方法
`

func (t MessageType) Switch(Normal func(_normal int), Hot func(_hot bool)) error {
	enum := t.Enum()
	switch t {
	default:
		return errors.New("MessageType value error, (" + t.String() + ")")
	case enum.Normal:
		Normal(0)
		return nil
	case enum.Hot:
		Hot(false)
		return nil
	}
}

func ExampleMessageTypeSwitch () {
	msg := Message{
		Type: Message{}.Type.Enum().Normal,
		Content: "OMG",
	}
	err := msg.Type.Switch(func(_normal int) {
		log.Print(msg.Content)
	}, func(_hot bool) {
		log.Print("HOT: " + msg.Content)
	}) ; if err != nil {
		log.Print(err)
	}
}

var _ = `
此时实现的是 ^const MessageTypeNormal MessageType = "normal"^ ，而不是 ^func(t MessageType) Enum()^。

func MessageTypeSwitch(t MessageType, Normal func(_normal int), Hot func(_hot bool)){
	// ...
}

为了将 enum 的相关实现都"聚合"起来，请务必实现 ^value.Enum() (enum struct{...})^ ^value.Switch(func ...)^
`