package user_service_test

import (
	"testing"

	user_service "github.com/luizrgf2/chat_real_time_go/internal/user/services"
	assert "github.com/stretchr/testify/assert"
)

func TestCreateValidUser(t *testing.T) {
	userService := user_service.NewUserService()
	user, err := userService.CreateUser("Luiz")
	if err != nil {
		t.Error(err)
	}
	assert.Nil(t, err)
	assert.Equal(t, "Luiz", user.Name)
}

func TestCreateInvalidUserWithInvalidName(t *testing.T) {
	userService := user_service.NewUserService()
	user, err := userService.CreateUser("Lu")
	if user != nil {
		t.Error(err)
	}

	assert.Contains(t, err.Error(), "Name: the length must be between 3 and 50")
}
