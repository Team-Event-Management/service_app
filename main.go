package main

import (
	"event_management/configs"
	datasources "event_management/internal/dataSources"
	"event_management/pkg/workers/producer"
	"event_management/routes"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	configs.LoadEnv()

	db := configs.InitDB()

	configs.RunMigrations(db)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH},
	}))

	for _, r := range e.Routes() {
		log.Printf("ROUTE %s %s", r.Method, r.Path)
	}

	cloudinarySvc, err := datasources.NewCloudinaryService()
	if err != nil {
		log.Fatalf("Failed to initialize Cloudinary service: %v", err)
	}

	configs.InitRabbitMQ()
	defer configs.CloseConnections()

	go producer.StartWorker()

	routes.Routes(e, db, &cloudinarySvc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(e.Start(":" + port))
}
