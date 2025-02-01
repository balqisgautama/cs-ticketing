package queries

import (
	modeldb "cs-ticketing/internal/models/db"

	"gorm.io/gorm"
)

type TicketStatusQueries struct {
	db *gorm.DB
}

func NewTicketStatusQueries(db *gorm.DB) *TicketStatusQueries {
	return &TicketStatusQueries{db: db}
}

func (q *TicketStatusQueries) FindTicketStatusByID(statusID uint) (*modeldb.TicketStatus, error) {
	var status modeldb.TicketStatus
	if err := q.db.First(&status, statusID).Error; err != nil {
		return nil, err
	}
	return &status, nil
}
