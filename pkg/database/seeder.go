package database

import (
	"IdEmpotencia/internal/product"
	"fmt"
	"log"
)

func SeedData() {
	var count int64
	if err := DB.Model(&product.Products{}).Count(&count).Error; err != nil {
		log.Panicf("Error verificando productos: %v", err)
	}

	if count > 0 {
		fmt.Println("Seeders: Los productos ya están en la base de datos, saltando inserciones.")
		return
	}

	products := []product.Products{
		{Name: "Laptop Gamer", Price: 1200.00, Stock: 10},
		{Name: "Mouse Inalámbrico", Price: 50.99, Stock: 50},
		{Name: "Teclado Mecánico", Price: 89.99, Stock: 30},
		{Name: "Monitor 27''", Price: 300.50, Stock: 15},
	}

	if err := DB.Create(&products).Error; err != nil {
		log.Panicf("Error insertando productos iniciales: %v", err)
	}

	fmt.Println("Seeders ejecutados: Productos insertados correctamente.")
}
