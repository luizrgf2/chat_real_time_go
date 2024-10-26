package user_service_test

import (
	"strings"
	"testing"

	user_service "github.com/luizrgf2/chat_real_time_go/internal/user/services"
)

func TestCreateValidUser(t *testing.T) {
	userService := user_service.NewUserService()
	user, err := userService.CreateUser("Luiz")
	if err != nil {
		t.Error(err)
	}
	if user.Name != "Luiz" {
		t.Fatal("Error name expected", "Luiz")
	}
}

func TestCreateInvalidUserWithInvalidName(t *testing.T) {
	userService := user_service.NewUserService()
	user, err := userService.CreateUser("Lu")
	if user != nil {
		t.Error(err)
	}
	if !strings.Contains(err.Error(), "Name: the length must be between 3 and 50") {
		t.Error("Error expected", "Name: the length must be between 3 and 50", "error finded", err.Error())
	}
}
