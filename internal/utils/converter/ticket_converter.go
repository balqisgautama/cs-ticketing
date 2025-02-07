package utilsconverter

import (
	dtoout "cs-ticketing/internal/dto/out"
	modeldb "cs-ticketing/internal/models/db"
	"cs-ticketing/internal/queries"
	"time"
)

type UtilsConverterTiket struct {
	ticketStatusQueries *queries.TicketStatusQueries
}

func NewUtilsConverterTicket(ticketStatusQueries *queries.TicketStatusQueries) *UtilsConverterTiket {
	return &UtilsConverterTiket{ticketStatusQueries: ticketStatusQueries}
}

func (u *UtilsConverterTiket) TranslateStatus(status string) string {
	switch status {
	case "opn":
		return "Open"
	case "cld":
		return "Closed"
	case "asn":
		return "Assigned"
	default:
		return "Unknown"
	}
}

// func for unix timestamp to time.Time
// zone is Jakarta
func (u *UtilsConverterTiket) UnixToTime(unix int64) string {
	t := time.Unix(unix, 0)
	loc := time.FixedZone("WIB", 7*3600) // WIB (Western Indonesian Time) UTC+7
	return t.In(loc).Format("2006-01-02 15:04:05")
}

func (u *UtilsConverterTiket) TicketListToTicketListResponse(
	tickets []modeldb.Ticket,
	total int,
	pageSize int,
	page int,
) dtoout.TicketList {
	var result []dtoout.Ticket
	for _, ticket := range tickets {
		ticketResponse, err := u.TicketToTicketResponse(ticket)
		if err != nil {
			continue
		}
		result = append(result, *ticketResponse)
	}

	totalPages := total / pageSize
	if total%pageSize != 0 {
		totalPages++
	}

	currentPage := page
	previousPage := currentPage - 1
	if previousPage < 1 {
		previousPage = 0 // or null
	}

	nextPage := currentPage + 1
	if nextPage > totalPages {
		nextPage = 0 // or null
	}

	return dtoout.TicketList{
		Tickets:      result,
		TotalPage:    totalPages,
		CurrentPage:  currentPage,
		PreviousPage: previousPage,
		NextPage:     nextPage,
		TotalItems:   total,
		PageSize:     pageSize,
	}
}

func (u *UtilsConverterTiket) TicketToTicketResponse(ticket modeldb.Ticket) (*dtoout.Ticket, error) {
	status, err := u.ticketStatusQueries.FindTicketStatusByID(ticket.StatusID)
	if err != nil {
		return nil, err
	}

	result := &dtoout.Ticket{
		ID:        ticket.ID,
		Title:     ticket.Title,
		Msg:       ticket.Msg,
		UserID:    ticket.UserID,
		Status:    u.TranslateStatus(status.Name),
		CreatedAt: u.UnixToTime(ticket.CreatedAt),
	}

	return result, nil
}
