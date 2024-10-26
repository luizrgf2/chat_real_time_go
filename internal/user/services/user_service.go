package user_service

import (
	"time"

	"github.com/luizrgf2/chat_real_time_go/internal/helpers"
	user_entity "github.com/luizrgf2/chat_real_time_go/internal/user/entities"
	user_validator "github.com/luizrgf2/chat_real_time_go/internal/user/validators"
)

type UserService struct {
	validator helpers.Validator[user_entity.UserEntity]
}

func NewUserService() UserService {
	return UserService{
		validator: &user_validator.OzzoUserValidator{},
	}
}

func (u *UserService) CreateUser(Name string) (*user_entity.UserEntity, error) {
	user := user_entity.UserEntity{
		Id:        1,
		Name:      Name,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
	err := u.validator.Validate(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
