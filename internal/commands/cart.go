package commands

import (
	"errors"
	"fmt"
)

type Cart struct {
	Products []*ProductRequest `json:"products"`
}

type ProductRequest struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

func (c Cart) Validate() error {
	if c.Products == nil || len(c.Products) == 0 {
		return errors.New("products: must be informed")
	}

	for k, v := range c.Products {
		if v.ID <= 0 {
			return fmt.Errorf("products[%v].id: must be greater than zero", k)
		}

		if v.Quantity <= 0 {
			return fmt.Errorf("products[%v].quantity: must be greater than zero", k)
		}
	}

	return nil
}
