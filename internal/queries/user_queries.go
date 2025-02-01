package queries

import (
	modeldb "cs-ticketing/internal/models/db"

	"gorm.io/gorm"
)

type UserQueries struct {
	db *gorm.DB
}

func NewUserQueries(db *gorm.DB) *UserQueries {
	return &UserQueries{db: db}
}

func (q *UserQueries) IsUserExistsByID(userID uint) (bool, error) {
	var count int64
	if err := q.db.Model(&modeldb.User{}).Where("id = ?", userID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
