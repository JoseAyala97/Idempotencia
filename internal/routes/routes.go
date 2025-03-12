package routes

import (
	"github.com/gorilla/mux"

	orderHandler "IdEmpotencia/internal/order/handler"
	productHandler "IdEmpotencia/internal/product/handler"
)

func RegisterRoutes(
	r *mux.Router,
	productHandler *productHandler.ProductHandler,
	orderHandler *orderHandler.OrderHandler,
) {
	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}/stock", productHandler.UpdateStock).Methods("PUT")

	r.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{id}", orderHandler.GetOrder).Methods("GET")
}

func InitRouter(productHandler *productHandler.ProductHandler, orderHandler *orderHandler.OrderHandler) *mux.Router {
	r := mux.NewRouter()
	RegisterRoutes(r, productHandler, orderHandler)
	return r
}
