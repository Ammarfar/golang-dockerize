package user

import (
	"errors"

	"github.com/Ammarfar/mezink-golang-assignment/config"
	"github.com/Ammarfar/mezink-golang-assignment/internal/domain"
	"github.com/Ammarfar/mezink-golang-assignment/internal/domain/models"
	"github.com/Ammarfar/mezink-golang-assignment/internal/domain/repository"
)

type UserUseCase interface {
	GetAll(userFilter *models.UserFilter) ([]*models.UserWithTotalMarks, error)
}

type UserUseCaseImpl struct {
	logger   domain.Logger
	userRepo repository.UserRepository
}

func NewUserUseCase(
	logger domain.Logger,
	userRepo repository.UserRepository,
) UserUseCase {
	return &UserUseCaseImpl{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (u *UserUseCaseImpl) GetAll(userFilter *models.UserFilter) ([]*models.UserWithTotalMarks, error) {
	user, err := u.userRepo.FindAll(userFilter)
	if err != nil {
		u.logger.Error(&domain.LoggerPayload{
			Loc: "user.userRepo.FindAll",
			Msg: err.Error(),
			Req: userFilter,
		})
		return nil, errors.New(config.ERROR_DATABASE)
	}

	return user, nil
}
