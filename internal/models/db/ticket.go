package modeldb

type Ticket struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Msg       string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	StatusID  uint   `gorm:"not null"`
	CreatedAt int64  `gorm:"not null"`
	UpdatedAt int64  `gorm:"null"`
	DeletedAt int64  `gorm:"null"`
}
