package apis

import (
	"github.com/jim/services/common"
	"github.com/jim/services/router/store"
	"github.com/kataras/iris/v12"
)

// BusinessHandler 业务出来Handler
type BusinessHandler struct {
	service store.BusinessService
}

// NewBusinessHandler 构造业务Handler
func NewBusinessHandler() *BusinessHandler {
	return &BusinessHandler{service: store.InitBusinessNoServiceService()}
}

func (b *BusinessHandler) Add(ctx iris.Context) {
	business := new(store.Business)
	err := ctx.ReadJSON(business)
	if err != nil {
		common.FailJSON(ctx, iris.StatusBadRequest, err, "Malformed request payload")
		return
	}
	err = b.service.Create(nil, business)
	if err != nil {
		common.InternalServerErrorJSON(ctx, err, "Server was unable to create a movie")
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(business)
}

func (b *BusinessHandler) Get(ctx iris.Context) {
	all, err := b.service.GetAll(nil)
	if err != nil {
		common.InternalServerErrorJSON(ctx, err, "Server was unable to create a movie")
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(all)
}
