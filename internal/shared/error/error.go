package error

import "fmt"

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
