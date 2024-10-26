package user_usercases_test

import (
	"testing"
	"time"

	"github.com/luizrgf2/chat_real_time_go/internal/infra/repositories"
	user_entity "github.com/luizrgf2/chat_real_time_go/internal/user/entities"
	user_errors "github.com/luizrgf2/chat_real_time_go/internal/user/errors"
	user_usercases "github.com/luizrgf2/chat_real_time_go/internal/user/usecases"
	user_usecases_dto "github.com/luizrgf2/chat_real_time_go/internal/user/usecases/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateValidUser(t *testing.T) {
	mockRepo := repositories.UserRepoMock{}

	user := user_usecases_dto.CreateUserInputDTO{
		NameUser: "Luiz",
	}

	expected := user_usecases_dto.CreateUserOutputDTO{
		ID:        1,
		Name:      "Luiz",
		CreatedAt: time.Date(2022, 5, 1, 0, 0, 0, 0, time.Local),
		UpdatedAt: time.Date(2022, 5, 1, 0, 0, 0, 0, time.Local),
	}

	userData := user_entity.UserEntity{
		Id:        expected.ID,
		Name:      expected.Name,
		CreatedAt: time.Date(2022, 5, 1, 0, 0, 0, 0, time.Local),
		UpdateAt:  time.Date(2022, 5, 1, 0, 0, 0, 0, time.Local),
	}

	mockRepo.On("Create", mock.MatchedBy(func(user *user_entity.UserEntity) bool {
		return user.Name == userData.Name && user.Id == userData.Id
	})).Return(&userData, nil)
	mockRepo.On("FindByName", user.NameUser).Return(nil, user_errors.ErrUserNotExists)

	sut := user_usercases.CreateUserUseCase{UserRepo: &mockRepo}

	res, err := sut.Exec(user)

	assert.Nil(t, err)
	assert.Equal(t, expected.Name, res.Name)
	assert.Equal(t, expected.ID, res.ID)
	assert.NotEmpty(t, expected.CreatedAt)
	assert.NotEmpty(t, expected.CreatedAt)
	assert.NotEmpty(t, expected.UpdatedAt)
}

func TestCreateUserAlreadyExists(t *testing.T) {
	mockRepo := repositories.UserRepoMock{}

	user := user_usecases_dto.CreateUserInputDTO{
		NameUser: "Luiz",
	}

	mockRepo.On("FindByName", user.NameUser).Return(nil, user_errors.ErrUserAlreadyExists)

	sut := user_usercases.CreateUserUseCase{UserRepo: &mockRepo}

	res, err := sut.Exec(user)

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.ErrorIs(t, user_errors.ErrUserAlreadyExists, err)
}
