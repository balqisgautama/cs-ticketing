package dtoin

import "errors"

type Ticket struct {
	Title  string `json:"ticket_title" validate:"required,min=10,max=100"`
	Msg    string `json:"ticket_msg" validate:"required,min=100"`
	UserID uint   `json:"user_id" validate:"required,number"`
}

type TicketList struct {
	Filter   *TicketFilter `json:"filter" validate:"omitempty"`
	Sort     *TicketSort   `json:"sort" validate:"omitempty"`
	PageSize int           `json:"page_size" validate:"required,number,min=1"`
	Page     int           `json:"page" validate:"required,number,min=1"`
}

type TicketFilter struct {
	FilterName  string `json:"filter_name" validate:"omitempty,oneof=created_at"`
	FilterType  string `json:"filter_type" validate:"omitempty,oneof=before after between"`
	FilterValue string `json:"filter_value" validate:"omitempty"`
}

type TicketSort struct {
	SortName string `json:"sort_name" validate:"omitempty"`
	SortDir  string `json:"sort_dir" validate:"omitempty,oneof=asc desc"`
}

func ValidateTicketFilter(filter *TicketFilter) error {
	if filter == nil {
		return nil
	}

	if filter.FilterName == "" || filter.FilterType == "" || filter.FilterValue == "" {
		return errors.New("filter_name, filter_type, and filter_value must be filled if filter is not nil")
	}

	return nil
}

func ValidateTicketSort(sort *TicketSort) error {
	if sort == nil {
		return nil
	}

	if sort.SortName == "" || sort.SortDir == "" {
		return errors.New("sort_name and sort_dir must be filled if sort is not nil")
	}

	return nil
}
