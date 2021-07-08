package products

import (
	"github.com/murilosrg/checkout-api/internal/entities"
	"github.com/murilosrg/checkout-api/pkg/log"
)

type Service interface {
	Get(id int, isGift bool) (entities.Product, error)
}

type service struct {
	repo   Repository
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

func (s service) Get(id int, isGift bool) (entities.Product, error) {
	if isGift {
		return s.repo.GetGift()
	}

	return s.repo.Find(id)
}