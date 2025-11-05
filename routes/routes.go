package routes

import (
	roleroute "event_management/routes/role_route"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	v1 := e.Group("/api/v1")
	roleroute.RoleRoutes(v1.Group("/role"), db)
}