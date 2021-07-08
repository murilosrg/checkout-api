package commands

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCart_validate(t *testing.T) {
	tests := []struct {
		name     string
		cart     Cart
		expected error
	}{
		{
			name:     "validate success",
			cart:     Cart{Products: []*ProductRequest{{ID: 1, Quantity: 1}}},
			expected: nil,
		},
		{
			name:     "validate error: products must be informed",
			cart:     Cart{},
			expected: errors.New("products: must be informed"),
		},
		{
			name:     "validate error: product.id must be greater than zero",
			cart:     Cart{Products: []*ProductRequest{{ID: 0, Quantity: 1}}},
			expected: errors.New("products[0].id: must be greater than zero"),
		},
		{
			name:     "validate error: product.quantity must be greater than zero",
			cart:     Cart{Products: []*ProductRequest{{ID: 1, Quantity: 0}}},
			expected: errors.New("products[0].quantity: must be greater than zero"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.cart.Validate()

			assert.Equal(t, tt.expected, actual)
		})
	}
}
