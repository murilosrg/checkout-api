package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{
			name:     "success",
			filename: "../../products.json",
		},
		{
			name:     "failed",
			filename: "invalid.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitializeDb(tt.filename)

			if tt.name == "success" {
				assert.NotNil(t, db)
				assert.NotNil(t, db.Products)
				assert.Nil(t, err)
			} else {
				assert.Nil(t, db)
				assert.NotNil(t, err)
			}
		})
	}
}
