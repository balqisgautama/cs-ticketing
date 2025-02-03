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

func (r *TicketQueries) GetTickets(
	filter *modeldb.TicketFilter,
	sort *modeldb.TicketSort,
	pageSize, page int,
) ([]modeldb.Ticket, int64, error) {
	var tickets []modeldb.Ticket
	var total int64

	query := r.db.Model(&modeldb.Ticket{})

	if filter != nil {
		switch filter.FilterType {
		case "before":
			query = query.Where("created_at < ?", filter.FilterValue)
		case "after":
			query = query.Where("created_at > ?", filter.FilterValue)
		case "between":
			query = query.Where("created_at BETWEEN ? AND ?", filter.FilterValue, filter.FilterValue2)
		}
	}

	switch sort.SortName {
	case "created_at":
		if sort.SortDir == "asc" {
			query = query.Order("created_at ASC")
		} else {
			query = query.Order("created_at DESC")
		}
	case "user_id":
		if sort.SortDir == "asc" {
			query = query.Order("user_id ASC")
		} else {
			query = query.Order("user_id DESC")
		}
	default:
		query = query.Order("created_at DESC")
	}

	// Hitung total record sebelum menerapkan offset dan limit
	query.Count(&total)

	// Terapkan offset dan limit
	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	if err := query.Find(&tickets).Error; err != nil {
		return nil, 0, err
	}

	return tickets, total, nil
}
