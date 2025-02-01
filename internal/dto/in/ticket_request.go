package dtoin

type Ticket struct {
	Title  string `json:"ticket_title" validate:"required,min=10,max=100"`
	Msg    string `json:"ticket_msg" validate:"required,min=100"`
	UserID uint   `json:"user_id" validate:"required,number"`
}
