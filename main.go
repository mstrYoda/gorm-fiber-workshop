package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"goworkshop/handler"
	"goworkshop/repository"
	"goworkshop/server"
)

func main() {
	dsn := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s",
		"postgres",
		"postgres",
		"123456",
		"localhost",
		"5432")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&repository.Product{})
	productRepo := repository.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepo)
	srv := server.NewServer(productHandler)

	srv.Listen(":8080")
}
