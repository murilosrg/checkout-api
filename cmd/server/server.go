package main

import (
	"github.com/gorilla/mux"
	"github.com/murilosrg/checkout-api/internal/checkout"
	"github.com/murilosrg/checkout-api/internal/config"
	"github.com/murilosrg/checkout-api/internal/discount"
	"github.com/murilosrg/checkout-api/internal/products"
	"github.com/murilosrg/checkout-api/pkg/db"
	"github.com/murilosrg/checkout-api/pkg/log"
	"google.golang.org/grpc"
	"net/http"
	"time"
)

func main() {
	conf, err := config.NewConfig(".env")
	check(err)

	handle := handleServices(conf)
	router := buildHandler(handle)

	http.HandleFunc("/", router.ServeHTTP)
	err = http.ListenAndServe(":8080", nil)
	check(err)
}

func handleServices(conf config.Config) checkout.Handler {
	DB, err := db.InitializeDb(conf.DatabaseFile)
	check(err)

	logger := log.New()

	conn, _ := grpc.Dial(conf.Discount, grpc.WithInsecure())

	productRepo := products.NewRepository(DB, logger)
	productService := products.NewService(productRepo, logger)
	discountService := discount.NewService(conn, time.Second, logger)
	checkoutService := checkout.NewService(conf, productService, discountService, logger)
	checkoutHandler := checkout.NewHandler(checkoutService, logger)

	return checkoutHandler
}

func buildHandler(handler checkout.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/checkout", handler.Post).Methods("POST")
	return router
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
