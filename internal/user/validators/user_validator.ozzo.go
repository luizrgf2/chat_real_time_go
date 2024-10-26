package user_validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	user_entity "github.com/luizrgf2/chat_real_time_go/internal/user/entities"
)

type OzzoUserValidator struct{}

func (v *OzzoUserValidator) Validate(user *user_entity.UserEntity) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Name, validation.Required, validation.Length(3, 50)),
	)
}
