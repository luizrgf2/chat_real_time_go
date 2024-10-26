package repositories

import (
	user_entity "github.com/luizrgf2/chat_real_time_go/internal/user/entities"
	"github.com/stretchr/testify/mock"
)

type UserRepoMock struct{ mock.Mock }

func (u *UserRepoMock) Create(user *user_entity.UserEntity) (*user_entity.UserEntity, error) {
	args := u.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user_entity.UserEntity), args.Error(1)
}

func (u *UserRepoMock) FindOne(id uint) (*user_entity.UserEntity, error) {
	args := u.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user_entity.UserEntity), args.Error(1)

}

func (u *UserRepoMock) FindByName(name string) (*user_entity.UserEntity, error) {
	args := u.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*user_entity.UserEntity), args.Error(1)
}
