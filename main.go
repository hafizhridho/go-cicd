package main

import (
	"latihan/configs"
	"latihan/controllers"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.Loadenv()
	configs.InitDb()
	e := echo.New()
	e.GET("/books", controllers.GetAll)
	e.POST("/books", controllers.CreateBook)
	e.DELETE("/books/:id",controllers.DeleteController)
	e.GET("/books/:id", controllers.GetByID)
	e.PUT("books/:id", controllers.UpdateBook)
	 e.Start(":8000")
}