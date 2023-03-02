package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB, e *echo.Echo) {
	// data := _userData.New(db)
	// service := _userService.New(data)
	// handler := _userHandler.New(service)

	// g := e.Group("/users")
	// g.GET("", handler.GetAll)
	// g.GET("/:id", handler.GetById)
	// g.POST("", handler.Create)
	// g.PUT("/:id", handler.Create)
	// g.DELETE("/:id", handler.Create)

}
