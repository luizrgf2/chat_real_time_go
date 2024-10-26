package ticket_service

import (
	"time"

	"github.com/luizrgf2/chat_real_time_go/internal/helpers"
	ticket_entity "github.com/luizrgf2/chat_real_time_go/internal/ticket/entities"
	ticket_validator "github.com/luizrgf2/chat_real_time_go/internal/ticket/validators"
)

type TicketService struct {
	validator helpers.Validator[ticket_entity.TicketEntity]
}

func CreateNewTicketService() TicketService {
	validator := ticket_validator.OzzoTicketValidator{}
	ticketService := TicketService{
		validator: &validator,
	}
	return ticketService
}

func (t *TicketService) CreateNewTicket(userPrimaryId uint, userSecondaryId uint) (*ticket_entity.TicketEntity, error) {
	ticket := ticket_entity.TicketEntity{
		Id:              1,
		PrimaryUserId:   userPrimaryId,
		SecondaryUserId: userSecondaryId,
		CreatedAt:       time.Now(),
		UpdateAt:        time.Now(),
	}

	err := t.validator.Validate(&ticket)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}
