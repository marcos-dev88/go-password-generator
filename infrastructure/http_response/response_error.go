package http_response

import "log"

type CustomError interface {
	DefaultLogResponse()
}

type customError struct {
	StatusCode int
	Message    string
}

func NewCustomError(statusCode int, message string) *customError {
	return &customError{StatusCode: statusCode, Message: message}
}

func (ce *customError) Error() string {
	return ce.Message
}

func GetErrorMessage(err error) string {
	return err.(*customError).Message
}

func GetErrorStatus(err error) int {
	return err.(*customError).StatusCode
}

func (ce *customError) DefaultLogResponse() {
	log.Printf("\nstatus: %v \nmessage: %v", ce.StatusCode, ce.Message)
}
