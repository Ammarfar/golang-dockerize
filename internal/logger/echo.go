package logger

import (
	"github.com/Ammarfar/mezink-golang-assignment/internal/domain"
	"github.com/Ammarfar/mezink-golang-assignment/pkg/utils"
	"github.com/labstack/echo/v4"
)

type EchoLogger struct {
	logger echo.Logger
}

func NewEchoLogger(logger echo.Logger) domain.Logger {
	return &EchoLogger{
		logger: logger,
	}
}

func (el *EchoLogger) Error(payload *domain.LoggerPayload) {
	el.logger.Error(utils.WriteToErrorLog(payload))
}

func (el *EchoLogger) Info(payload *domain.LoggerPayload) {
	el.logger.Info(utils.WriteToErrorLog(payload))
}
