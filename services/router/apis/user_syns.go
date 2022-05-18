package apis

import (
	"github.com/jim/message/feishu"
	"github.com/jim/services/common"
	"github.com/jim/services/router/store"
	"github.com/kataras/iris/v12"
)

type UserSynHandler struct {
	service store.UserAccountService
}

// NewUserSynHandler 获取用户数据同步
func NewUserSynHandler() *UserSynHandler {
	return &UserSynHandler{service: store.InitUserAccountService()}
}

func (h *UserSynHandler) Syn(ctx iris.Context) {
	app := feishu.Feishu{AppSecret: "TFzUlKx3FJrDFhZdrwpxahXrtpIU31xP", AppId: "cli_a208d1b17ef8900d"}
	deptUsers := app.GetUsersByDept("0")
	for _, item := range deptUsers.Data.Items {
		account := &store.UserAccount{
			Phone: item.Mobile,
			Name:  item.Name,
			Email: item.Email,
			Mode: []store.Mode{
				{
					OpenId:   item.OpenID,
					UnionId:  item.UnionID,
					UserId:   item.UserID,
					ModeType: "Feishu",
				},
			},
		}
		h.service.Save(nil, account)
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(common.Results{Code: iris.StatusOK, Message: "数据同步中..."})
}
