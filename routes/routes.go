package routes

import (
	datasources "event_management/internal/dataSources"
	adminroute "event_management/routes/admin_route"
	eventroute "event_management/routes/event_route"
	instanceroute "event_management/routes/instance_route"
	roleroute "event_management/routes/role_route"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB, cloudinarySvc *datasources.CloudinaryService) {
	v1 := e.Group("/api/v1")
	roleroute.RoleRoutes(v1.Group("/role"), db)
	adminroute.AdminRoutes(v1.Group("/admin"), db)
	instanceroute.InstanceRoutes(v1.Group("/instance"), db)
	eventroute.EventRoutes(v1.Group("/event"), db, cloudinarySvc)
}
