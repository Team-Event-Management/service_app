package eventroute

import (
	datasources "event_management/internal/dataSources"
	eventhandler "event_management/internal/handlers/event_handler"
	eventrepository "event_management/internal/repositories/event_repositories"
	imagerepository "event_management/internal/repositories/image_repositories"
	eventservice "event_management/internal/services/event_service"
	imageservice "event_management/internal/services/image_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EventRoutes(e *echo.Group, db *gorm.DB, cloudinarySvc *datasources.CloudinaryService) {
	eventRepo := eventrepository.NewEventRepositoryImpl(db)
	imageRepo := imagerepository.NewImageRepositoryImpl(db)

	eventSvc := eventservice.NewEventServiceImpl(eventRepo)
	imageSvc := imageservice.NewImageServiceImpl(eventRepo, imageRepo, *cloudinarySvc)

	handler := eventhandler.NewEventHandler(eventSvc, imageSvc)

	e.POST("/create", handler.CreateEvent)
	e.GET("/all", handler.GetAllEvent)
	e.GET("/:eventId", handler.GetByIdEvent)
	e.PUT("/:eventId/edit", handler.UpdateEvent)
	e.DELETE("/:eventId/delete", handler.DeleteEvent)

	e.GET("/:eventId/images", handler.ListEventImages)
	e.DELETE("/:eventId/images/:imageId", handler.DeleteEventImage)
}
