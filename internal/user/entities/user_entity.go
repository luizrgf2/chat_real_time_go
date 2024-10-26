package user_entity

import (
	"time"
)

type UserEntity struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdateAt  time.Time
}
