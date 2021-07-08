package products

import (
	"github.com/murilosrg/checkout-api/internal/mocks"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	productService Service
)

func TestService_GetProduct(t *testing.T) {
	loadService()

	actual, err := productService.Get(1, false)
	assert.NotNil(t, actual)
	assert.Equal(t, 1, actual.ID)
	assert.Nil(t, err)
}

func TestService_GetProductNotFound(t *testing.T) {
	loadService()

	actual, err := productService.Get(9999, false)
	assert.Equal(t, 0, actual.ID)
	assert.NotNil(t, err)
}

func TestService_GetGift(t *testing.T) {
	loadService()

	actual, err := productService.Get(0, true)
	assert.NotNil(t, actual)
	assert.Equal(t, true, actual.IsGift)
	assert.Nil(t, err)
}

func loadService() {
	if productService == nil {
		repo := mocks.NewMockRepository()
		productService = NewService(repo, logrus.New())
	}
}
