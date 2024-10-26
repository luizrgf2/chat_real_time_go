package user_usercases

import (
	"errors"

	user_entity "github.com/luizrgf2/chat_real_time_go/internal/user/entities"
	user_errors "github.com/luizrgf2/chat_real_time_go/internal/user/errors"
	user_interfaces_repo "github.com/luizrgf2/chat_real_time_go/internal/user/interfaces/repository"
	user_service "github.com/luizrgf2/chat_real_time_go/internal/user/services"
	user_usecases_dto "github.com/luizrgf2/chat_real_time_go/internal/user/usecases/dto"
)

type CreateUserUseCase struct {
	UserRepo user_interfaces_repo.IUserRepository
}

func (c *CreateUserUseCase) checkIfUserExists(userName string) error {
	user, err := c.UserRepo.FindByName(userName)
	if err != nil {
		if errors.Is(err, user_errors.ErrUserNotExists) {
			return nil
		}
		return err
	}
	if user != nil {
		return user_errors.ErrUserAlreadyExists
	}
	return nil
}

func (c *CreateUserUseCase) saveUser(userData user_entity.UserEntity) (*user_entity.UserEntity, error) {
	userSaved, err := c.UserRepo.Create(&userData)
	if err != nil {
		return nil, err
	}
	return userSaved, nil
}

func (c *CreateUserUseCase) createUserEntity(input user_usecases_dto.CreateUserInputDTO) (*user_entity.UserEntity, error) {
	userService := user_service.NewUserService()
	user, err := userService.CreateUser(input.NameUser)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *CreateUserUseCase) Exec(input user_usecases_dto.CreateUserInputDTO) (*user_usecases_dto.CreateUserOutputDTO, error) {
	userData, err := c.createUserEntity(input)
	if err != nil {
		return nil, err
	}

	if err := c.checkIfUserExists(userData.Name); err != nil {
		return nil, err
	}

	savedUsed, err := c.saveUser(*userData)
	if err != nil {
		return nil, err
	}

	return &user_usecases_dto.CreateUserOutputDTO{
		ID:        savedUsed.Id,
		Name:      savedUsed.Name,
		CreatedAt: savedUsed.CreatedAt,
		UpdatedAt: savedUsed.UpdateAt,
	}, nil
}
