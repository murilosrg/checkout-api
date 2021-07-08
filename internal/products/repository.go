package products

import (
	"errors"
	"fmt"
	"github.com/murilosrg/checkout-api/internal/entities"
	"github.com/murilosrg/checkout-api/pkg/db"
	"github.com/murilosrg/checkout-api/pkg/log"
	"math/rand"
	"time"
)

type Repository interface {
	Find(id int) (entities.Product, error)
	GetGift() (entities.Product, error)
}

type repository struct {
	db     *db.DB
	logger log.Logger
}

func NewRepository(db *db.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) Find(id int) (entities.Product, error) {
	for _, val := range r.db.Products {
		if val.ID == id {
			return val, nil
		}
	}

	return entities.Product{}, fmt.Errorf("product %v not found", id)
}

func (r repository) GetGift() (entities.Product, error) {
	gifts := make([]entities.Product, 0)

	for _, val := range r.db.Products {
		if val.IsGift {
			gifts = append(gifts, val)
		}
	}

	if len(gifts) == 0 {
		return entities.Product{}, errors.New("gift not found")
	}

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(gifts)
	return gifts[n], nil
}