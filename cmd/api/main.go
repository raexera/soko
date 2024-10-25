package main

import (
	echo "github.com/labstack/echo/v4"

	"github.com/raexera/soko/internal/app/handlers"
	"github.com/raexera/soko/internal/app/storage"
)

func main() {
	e := echo.New()

	storage.InitDB()

	e.POST("/products", handlers.CreateProduct)
	e.PUT("/products/:id", handlers.UpdateProduct)
	e.GET("/products/:id", handlers.GetProduct)
	e.GET("/products", handlers.GetAllProducts)
	e.DELETE("/products/:id", handlers.DeleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
