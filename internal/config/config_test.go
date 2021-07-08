package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_newConfig(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{
			name:     "success",
			filename: "../../.env",
		},
		{
			name:     "failed",
			filename: "invalid",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf, err := NewConfig(tt.filename)

			if tt.name == "success" {
				assert.NotNil(t, conf)
				assert.NotEmpty(t, conf.Discount)
				assert.NotEmpty(t, conf.DatabaseFile)
				assert.NotEmpty(t, conf.BlackFridayDate)
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

