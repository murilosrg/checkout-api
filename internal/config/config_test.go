package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfig_NewConfigSuccess(t *testing.T) {
	os.Setenv("BLACK_FRIDAY_DATE", "2021-07-07")
	os.Setenv("DATABASE_FILE", "test")
	os.Setenv("DISCOUNT_URI", "test")

	conf, err := NewConfig()

	assert.NotNil(t, conf)
	assert.NotEmpty(t, conf.Discount)
	assert.NotEmpty(t, conf.DatabaseFile)
	assert.NotEmpty(t, conf.BlackFridayDate)
	assert.Nil(t, err)

	os.Unsetenv("BLACK_FRIDAY_DATE")
	os.Unsetenv("DATABASE_FILE")
	os.Unsetenv("DISCOUNT_URI")
}

func TestConfig_NewConfigFailed(t *testing.T) {
	_, err := NewConfig()

	assert.NotNil(t, err)
}

