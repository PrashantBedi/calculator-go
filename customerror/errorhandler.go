package customerror

import (
	"calculator/errorcodes"
)

type ErrorDetails struct {
	Message string
	ErrorCode errorcodes.ErrorCodes
}