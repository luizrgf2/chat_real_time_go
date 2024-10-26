package user_interfaces_repo

import user_entity "github.com/luizrgf2/chat_real_time_go/internal/user/entities"

type IUserRepository interface {
	Create(user *user_entity.UserEntity) (*user_entity.UserEntity, error)
	FindOne(id uint) (*user_entity.UserEntity, error)
	FindByName(name string) (*user_entity.UserEntity, error)
}
