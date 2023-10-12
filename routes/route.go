package routes

import (
	"NOMOR1/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.GET("/books",controllers.GetAllBook)
	e.GET("/books/:id",controllers.GetBookById)
	e.POST("/books",controllers.CreateBook)
	e.PUT("/books/:id",controllers.UpdateBook)
	e.DELETE("/books/:id",controllers.DeleteBook)

	return e

}