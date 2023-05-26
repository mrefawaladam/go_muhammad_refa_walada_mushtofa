package http

import (
	db "capston-lms/internal/adapters/db/mysql"
	handler "capston-lms/internal/adapters/http/handler"
	middlewares "capston-lms/internal/adapters/http/middleware"
	repository "capston-lms/internal/adapters/repository"
	usecase "capston-lms/internal/application/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// user management
	userRepo    repository.UserRepository
	userHandler handler.UserHandler
	userUsecase usecase.UserUseCase
	// user management
	itemRepo    repository.ItemRepository
	itemHandler handler.ItemsHandler
	itemUsecase usecase.ItemUseCase
	// auth
	AuthHandler handler.AuthHandler
)

func declare() {
	// user
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUseCase{Repo: userRepo}
	userHandler = handler.UserHandler{UserUsecase: userUsecase}
	// item
	itemRepo = repository.ItemRepository{DB: db.DbMysql}
	itemUsecase = usecase.ItemUseCase{Repo: itemRepo}
	itemHandler = handler.ItemsHandler{ItemUseCase: itemUsecase}
	// auth
	AuthHandler = handler.AuthHandler{Usecase: userUsecase}

}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()
	e := echo.New()
	LogMiddleware(e)
	e.Use(middleware.Logger())
	e.POST("/login", AuthHandler.Login())
	e.POST("/registrasi", AuthHandler.Register())

	// montor group
	admin := e.Group("/admin")
	admin.Use(middleware.Logger())
	admin.Use(middlewares.AuthMiddleware())
	admin.Use(middlewares.RequireRole("admin"))
	// users
	admin.GET("/users", userHandler.GetAllUsers())
	admin.GET("/users/:id", userHandler.GetUser())
	admin.POST("/users", userHandler.CreateUser())
	admin.DELETE("/users/:id", userHandler.DeleteUser())
	// items
	admin.GET("/item", itemHandler.GetItems())
	admin.GET("/item/category/:category_id", itemHandler.GetItemByCategory())

	admin.GET("/item/:id", itemHandler.GetItem())
	admin.POST("/item", itemHandler.CreateItem())
	admin.PUT("/item/:id", itemHandler.UpdateItem())
	admin.DELETE("/item/:id", itemHandler.DeleteItem())

	return e
}
func LogMiddleware(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri},latency_human=${latency_human} ,status=${status}\n",
	}))
}
