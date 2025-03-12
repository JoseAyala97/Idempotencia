package database

import (
	"IdEmpotencia/internal/order"
	"IdEmpotencia/internal/orderitem"
	"IdEmpotencia/internal/product"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	config := NewDbConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Name)

	var err error
	for i := 0; i < 5; i++ {
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			fmt.Println("Conectado a MySQL con GORM")
			autoMigrateModels()
			SeedData()
			return
		}
		fmt.Println("⏳ Esperando MySQL... Reintentando en 5s")
		time.Sleep(5 * time.Second)
	}
	log.Panicf("Error conectando a MySQL: %v", err)
}

func Init() {
	InitRedis()
	InitDB()
	fmt.Println("Redis y MySQL inicializados")
}

func autoMigrateModels() {
	err := DB.AutoMigrate(
		&product.Products{},
		&order.Order{},
		&orderitem.OrderItem{},
	)
	if err != nil {
		log.Panicf("Error en la migración: %v", err)
	}
}
