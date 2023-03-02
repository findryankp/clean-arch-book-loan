package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_ "clean-arch/features/users/data"
	_ "clean-arch/features/users/delivery"
	_ "clean-arch/features/users/service"
)

func AuthRouter(db *gorm.DB, e *echo.Echo) {
	// data := _userData.New(db)
	// service := _userService.New(data)
	// handler := _userHandler.New(service)

	// g := e.Group("/auth")
	// g.POST("/register", handler.Create)
	// g.POST("/login", handler.Create)
	// g.GET("/users", handler.Create)

	// g.POST("/change-password", handler.Create)
	// g.POST("/forget-password", handler.Create)
}
