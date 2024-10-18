package main

import (
	"game-shop-api/controller"
	"game-shop-api/db"
	"game-shop-api/repository"
	"game-shop-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnction, err := db.ConnectDb()
	if err != nil {
		panic(err)
	}

	ProductRespository := repository.NewProductRepository(dbConnction)
	productUsecase := usecase.NewProductUsecase(ProductRespository)
	ProductController := controller.NewProductController(productUsecase)

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:id", ProductController.GetProductsById)
	server.DELETE("product/:id", ProductController.DeleteProducts)

	server.Run(":8083")
}
