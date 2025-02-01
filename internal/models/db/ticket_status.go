package modeldb

type TicketStatus struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Note      string `gorm:"null"`
	CreatedAt string `gorm:"not null"`
	UpdatedAt string `gorm:"null"`
	DeleteAt  string `gorm:"null"`
}
