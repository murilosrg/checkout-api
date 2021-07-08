package db

import (
	"encoding/json"
	"github.com/murilosrg/checkout-api/internal/entities"
	"io/ioutil"
)

type DB struct {
	Products []entities.Product
}

func InitializeDb(filename string) (*DB, error) {
	db := &DB{}

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &db.Products)

	return db, err
}
