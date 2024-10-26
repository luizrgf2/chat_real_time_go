package user_usecases_dto

import "time"

type CreateUserInputDTO struct {
	NameUser string `json:"nameUser"`
}

type CreateUserOutputDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
