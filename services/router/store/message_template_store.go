package store

import (
	"context"
	"fmt"
	"github.com/jim/services/database"
	"go.mongodb.org/mongo-driver/bson"
)

// MessageTemplate 消息模版
type MessageTemplate struct {
	TemplateNo      string `bson:"template_no"`
	TemplateContent string `bson:"template_content"`
	Mode            string `bson:"mode"`
	MsgType         string `bson:"msg_type"`
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
	fmt.Println(database.MessageTemplateCollection)
	_, err := database.MessageTemplateCollection.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *messageTemplateService) matchTemplateNo(templateNo string) bson.D {
	filter := bson.D{{Key: "template_no", Value: templateNo}}
	return filter
}

func (s *messageTemplateService) GetByTemplateNo(ctx context.Context, templateNo string) (MessageTemplate, error) {
	var messageTemplate MessageTemplate
	database.MessageTemplateCollection.FindOne(ctx, s.matchTemplateNo(templateNo)).Decode(&messageTemplate)
	return messageTemplate, nil
}

func (s *messageTemplateService) Update(ctx context.Context, templateNo string, m MessageTemplate) error {
	_, err := database.MessageTemplateCollection.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *messageTemplateService) Delete(ctx context.Context, id string) error {
	return nil
}
