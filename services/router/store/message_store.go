package store

import "context"

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
	TemplateNos []string
}

type MessageService interface {
	Send(ctx context.Context, message *Message) error
}

var _ MessageService = (*messageService)(nil)

func InitMessageService() MessageService {
	return &messageService{}
}

type messageService struct {
}

func (s messageService) Send(ctx context.Context, message *Message) error {

	return nil
}
