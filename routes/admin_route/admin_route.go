package adminroute

import (
	adminhandler "event_management/internal/handlers/admin_handler"
	adminrepository "event_management/internal/repositories/admin_repositories"
	adminservice "event_management/internal/services/admin_service"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AdminRoutes(e *echo.Group, db *gorm.DB) {
	adminRepo := adminrepository.NewAdminRepositoryImpl(db)
	adminService := adminservice.NewAdminServiceImpl(adminRepo)
	adminHandler := adminhandler.NewAdminHandler(adminService)

	e.POST("/register", adminHandler.RegisterAdmin)
	e.POST("/login", adminHandler.LoginAdmin)

	auth := e.Group("", echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	auth.PUT("/edit", adminHandler.UpdateProfileAdmin)
	auth.GET("/me", adminHandler.GetProfileAdmin)
	auth.POST("/logout", adminHandler.LogoutAdmin)
}
