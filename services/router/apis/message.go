package apis

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/_examples/tutorial/mongodb/httputil"
)

type Message struct {
	//业务编号
	BusinessNo string
	//消息内容
	Content string
	//消息参数
	Params []string
	//消息接收用户
	UserPhones []string
	//消息模版编号
	TemplateNo string
	//消息发送渠道
	SendChannel []string
}

func (m *Message) Send(ctx iris.Context) {
	message := new(Message)
	err := ctx.ReadJSON(message)
	if err != nil {
		httputil.FailJSON(ctx, iris.StatusBadRequest, err, "Malformed request payload")
		return
	}
	print(message)
}
