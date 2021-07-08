package mocks

import (
	"errors"
	"github.com/murilosrg/checkout-api/internal/entities"
)

type ProductServiceMock interface {
	Get(id int, isGift bool) (entities.Product, error)
}

type productMock struct{}

func NewProductServiceMock() ProductServiceMock {
	return productMock{}
}

func (s productMock) Get(id int, isGift bool) (entities.Product, error) {
	if isGift {
		return entities.Product{
			ID:     2,
			IsGift: true,
		}, nil
	}

	if id != 1 {
		return entities.Product{}, errors.New("error")
	}

	return entities.Product{
		ID:          1,
		Title:       "test",
		Description: "test",
		Amount:      1000,
		IsGift:      false,
	}, nil
}
