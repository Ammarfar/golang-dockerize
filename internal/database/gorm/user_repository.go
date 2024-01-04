package gorm_repository

import (
	"fmt"

	"github.com/Ammarfar/mezink-golang-assignment/internal/domain/models"
	"github.com/Ammarfar/mezink-golang-assignment/internal/domain/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepositoryGorm(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryGorm{
		db: db,
	}
}

func (ur *UserRepositoryGorm) FindAll(filter *models.UserFilter) ([]*models.UserWithTotalMarks, error) {
	var users []*models.UserWithTotalMarks

	clauses := make([]clause.Expression, 0)
	if filter.StartDate != nil && *filter.StartDate != "" {
		clauses = append(clauses, clause.Gte{Column: "created_at", Value: fmt.Sprintf("%s 00:00:00", *filter.StartDate)})
	}
	if filter.EndDate != nil && *filter.EndDate != "" {
		clauses = append(clauses, clause.Lte{Column: "created_at", Value: fmt.Sprintf("%s 23:59:59", *filter.EndDate)})
	}

	if *filter.MaxCount == 0 {
		maxCount := 9999999999
		filter.MaxCount = &maxCount
	}

	if err := ur.db.
		Select("records.id, records.created_at, sum(mark) as total_marks").
		Joins("JOIN marks ON marks.user_id = records.id").
		Clauses(clauses...).
		Group("records.id, records.created_at").
		Having("total_marks >= ? AND total_marks <= ?", filter.MinCount, filter.MaxCount).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
