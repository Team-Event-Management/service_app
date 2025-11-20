package standroute

import (
	standhandler "event_management/internal/handlers/stand_handler"
	standrepositories "event_management/internal/repositories/stand_repositories"
	standservice "event_management/internal/services/stand_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StandRoutes(e *echo.Group, db *gorm.DB) {
	standRepo := standrepositories.NewStandRepositoryImpl(db)
	standService := standservice.NewStandServiceImpl(standRepo)
	standHandler := standhandler.NewStandHandler(standService)

	e.POST("/create", standHandler.CreateStand)
	e.GET("/all", standHandler.GetAllStand)
	e.GET("/:standId", standHandler.GetByIdStand)
	e.PUT("/:standId/edit", standHandler.UpdateStand)
	e.DELETE("/:standId/delete", standHandler.DeleteStand)
}
