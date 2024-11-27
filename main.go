package main

import (
	"camping-backend-with-go/api/routes"
	"camping-backend-with-go/pkg/auth"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/healthcheck"
	"camping-backend-with-go/pkg/proxy"
	"camping-backend-with-go/pkg/spot"
	"camping-backend-with-go/pkg/user"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

// @title ggocamping App
// @version 1.0
// @description This is an API for ggocamping Application

// @contact.name ggorockee
// @contact.email ggorockee@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /v1
func main() {
	db := databaseConnection()

	healthcheckService := healthcheck.NewService()

	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo)

	authRepo := auth.NewRepo(db)
	authService := auth.NewService(authRepo)

	spotRepo := spot.NewRepo(db)
	spotService := spot.NewService(spotRepo)

	app := fiber.New()
	app.Use(cors.New())

	// swagger settings
	swaggerCfg := swagger.Config{
		BasePath: "/v1",
		FilePath: "./docs/swagger.yaml",
		Path:     "docs",
	}
	app.Use(swagger.New(swaggerCfg))

	v1 := app.Group("/v1")

	routes.UserRouter(v1, userService)
	routes.AuthRouter(v1, authService)

	routes.SpotRouter(v1, spotService)
	routes.HealthCheckRouter(v1, healthcheckService)
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() *gorm.DB {
	// Local에서 Teleport 작업할 때만 사용
	// 배포시에는 comment 활성화
	if err := os.Setenv("PROXY", "true"); err != nil {
		log.Println(err.Error())
	}

	dsn := proxy.GetProxyDatabase()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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
