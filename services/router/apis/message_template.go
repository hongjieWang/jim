package apis

import (
	"github.com/jim/services/common"
	"github.com/jim/services/router/store"
	"github.com/kataras/iris/v12"
)

type MessageTemplateHandler struct {
	service store.MessageTemplateService
}

func NewMessageTemplateHandler() *MessageTemplateHandler {
	return &MessageTemplateHandler{service: store.InitMessageTemplateService()}
}

func (h *MessageTemplateHandler) Add(ctx iris.Context) {
	m := new(store.MessageTemplate)
	err := ctx.ReadJSON(m)
	if err != nil {
		common.FailJSON(ctx, iris.StatusBadRequest, err, "Malformed request payload")
		return
	}
	err = h.service.Create(nil, m)
	if err != nil {
		common.InternalServerErrorJSON(ctx, err, "Server was unable to create a movie")
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(m)
}
