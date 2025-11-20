package standcategoryroute

import (
	standcategoryhandler "event_management/internal/handlers/stand_category_handler"
	standcategoryrepositories "event_management/internal/repositories/stand_category_repositories"
	standcategoryservice "event_management/internal/services/stand_category_service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StandCategoryRoutes(e *echo.Group, db *gorm.DB) {
	standcategoryRepo := standcategoryrepositories.NewStandCategoryRepositoryImpl(db)
	standcategoryService := standcategoryservice.NewStandCategoryServiceImpl(standcategoryRepo)
	standcategoryHandler := standcategoryhandler.NewStandCategoryHandler(standcategoryService)

	e.POST("/create", standcategoryHandler.CreateStandCategory)
	e.GET("/all", standcategoryHandler.GetAllStandCategory)
	e.GET("/:standCategoryId", standcategoryHandler.GetByIdStandCategory)
	e.PUT("/:standCategoryId/edit", standcategoryHandler.UpdateStandCategory)
	e.DELETE("/:standCategoryId/delete", standcategoryHandler.DeleteStandCategory)
}