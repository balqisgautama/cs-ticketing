package service

import (
	dtoin "cs-ticketing/internal/dto/in"
	dtoout "cs-ticketing/internal/dto/out"
	modeldb "cs-ticketing/internal/models/db"
	"cs-ticketing/internal/queries"
	utilsconverter "cs-ticketing/internal/utils/converter"
	"errors"
	"time"
)

type TicketService struct {
	ticketQueries       *queries.TicketQueries
	userQueries         *queries.UserQueries
	ticketStatusQueries *queries.TicketStatusQueries
	utilsConverterTiket *utilsconverter.UtilsConverterTiket
}

func NewTicketService(
	ticketQueries *queries.TicketQueries,
	userQueries *queries.UserQueries,
	ticketStatusQueries *queries.TicketStatusQueries,
	utilsConverterTiket *utilsconverter.UtilsConverterTiket,
) *TicketService {
	return &TicketService{
		ticketQueries:       ticketQueries,
		userQueries:         userQueries,
		ticketStatusQueries: ticketStatusQueries,
		utilsConverterTiket: utilsConverterTiket,
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

	result, err := s.utilsConverterTiket.TicketToTicketResponse(*ticketFromDB)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *TicketService) GetTickets(req dtoin.TicketList, startDate int64, endDate int64) (*dtoout.TicketList, error) {
	ticketFilterModel := &modeldb.TicketFilter{}
	ticketSortModel := &modeldb.TicketSort{}

	if req.Filter != nil {
		ticketFilterModel.FilterValue = startDate
		ticketFilterModel.FilterValue2 = endDate
		ticketFilterModel.FilterName = req.Filter.FilterName
		ticketFilterModel.FilterType = req.Filter.FilterType
	}

	if req.Sort != nil {
		ticketSortModel.SortName = req.Sort.SortName
		ticketSortModel.SortDir = req.Sort.SortDir
	}

	tickets, total, err := s.ticketQueries.GetTickets(ticketFilterModel, ticketSortModel, req.PageSize, req.Page)
	if err != nil {
		return nil, err
	}

	result := &dtoout.TicketList{}
	if len(tickets) > 0 {
		result = s.utilsConverterTiket.TicketListToTicketListResponse(tickets, int(total), req.PageSize, req.Page)
	}

	return result, nil
}
