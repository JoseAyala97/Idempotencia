package injection

import (
	orderHandler "IdEmpotencia/internal/order/handler"
	orderInfra "IdEmpotencia/internal/order/infrastructure"
	orderService "IdEmpotencia/internal/order/service"

	productHandler "IdEmpotencia/internal/product/handler"
	productInfra "IdEmpotencia/internal/product/infrastructure"
	productService "IdEmpotencia/internal/product/service"
)

func InjectDependencies() (*productHandler.ProductHandler, *orderHandler.OrderHandler) {
	productRepo := productInfra.NewProductRepository()
	orderRepo := orderInfra.NewOrderRepository()

	productService := productService.NewProductService(productRepo)
	orderService := orderService.NewOrderService(orderRepo, productRepo)

	productHandler := productHandler.NewProductHandler(productService)
	orderHandler := orderHandler.NewOrderHandler(orderService)

	return productHandler, orderHandler
}
