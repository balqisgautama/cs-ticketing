package dtoout

type Ticket struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Msg    string `json:"msg"`
	UserID uint   `json:"user_id"`
	Status string `json:"status"`
}

// TranslateStatus translates the status code to a user-friendly string
func TranslateStatus(statusID uint) string {
	switch statusID {
	case 1:
		return "Open"
	case 2:
		return "Closed"
	case 3:
		return "Assigned"
	default:
		return "Unknown"
	}
}
