package routes

import (
	adminroute "event_management/routes/admin_route"
	instanceroute "event_management/routes/instance_route"
	roleroute "event_management/routes/role_route"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	v1 := e.Group("/api/v1")
	roleroute.RoleRoutes(v1.Group("/role"), db)
	adminroute.AdminRoutes(v1.Group("/admin"), db)
	instanceroute.InstanceRoutes(v1.Group("/instance"), db)
}