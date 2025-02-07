package dtoout

type Ticket struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Msg       string `json:"msg"`
	UserID    uint   `json:"user_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type TicketList struct {
	Tickets      []Ticket `json:"tickets"`
	TotalPage    int      `json:"total_page"`
	CurrentPage  int      `json:"current_page"`
	PreviousPage int      `json:"previous_page"`
	NextPage     int      `json:"next_page"`
	TotalItems   int      `json:"total_items"`
	PageSize     int      `json:"page_size"`
}
