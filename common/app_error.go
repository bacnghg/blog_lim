package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"_"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func ErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func FullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func Unauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func CustomError(root error, msg, key string) *AppError {
	if root != nil {
		return ErrorResponse(root, msg, root.Error(), key)
	}

	return ErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()

	}
	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrorDB(err error) *AppError {
	return ErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}
func InvalidRequest(err error) *AppError {
	return ErrorResponse(err, "Invalid request", err.Error(), "InvalidRequest")
}
func ErrorInternal(err error) *AppError {
	return FullErrorResponse(http.StatusInternalServerError, err, "Something went wrong in the server", err.Error(), "InternalError")
}
