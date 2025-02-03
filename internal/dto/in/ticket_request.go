package dtoin

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
