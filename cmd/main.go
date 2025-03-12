package main

import (
	"fmt"
	"log"
	"net/http"

	"IdEmpotencia/internal/routes"

	"IdEmpotencia/pkg/database"
	"IdEmpotencia/pkg/injection"
	"IdEmpotencia/pkg/middleware"
)

func main() {
	database.Init()

	productHandler, orderHandler := injection.InjectDependencies()

	r := routes.InitRouter(productHandler, orderHandler)
	r.Use(middleware.ErrorHandler)

	port := ":8080"
	fmt.Println("Servidor corriendo en", port)
	log.Fatal(http.ListenAndServe(port, r))
}
