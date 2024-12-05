package main

import (
	"camping-backend-with-go/api/routes"
	_ "camping-backend-with-go/docs"
	"camping-backend-with-go/pkg/config"
	"camping-backend-with-go/pkg/entities"
	"camping-backend-with-go/pkg/service/amenity"
	"camping-backend-with-go/pkg/service/category"
	"camping-backend-with-go/pkg/service/healthcheck"
	"camping-backend-with-go/pkg/service/spot"
	"camping-backend-with-go/pkg/service/user"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

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
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	db := databaseConnection()

	healthcheckService := healthcheck.NewService()

	userRepo := user.NewRepo(db)
	userService := user.NewService(userRepo)

	spotRepo := spot.NewRepo(db, userRepo)
	spotService := spot.NewService(spotRepo)

	// Category
	categoryRepo := category.NewRepo(db, userRepo)
	categoryService := category.NewService(categoryRepo)

	//Amenity
	amenityRepo := amenity.NewRepo(db, userRepo)
	amenityService := amenity.NewService(amenityRepo)

	app := fiber.New()
	app.Use(cors.New())

	// db instance를 middleware로 보내기
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// swagger settings
	//swaggerCfg := swagger.Config{
	//	BasePath: "/v1",
	//	FilePath: "./docs/swagger.json",
	//	Path:     "docs",
	//}
	//app.Use(swagger.New(swaggerCfg))

	v1 := app.Group("/v1")

	routes.UserRouter(v1, userService)
	routes.AuthRouter(v1, userService)
	routes.SpotRouter(v1, spotService)
	routes.CategoryRouter(v1, categoryService)
	routes.AmenityRouter(v1, amenityService)
	routes.SwaggerRouter(v1)
	routes.HealthCheckRouter(v1, healthcheckService)
	log.Fatal(app.Listen(":3000"))
}

func databaseConnection() *gorm.DB {
	// Local에서 Teleport 작업할 때만 사용
	// 배포시에는 comment 활성화
	MYSQL := os.Getenv("GO_MYSQL")
	var dsn string
	var dbSetting string
	var db *gorm.DB
	var err error

	if MYSQL == "" {
		dbSetting = "local"
		dsn = "test.db"
		log.Printf("[INFO] %s DB setting enabled...\n", dbSetting)
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	} else {
		dbSetting = "mysql"
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Config("DB_USER"),
			config.Config("DB_PASSWORD"),
			config.Config("DB_HOST"),
			config.Config("DB_PORT"),
			config.Config("DB_NAME"),
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	err = db.AutoMigrate(
		&entities.Spot{},
		&entities.User{},
		&entities.Category{},
		&entities.Amenity{},
	)
	if err != nil {
		log.Println(err.Error())
	}

	return db
}
