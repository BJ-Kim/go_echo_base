package modules

import (
	"github.com/labstack/echo/v4"
)

func ErrorHandler(ctx echo.Context, err error) error {
	if ce, ok := err.(*CustomError); ok {
		return ctx.JSON(ce.Status, ce)
	}
	return ctx.JSON(500, CustomError{
		Status:    500,
		ErrorCode: 500,
		Message:   "Internal Server Error",
	})
}

type CustomError struct {
	Status    int    `json:"status"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func NewInvalidValueError() *CustomError {
	return &CustomError{
		Status:    400,
		ErrorCode: 100,
		Message:   "invalid value error",
	}
}
