package ticket_service_test

import (
	"testing"

	ticket_service "github.com/luizrgf2/chat_real_time_go/internal/ticket/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateValidTicket(t *testing.T) {
	ticketService := ticket_service.CreateNewTicketService()
	ticket, err := ticketService.CreateNewTicket(1, 2)

	assert.Nil(t, err)
	assert.Equal(t, ticket.PrimaryUserId, uint(1))
	assert.Equal(t, ticket.SecondaryUserId, uint(2))
}
