package checkout

import (
	"github.com/murilosrg/checkout-api/internal/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	checkoutService Service
)

func TestService_ValidRequest(t *testing.T) {
	loadService()
	actual, _ := checkoutService.Checkout(mocks.ValidCart())
	assert.NotNil(t, actual)
}

func TestService_InvalidRequest(t *testing.T) {
	loadService()
	_, err := checkoutService.Checkout(mocks.InvalidCart())
	assert.NotNil(t, err)
}

func loadService() {
	if checkoutService == nil {
		product := mocks.NewProductServiceMock()
		discount := mocks.NewDiscountServiceMock()
		conf := mocks.NewConfigMock()
		checkoutService = NewService(conf, product, discount, logrus.New())
	}
}
