package ticket_validator

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	ticket_entity "github.com/luizrgf2/chat_real_time_go/internal/ticket/entities"
)

type OzzoTicketValidator struct{}

func (v *OzzoTicketValidator) Validate(user *ticket_entity.TicketEntity) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.PrimaryUserId, validation.Required, validation.By(isUint)),
		validation.Field(&user.SecondaryUserId, validation.Required, validation.By(isUint)),
	)
}

func isUint(value interface{}) error {
	if _, ok := value.(uint); !ok {
		return errors.New("validation_uint must be a uint")
	}
	return nil
}
