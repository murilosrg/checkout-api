package errors

const (
	InternalServerErrorStatusCode = 500
	UnprocessableEntityStatusCode = 422
	BadRequestStatusCode          = 400
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return e.Message
}

func (e ErrorResponse) StatusCode() int {
	return e.Status
}

func InternalServerError(msg string) ErrorResponse {
	return ErrorResponse{
		Status:  InternalServerErrorStatusCode,
		Message: msg,
	}
}

func UnprocessableEntity(msg string) ErrorResponse {
	return ErrorResponse{
		Status:  UnprocessableEntityStatusCode,
		Message: msg,
	}
}

func BadRequest(msg string) ErrorResponse {
	return ErrorResponse{
		Status:  BadRequestStatusCode,
		Message: msg,
	}
}
