package eventroute

import (
	eventhandler "event_management/internal/handlers/event_handler"
	eventrepository "event_management/internal/repositories/event_repositories"
	imagerepository "event_management/internal/repositories/image_repositories"
	eventservice "event_management/internal/services/event_service"
	imageservice "event_management/internal/services/image_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EventRoutes(e *echo.Group, db *gorm.DB) {
	eventRepo := eventrepository.NewEventRepositoryImpl(db)
	imageRepo := imagerepository.NewImageRepositoryImpl(db)

	eventSvc := eventservice.NewEventServiceImpl(eventRepo)
	imageSvc := imageservice.NewImageServiceImpl(eventRepo, imageRepo, "uploads")

	handler := eventhandler.NewEventHandler(eventSvc, imageSvc)

	e.POST("/create", handler.CreateEvent)
	e.GET("/all", handler.GetAllEvent)
	e.GET("/:eventId", handler.GetByIdEvent)
	e.PUT("/:eventId/edit", handler.UpdateEvent)
	e.DELETE("/:eventId/delete", handler.DeleteEvent)

	e.POST("/:eventId/upload-image", handler.UploadEventImage)
	e.GET("/:eventId/images", handler.ListEventImages)
	e.DELETE("/:eventId/images/:imageId", handler.DeleteEventImage)
}