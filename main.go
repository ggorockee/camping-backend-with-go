package main

import (
	"camping-backend-with-go/pkg/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := databaseConnection()

	app.Get("/v1/helloworld", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status": "ok",
			"error":  false,
		})
	})

	app.Listen(":3000")
}

func databaseConnection() *gorm.DB {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	err = db.AutoMigrate(&entities.Spot{})
	if err != nil {
		log.Println(err.Error())
	}

	return db
}
