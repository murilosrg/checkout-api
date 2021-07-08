package mocks

import (
	"github.com/murilosrg/checkout-api/internal/config"
	"time"
)

func NewConfigMock() config.Config {
	return config.Config{
		BlackFridayDate: time.Now().UTC().Truncate(24 * time.Hour),
	}
}
