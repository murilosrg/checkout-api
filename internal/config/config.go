package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseFile    string
	BlackFridayDate time.Time
	Discount        string
}

func NewConfig(filename string) (Config, error) {
	if err := godotenv.Load(filename); err != nil {
		return Config{}, err
	}

	stringDate := os.Getenv("BLACK_FRIDAY_DATE")
	date, _ := time.Parse("2006-01-02", stringDate)

	return Config{
		DatabaseFile:    os.Getenv("DATABASE_FILE"),
		BlackFridayDate: date,
		Discount:        os.Getenv("DISCOUNT_URI"),
	}, nil
}
