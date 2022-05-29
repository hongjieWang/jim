package apis

import (
	"fmt"
	"github.com/jim/services/common"
	"github.com/jim/services/router/store"
	"github.com/kataras/iris/v12"
)

type MessageHandler struct {
	service store.MessageService
}

func NewMessageHandler() MessageHandler {
	return MessageHandler{
		service: store.InitMessageService(),
	}
}

func (m *MessageHandler) Send(ctx iris.Context) {
	message := new(store.Message)
	err := ctx.ReadJSON(message)
	if err != nil {
		common.FailJSON(ctx, iris.StatusBadRequest, err, "Malformed request payload")
		return
	}
	err = m.service.Send(nil, message)
	fmt.Println(message)
}
