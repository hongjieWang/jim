package store

import (
	"context"
	"github.com/jim/services/database"
)

type UserAccount struct {
	Phone string `json:"phone"`
	Mode  []Mode `json:"modes"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Mode struct {
	OpenId   string `json:"open_id"`
	UserId   string `json:"user_id"`
	UnionId  string `json:"union_id"`
	OtherId  string `json:"other_id"`
	ModeType string `json:"mode_type"`
}

// UserAccountService 用户账号配置Service
type UserAccountService interface {
	GetAll(ctx context.Context) ([]UserAccount, error)
	GetUserAccountByPhone(ctx context.Context, phone string) (UserAccount, error)
	Save(ctx context.Context, account *UserAccount) error
	Update(ctx context.Context, phone string, account *UserAccount) error
	Delete(ctx context.Context, phone string) error
}

type userAccountService struct {
}

func InitUserAccountService() UserAccountService {
	return &userAccountService{}
}

var _ UserAccountService = (*userAccountService)(nil)

func (user userAccountService) GetAll(ctx context.Context) ([]UserAccount, error) {
	return nil, nil
}

func (user userAccountService) GetUserAccountByPhone(ctx context.Context, phone string) (UserAccount, error) {
	return UserAccount{}, nil
}
func (user userAccountService) Save(ctx context.Context, account *UserAccount) error {
	_, err := database.UserAccountCollection.InsertOne(ctx, account)
	return err
}
func (user userAccountService) Update(ctx context.Context, phone string, account *UserAccount) error {
	return nil
}
func (user userAccountService) Delete(ctx context.Context, phone string) error {
	return nil
}
