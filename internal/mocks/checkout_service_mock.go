package mocks

import (
	"errors"
	"github.com/murilosrg/checkout-api/internal/commands"
)

type CheckoutServiceMock interface {
	Checkout(cart commands.Cart) (commands.Checkout, error)
}

type checkoutMock struct{}

func NewCheckoutServiceMock() CheckoutServiceMock {
	return checkoutMock{}
}

func (c checkoutMock) Checkout(cart commands.Cart) (commands.Checkout, error) {
	if cart.Products[0].ID != 1 {
		return commands.Checkout{}, errors.New("error")
	}

	return commands.Checkout{
		TotalAmount:             1000,
		TotalAmountWithDiscount: 950,
		TotalDiscount:           50,
		Products: []commands.ProductResponse{
			{
				ID:          1,
				Quantity:    1,
				UnitAmount:  1000,
				TotalAmount: 1000,
				Discount:    50,
				IsGift:      false,
			},
		},
	}, nil
}
