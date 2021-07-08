package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckout_NewGift(t *testing.T) {
	p := NewProductGift(1)

	assert.Equal(t, 1, p.ID)
	assert.Equal(t, 1, p.Quantity)
	assert.Equal(t, 0, p.UnitAmount)
	assert.Equal(t, 0, p.TotalAmount)
	assert.Equal(t, 0, p.Discount)
	assert.Equal(t, true, p.IsGift)
}
