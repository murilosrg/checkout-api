package mocks

import (
	"errors"
	"github.com/murilosrg/checkout-api/internal/entities"
)

type MockRepository interface {
	Find(id int) (entities.Product, error)
	GetGift() (entities.Product, error)
}

type productMockRepo struct{}

func NewMockRepository() MockRepository {
	return productMockRepo{}
}

func (m productMockRepo) Find(id int) (entities.Product, error) {
	if id != 1 {
		return entities.Product{}, errors.New("error")
	}

	return entities.Product{
		ID:          1,
		Title:       "test",
		Description: "test",
		Amount:      10,
		IsGift:      false,
	}, nil
}

func (m productMockRepo) GetGift() (entities.Product, error) {
	return entities.Product{
		ID:          1,
		Title:       "test",
		Description: "test",
		Amount:      10,
		IsGift:      true,
	}, nil
}
