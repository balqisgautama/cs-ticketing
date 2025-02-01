package service

import (
	dtoin "cs-ticketing/internal/dto/in"
	dtoout "cs-ticketing/internal/dto/out"
	modeldb "cs-ticketing/internal/models/db"
	"cs-ticketing/internal/queries"
	"errors"
	"time"
)

type TicketService struct {
	ticketQueries       *queries.TicketQueries
	userQueries         *queries.UserQueries
	ticketStatusQueries *queries.TicketStatusQueries
}

func NewTicketService(
	ticketQueries *queries.TicketQueries,
	userQueries *queries.UserQueries,
	ticketStatusQueries *queries.TicketStatusQueries,
) *TicketService {
	return &TicketService{
		ticketQueries:       ticketQueries,
		userQueries:         userQueries,
		ticketStatusQueries: ticketStatusQueries,
	}
}

func (s *TicketService) CreateTicket(dtoin *dtoin.Ticket) (*dtoout.Ticket, error) {
	// Check if the user exists
	exists, err := s.userQueries.IsUserExistsByID(dtoin.UserID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("user does not exist")
	}

	// Create a ticket
	ticketModel := modeldb.Ticket{
		UserID:    dtoin.UserID,
		Title:     dtoin.Title,
		Msg:       dtoin.Msg,
		StatusID:  1, // 1 is the default status
		CreatedAt: time.Now().Unix(),
	}

	ticketFromDB, err := s.ticketQueries.CreateTicket(&ticketModel)
	if err != nil {
		return nil, err
	}

	status, err := s.ticketStatusQueries.FindTicketStatusByID(ticketFromDB.StatusID)
	if err != nil {
		return nil, err
	}

	result := dtoout.Ticket{
		ID:     ticketFromDB.ID,
		UserID: ticketFromDB.UserID,
		Title:  ticketFromDB.Title,
		Msg:    ticketFromDB.Msg,
		Status: dtoout.TranslateStatus(status.Name),
	}

	return &result, nil
}
