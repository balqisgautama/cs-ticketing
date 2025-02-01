package modeldb

type User struct {
	ID        uint   `gorm:"primaryKey"`
	ClientID  uint   `gorm:"not null,unique"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"not null,unique"`
	Phone     string `gorm:"null,unique"`
	IsActive  bool   `gorm:"not null"`
	CreatedAt string `gorm:"not null"`
	UpdatedAt string `gorm:"null"`
	DeleteAt  string `gorm:"null"`
}
