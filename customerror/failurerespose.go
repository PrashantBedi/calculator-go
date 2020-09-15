package customerror

import (
	"calculator/errorcodes"
)

type FailureResponse struct {
	Message errorcodes.ErrorCodes
	Reasons map[errorcodes.ErrorCodes]string
}