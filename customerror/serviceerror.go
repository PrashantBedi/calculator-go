package customerror

import (
	"calculator/errorcodes"
)

func (errorDetails ErrorDetails) ServiceError() FailureResponse{
	var failureResponse FailureResponse
	failureResponse.Message = errorDetails.ErrorCode
	failureResponse.Reasons = map[errorcodes.ErrorCodes]string{ errorDetails.ErrorCode: errorDetails.Message}
	return failureResponse
}