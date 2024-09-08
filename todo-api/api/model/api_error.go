package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiError struct {
	Msg  string
	Code int
}

func NewApiError(msg string, code int) *ApiError {
	return &ApiError{Msg: msg, Code: code}
}

func (a *ApiError) Error() string {
	return a.Msg
}

type ApiResponseErrorDto struct {
	ErrorMessage string `json:"error_message"`
}

func ResponseError(ctx echo.Context, err error) error {
	switch e := err.(type) {
	case *ApiError:
		return ctx.JSON(e.Code, ApiResponseErrorDto{ErrorMessage: e.Error()})
	default:
		return ctx.JSON(http.StatusInternalServerError, ApiResponseErrorDto{ErrorMessage: e.Error()})

	}
}
