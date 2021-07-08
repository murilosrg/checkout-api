package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorResponse_Error(t *testing.T) {
	e := ErrorResponse{
		Message: "test",
	}
	assert.Equal(t, "test", e.Error())
}

func TestErrorResponse_StatusCode(t *testing.T) {
	e := ErrorResponse{
		Status: 201,
	}
	assert.Equal(t, 201, e.StatusCode())
}

func TestInternalServerError(t *testing.T) {
	res := InternalServerError("internal error")

	assert.Equal(t, InternalServerErrorStatusCode, res.StatusCode())
	assert.Equal(t, "internal error", res.Error())
}

func TestUnprocessableEntity(t *testing.T) {
	res := UnprocessableEntity("unprocessable entity")

	assert.Equal(t, UnprocessableEntityStatusCode, res.StatusCode())
	assert.Equal(t, "unprocessable entity", res.Error())
}

func TestBadRequest(t *testing.T) {
	res := BadRequest("bad request")

	assert.Equal(t, BadRequestStatusCode, res.StatusCode())
	assert.Equal(t, "bad request", res.Error())
}