package config

import (
	"errors"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DatabaseFile    string
	BlackFridayDate time.Time
	Discount        string
}

func NewConfig() (Config, error) {
	stringDate := os.Getenv("BLACK_FRIDAY_DATE")
	dbFilename := os.Getenv("DATABASE_FILE")
	discountUri := os.Getenv("DISCOUNT_URI")

	if len(stringDate) == 0 || len(dbFilename) == 0 || len(discountUri) == 0 {
		return Config{}, errors.New("environments variables not set")
	}

	date, _ := time.Parse("2006-01-02", stringDate)

	return Config{
		DatabaseFile:    os.Getenv("DATABASE_FILE"),
		BlackFridayDate: date,
		Discount:        os.Getenv("DISCOUNT_URI"),
	}, nil
}
