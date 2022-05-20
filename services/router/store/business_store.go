package store

import (
	"context"
	"fmt"
	"github.com/jim/services/database"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

type Business struct {
	BusinessNo        string            `bson:"business_no"`
	MessageTemplates  []MessageTemplate `bson:"message_templates"`
	MessageTemplateNo []string          `bson:"message_template_no"`
	BusinessName      string            `bson:"business_name"`
	DeliverAfter      int64             `bson:"deliver_after"`
}

// BusinessService 业务管理Service
type BusinessService interface {
	GetAll(ctx context.Context) ([]Business, error)
	GetByTemplateNo(ctx context.Context, businessNo string) (Business, error)
	Create(ctx context.Context, m *Business) error
	Update(ctx context.Context, businessNo string, m Business) error
	Delete(ctx context.Context, id string) error
}

type businessService struct {
	messageTemplateService MessageTemplateService
}

var _ BusinessService = (*businessService)(nil)

func InitBusinessNoServiceService() BusinessService {
	return &businessService{
		messageTemplateService: InitMessageTemplateService(),
	}
}

func (s *businessService) GetAll(ctx context.Context) ([]Business, error) {
	var bu []Business
	cursor, err := database.BusinessCollection.Find(ctx, bson.D{})
	if err = cursor.All(context.TODO(), &bu); err != nil {
		panic(err)
	}
	return bu, nil
}

func (s *businessService) Create(ctx context.Context, m *Business) error {
	templateNos := m.MessageTemplateNo
	for _, no := range templateNos {
		fmt.Println(no)
		templateNo, _ := s.messageTemplateService.GetByTemplateNo(ctx, no)
		if strings.EqualFold(templateNo.TemplateContent, "") {
			continue
		}
		templates := append(m.MessageTemplates, templateNo)
		m.MessageTemplates = templates
	}
	_, err := database.BusinessCollection.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *businessService) GetByTemplateNo(ctx context.Context, templateNo string) (Business, error) {
	return Business{}, nil
}

func (s *businessService) Update(ctx context.Context, templateNo string, m Business) error {
	_, err := database.BusinessCollection.InsertOne(ctx, m)
	if err != nil {
		return err
	}
	return nil
}

func (s *businessService) Delete(ctx context.Context, id string) error {
	return nil
}
