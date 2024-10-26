package ticket_interfaces_repo

import ticket_entity "github.com/luizrgf2/chat_real_time_go/internal/ticket/entities"

type IUserRepository interface {
	Create(ticket *ticket_entity.TicketEntity) (*ticket_entity.TicketEntity, error)
	FindOne(id uint) (*ticket_entity.TicketEntity, error)
	FindByUserId(primaryUserId *uint, secondaryUserId *uint)
	FindAnyUserId(userId uint)
}
