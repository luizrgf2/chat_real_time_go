package ticket_entity

import "time"

type TicketEntity struct {
	Id              uint
	PrimaryUserId   uint
	SecondaryUserId uint
	CreatedAt       time.Time
	UpdateAt        time.Time
}
