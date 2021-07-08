package mocks

import (
	"context"
	"errors"
)

type DiscountServiceMock interface {
	Get(ctx context.Context, id int) (float32, error)
}

type discountMock struct{}

func NewDiscountServiceMock() DiscountServiceMock {
	return discountMock{}
}

func (s discountMock) Get(ctx context.Context, id int) (float32, error) {
	if id != 1 {
		return 0, errors.New("error")
	}

	return 0.05, nil
}
