package main

import (
	_ "camping-backend-with-go/docs"
	entities2 "camping-backend-with-go/internal/domain"
	"camping-backend-with-go/internal/middleware"
	"camping-backend-with-go/internal/repository"
	"camping-backend-with-go/internal/route"
	"camping-backend-with-go/internal/service"
	"camping-backend-with-go/pkg/config"
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

// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	db := databaseConnection()

	healthcheckService := healthcheck.NewService()

	userRepo := user.NewRepo(db)
	amenityRepo := amenity.NewRepo(db, userRepo)
	categoryRepo := category.NewRepo(db, userRepo)
	spotRepo := spot.NewRepo(db, userRepo, amenityRepo, categoryRepo)
	wishRepo := repository.NewRepo(db, spotRepo)

	userService := user.NewService(userRepo)
	spotService := spot.NewService(spotRepo)
	categoryService := category.NewService(categoryRepo)
	amenityService := amenity.NewService(amenityRepo)
	wishListService := service.NewService(wishRepo)

	app := fiber.New()
	app.Use(cors.New())

	// db instance를 middleware로 보내기
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// request middleware 구현
	app.Use(middleware.RequestAuthMiddleware())

	// swagger settings
	//swaggerCfg := swagger.Config{
	//	BasePath: "/v1",
	//	FilePath: "./docs/swagger.json",
	//	Path:     "docs",
	//}
	//app.Use(swagger.New(swaggerCfg))

	v1 := app.Group("/api/v1")

	route.UserRouter(v1, userService)
	route.AuthRouter(v1, userService)
	route.SpotRouter(v1, spotService)
	route.CategoryRouter(v1, categoryService)
	route.AmenityRouter(v1, amenityService)
	route.WishListRouter(v1, wishListService)
	route.SwaggerRouter(v1)
	route.HealthCheckRouter(v1, healthcheckService)
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
		log.Fatalf("database setting failed...")
	} else {
		dbSetting = "mysql"
		log.Printf("[INFO] %s DB setting enabled...\n", dbSetting)
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
	}

	log.Println("connected")
	err = db.AutoMigrate(
		&entities2.Spot{},
		&entities2.User{},
		&entities2.Category{},
		&entities2.Amenity{},
		&entities2.Review{},
		&entities2.WishList{},
	)
	if err != nil {
		log.Println(err.Error())
	}

	return db
}
