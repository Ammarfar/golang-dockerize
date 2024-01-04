package user

import (
	"net/http"
	"strconv"

	"github.com/Ammarfar/mezink-golang-assignment/internal/domain"
	"github.com/Ammarfar/mezink-golang-assignment/internal/domain/models"
	user "github.com/Ammarfar/mezink-golang-assignment/internal/modules/user/usecase"
	"github.com/Ammarfar/mezink-golang-assignment/pkg/utils"
	"github.com/labstack/echo/v4"
)

type UserHandlers interface {
	GetAll() echo.HandlerFunc
}

type UserHandlersImpl struct {
	logger domain.Logger
	userUC user.UserUseCase
}

func NewUserHandler(logger domain.Logger, userUC user.UserUseCase) UserHandlers {
	return &UserHandlersImpl{
		logger: logger,
		userUC: userUC,
	}
}

func (h *UserHandlersImpl) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		startDate := c.QueryParam("startDate")
		endDate := c.QueryParam("endDate")
		minCount, _ := strconv.Atoi(c.QueryParam("minCount"))
		maxCount, _ := strconv.Atoi(c.QueryParam("maxCount"))
		userFilter := &models.UserFilter{
			StartDate: &startDate,
			EndDate:   &endDate,
			MinCount:  &minCount,
			MaxCount:  &maxCount,
		}

		users, err := h.userUC.GetAll(userFilter)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.Response{
				Code: 1,
				Msg:  err.Error(),
			})
		}

		return c.JSON(http.StatusOK, utils.Response{
			Code:    0,
			Msg:     "success",
			Records: users,
		})
	}
}
