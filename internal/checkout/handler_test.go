package checkout

import (
	"bytes"
	"encoding/json"
	"github.com/murilosrg/checkout-api/internal/commands"
	"github.com/murilosrg/checkout-api/internal/errors"
	"github.com/murilosrg/checkout-api/internal/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Success(t *testing.T) {
	checkout := mocks.NewCheckoutServiceMock()
	handler := NewHandler(checkout, logrus.New())

	b, _ := json.Marshal(mocks.ValidCart())
	request := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	responseWriter := httptest.NewRecorder()

	handler.Post(responseWriter, request)
	response, actual := extractResponseFromResponseWriter(responseWriter)
	responseString, _ := json.Marshal(expectedResponse())

	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, string(responseString), actual)
}

func TestHandler_UnprocessableEntity(t *testing.T) {
	checkout := mocks.NewCheckoutServiceMock()
	handler := NewHandler(checkout, logrus.New())

	b, _ := json.Marshal(mocks.InvalidCart())
	request := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	responseWriter := httptest.NewRecorder()

	handler.Post(responseWriter, request)
	response, actual := extractResponseFromResponseWriter(responseWriter)
	errorResponse := errors.UnprocessableEntity("error")
	responseString, _ := json.Marshal(errorResponse)

	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(t, http.StatusUnprocessableEntity, response.StatusCode)
	assert.Equal(t, string(responseString), actual)
}

func TestHandler_BadRequest(t *testing.T) {
	checkout := mocks.NewCheckoutServiceMock()
	handler := NewHandler(checkout, logrus.New())

	b, _ := json.Marshal(commands.Cart{})
	request := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	responseWriter := httptest.NewRecorder()

	handler.Post(responseWriter, request)
	response, actual := extractResponseFromResponseWriter(responseWriter)
	errorResponse := errors.BadRequest("products: must be informed")
	responseString, _ := json.Marshal(errorResponse)

	assert.Equal(t, "application/json", response.Header.Get("Content-Type"))
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, string(responseString), actual)
}

func extractResponseFromResponseWriter(responseWriter *httptest.ResponseRecorder) (*http.Response, string) {
	resp := responseWriter.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := string(body)
	actual = strings.TrimSuffix(actual, "\n")
	return resp, actual
}

func expectedResponse() commands.Checkout {
	return commands.Checkout{
		TotalAmount:             1000,
		TotalAmountWithDiscount: 950,
		TotalDiscount:           50,
		Products: []commands.ProductResponse{{ID: 1, Quantity: 1, UnitAmount: 1000, TotalAmount: 1000, Discount: 50, IsGift: false}},
	}
}

