package modeldb

type Ticket struct {
	ID        int64  `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Msg       string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	StatusID  int64  `gorm:"not null"`
	CreatedAt int64  `gorm:"not null"`
	UpdatedAt int64  `gorm:"null"`
	DeletedAt int64  `gorm:"null"`
}

type TicketFilter struct {
	FilterName   string
	FilterType   string
	FilterValue  int64
	FilterValue2 int64 // For between filter
}

type TicketSort struct {
	SortName string
	SortDir  string
}
