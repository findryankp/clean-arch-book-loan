package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	// middlewares.BasicLogger(e)
	// _a.AuthRouter(db, e)
	UserRouter(db, e)
	// _b.BookRouter(db, e)
}
