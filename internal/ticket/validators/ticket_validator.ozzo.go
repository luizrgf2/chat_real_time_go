package ticket_validator

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	ticket_entity "github.com/luizrgf2/chat_real_time_go/internal/ticket/entities"
)

type OzzoTicketValidator struct{}

func (v *OzzoTicketValidator) Validate(user *ticket_entity.TicketEntity) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.PrimaryUserId, validation.Required, is.Int),
		validation.Field(&user.SecondaryUserId, validation.Required, is.Int),
	)
}
