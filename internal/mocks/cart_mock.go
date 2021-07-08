package mocks

import "github.com/murilosrg/checkout-api/internal/commands"

func ValidCart() commands.Cart {
	return commands.Cart{
		Products: []*commands.ProductRequest{
			{
				ID:       1,
				Quantity: 1,
			},
		},
	}
}

func InvalidCart() commands.Cart {
	return commands.Cart{
		Products: []*commands.ProductRequest{
			{
				ID:       999,
				Quantity: 1,
			},
		},
	}
}
