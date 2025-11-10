package instanceroute

import (
	instancehandler "event_management/internal/handlers/instance_handler"
	instancerepository "event_management/internal/repositories/instance_repositories"
	instanceservice "event_management/internal/services/instance_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InstanceRoutes(e *echo.Group, db *gorm.DB) {
	instanceRepo := instancerepository.NewInstanceRepositoryImpl(db)
	instanceService := instanceservice.NewInstanceServiceImpl(instanceRepo)
	instanceHandler := instancehandler.NewInstanceHandler(instanceService)

	e.POST("/create", instanceHandler.CreateInstance)
	e.GET("/all", instanceHandler.GetAllInstance)
	e.GET("/:instanceId", instanceHandler.GetByIdInstance)
	e.PUT("/:instanceId/edit", instanceHandler.UpdateInstance)
	e.DELETE("/:instanceId/delete", instanceHandler.DeleteInstance)
}