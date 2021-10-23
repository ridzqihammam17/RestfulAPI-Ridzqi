package routes

import (
	"restfulapi/ridzqi/constants"
	"restfulapi/ridzqi/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/login", controllers.LoginUsersController)
	e.POST("/users", controllers.RegisterUserController)

	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)

	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users", controllers.GetUsersController)
	r.GET("/users/:id", controllers.GetUserController)
	r.PUT("/users/:id", controllers.UpdateUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)

	r.POST("/books", controllers.CreateBookController)
	r.PUT("/books/:id", controllers.UpdateBookController)
	r.DELETE("/books/:id", controllers.DeleteBookController)

	return e
}
