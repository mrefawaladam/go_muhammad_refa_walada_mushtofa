package routes

import (
	"curdusers/constans"
	"curdusers/controllers"
	m "curdusers/middlewares"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	// users routes
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.POST("/login", controllers.LoginUserController)

	// blogs routes
	e.POST("/blogs", controllers.CreateBlogController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)

	// Book routes
	e.POST("/books", controllers.CreateBookController)
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)

	eAuthBasic := e.Group("/auth")
	eAuthBasic.Use(mid.BasicAuth(m.BasicAuthDB))
	eAuthBasic.GET("/users", controllers.GetUsersController)

	eJwt := e.Group("/jwt")
	eJwt.Use(mid.JWT([]byte(constans.SECRET_JWT)))
	eJwt.GET("/users", controllers.GetUsersController)

	return e
}
