package customerror

func (errorDetails ErrorDetails) DivisionException() FailureResponse{
	return errorDetails.ServiceError()
}