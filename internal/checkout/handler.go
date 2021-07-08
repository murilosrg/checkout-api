package checkout

import (
	"encoding/json"
	"github.com/murilosrg/checkout-api/internal/commands"
	"github.com/murilosrg/checkout-api/internal/errors"
	"github.com/murilosrg/checkout-api/pkg/log"
	"net/http"
)

type Handler interface {
	Post(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	checkout Service
	logger   log.Logger
}

func NewHandler(checkout Service, logger log.Logger) Handler {
	return &handler{checkout, logger}
}

func (h handler) Post(w http.ResponseWriter, r *http.Request) {
	var cart commands.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)

	if err := cart.Validate(); err != nil {
		SendError(w, errors.BadRequest(err.Error()))
		return
	}

	res, err := h.checkout.Checkout(cart)

	if err != nil {
		SendError(w, errors.UnprocessableEntity(err.Error()))
		return
	}

	SendJson(w, res)
}

func SendError(w http.ResponseWriter, err errors.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode())
	json.NewEncoder(w).Encode(err)
}

func SendJson(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payload)
}

