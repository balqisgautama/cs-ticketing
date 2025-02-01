package queries

import (
	modeldb "cs-ticketing/internal/models/db"

	"gorm.io/gorm"
)

type TicketQueries struct {
	db *gorm.DB
}

func NewTicketQueries(db *gorm.DB) *TicketQueries {
	return &TicketQueries{db: db}
}

func (q *TicketQueries) CreateTicket(ticket *modeldb.Ticket) (*modeldb.Ticket, error) {
	result := q.db.Create(ticket)
	if result.Error != nil {
		return nil, result.Error
	}
	return ticket, nil
}
