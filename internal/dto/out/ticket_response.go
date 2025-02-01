package dtoout

type Ticket struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Msg    string `json:"msg"`
	UserID uint   `json:"user_id"`
	Status string `json:"status"`
}

// TranslateStatus translates the status code to a user-friendly string
func TranslateStatus(status string) string {
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
