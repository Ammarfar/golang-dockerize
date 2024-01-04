package repository

import (
	"github.com/Ammarfar/mezink-golang-assignment/internal/domain/models"
)

type UserRepository interface {
	FindAll(*models.UserFilter) ([]*models.UserWithTotalMarks, error)
}
