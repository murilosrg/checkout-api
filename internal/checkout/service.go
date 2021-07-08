package checkout

import (
	"context"
	"github.com/murilosrg/checkout-api/internal/commands"
	"github.com/murilosrg/checkout-api/internal/config"
	"github.com/murilosrg/checkout-api/internal/discount"
	"github.com/murilosrg/checkout-api/internal/entities"
	"github.com/murilosrg/checkout-api/internal/products"
	"github.com/murilosrg/checkout-api/pkg/log"
	"math"
	"time"
)

type Service interface {
	Checkout(cart commands.Cart) (commands.Checkout, error)
}

type service struct {
	conf     config.Config
	product  products.Service
	discount discount.Service
	logger   log.Logger
}

func NewService(conf config.Config, product products.Service, discount discount.Service, logger log.Logger) Service {
	return service{conf, product, discount, logger}
}

func (s service) Checkout(cart commands.Cart) (commands.Checkout, error) {
	checkout := commands.Checkout{}
	productsMap := handleProductMap(cart.Products)

	for k, v := range productsMap {
		prod, disc, err := s.handleProductAndDiscount(k)

		if err != nil {
			return commands.Checkout{}, err
		}

		prdResp := mapProductResponse(prod, v, disc)

		checkout.TotalAmount = checkout.TotalAmount + prdResp.TotalAmount
		checkout.TotalDiscount = checkout.TotalDiscount + prdResp.Discount
		checkout.TotalAmountWithDiscount = checkout.TotalAmount - checkout.TotalDiscount
		checkout.Products = append(checkout.Products, prdResp)
	}

	if s.isBlackFriday() {
		checkout.Products = append(checkout.Products, s.getGift())
	}

	return checkout, nil
}

func handleProductMap(products []*commands.ProductRequest) map[int]int {
	productsMap := make(map[int]int)
	for _, val := range products {
		productsMap[val.ID] += val.Quantity
	}

	return productsMap
}

func (s service) handleProductAndDiscount(id int) (entities.Product, float32, error) {
	prod, err := s.product.Get(id, false)

	if err != nil {
		return entities.Product{}, 0, err
	}

	disc, err := s.discount.Get(context.Background(), prod.ID)

	if err != nil {
		s.logger.Warning("discount service offline")
	}

	return prod, disc, nil
}

func mapProductResponse(prod entities.Product, qtd int, discount float32) commands.ProductResponse {
	value := float64(prod.Amount) * float64(discount)
	discountAmount := int(math.Round(value))

	return commands.ProductResponse{
		ID:          prod.ID,
		Quantity:    qtd,
		UnitAmount:  prod.Amount,
		TotalAmount: prod.Amount * qtd,
		Discount:    discountAmount * qtd,
		IsGift:      false,
	}
}

func (s service) isBlackFriday() bool {
	return s.conf.BlackFridayDate.Truncate(24*time.Hour) == time.Now().UTC().Truncate(24*time.Hour)
}

func (s service) getGift() commands.ProductResponse {
	prod, _ := s.product.Get(0, true)
	return commands.NewProductGift(prod.ID)
}
