package store

import (
	"context"
	"github.com/jim/services/database"
	"go.mongodb.org/mongo-driver/mongo"
)

// mongodb数据表
const collection = "message_template"

var c = &mongo.Collection{}
var (
	Client, _ = database.Init("mongodb+srv://julywhj:XXX@cluster0.r1o1v.mongodb.net/test")
)

func init() {
	c = Client.Database("jim").Collection(collection)
}

// MessageTemplate 消息模版
type MessageTemplate struct {
	TemplateNo      string
	TemplateContent string
}

// MessageTemplateService 消息模版Service
type MessageTemplateService interface {
	GetAll(ctx context.Context) ([]MessageTemplate, error)
	GetByTemplateNo(ctx context.Context, templateNo string) (MessageTemplate, error)
	Create(ctx context.Context, m *MessageTemplate) error
	Update(ctx context.Context, templateNo string, m MessageTemplate) error
	Delete(ctx context.Context, id string) error
}

type messageTemplateService struct {
}

var _ MessageTemplateService = (*messageTemplateService)(nil)

func InitMessageTemplateService() MessageTemplateService {
	return &messageTemplateService{}
}

func (s *messageTemplateService) GetAll(ctx context.Context) ([]MessageTemplate, error) {
	return nil, nil
}

func (s *messageTemplateService) Create(ctx context.Context, m *MessageTemplate) error {
	_, err := c.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *messageTemplateService) GetByTemplateNo(ctx context.Context, templateNo string) (MessageTemplate, error) {
	return MessageTemplate{}, nil
}

func (s *messageTemplateService) Update(ctx context.Context, templateNo string, m MessageTemplate) error {
	_, err := c.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *messageTemplateService) Delete(ctx context.Context, id string) error {
	return nil
}
