package main

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
