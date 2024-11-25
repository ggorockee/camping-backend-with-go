package main

import (
	"camping-backend-with-go/api/routes"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/healthcheck"
	"camping-backend-with-go/pkg/spot"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := databaseConnection()

	healthcheckService := healthcheck.NewService()

	spotRepo := spot.NewRepo(db)
	spotService := spot.NewService(spotRepo)

	app := fiber.New()
	app.Use(cors.New())

	v1 := app.Group("/v1")
	routes.SpotRouter(v1, spotService)
	routes.HealthCheckRouter(v1, healthcheckService)
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() *gorm.DB {
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	err = db.AutoMigrate(
		&entities.Spot{},
		&entities.User{},
	)
	if err != nil {
		log.Println(err.Error())
	}

	return db
}
