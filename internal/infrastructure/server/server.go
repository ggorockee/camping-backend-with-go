package server

import (
	amenityrepository "camping-backend-with-go/internal/domain/repository/amenity"
	authrepository "camping-backend-with-go/internal/domain/repository/auth"
	categoryrepository "camping-backend-with-go/internal/domain/repository/category"
	amenityservice "camping-backend-with-go/internal/domain/service/amenity"
	authservice "camping-backend-with-go/internal/domain/service/auth"
	categoryservice "camping-backend-with-go/internal/domain/service/category"
	"camping-backend-with-go/internal/infrastructure/database"
	amenityroute "camping-backend-with-go/pkg/api/route/amenity"
	authroute "camping-backend-with-go/pkg/api/route/auth"
	categoryroute "camping-backend-with-go/pkg/api/route/category"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

var port int

func Start(aport ...int) {
	// port 인수 제어
	if len(aport) == 0 {
		port = 3000
	}

	if len(aport) == 1 {
		port = aport[0]
	} else {
		log.Fatalf("port는 정확히 한개여야 합니다.\n")
	}

	db := database.Connection()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${red}${status} ${blue}[${method}] ${white}${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Asia/Seoul",
	}))

	// Middleware로 db pointer 보내기
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	// repository 초기화
	authRepo := authrepository.NewAuthRepository(db)
	categoryRepo := categoryrepository.NewCategoryRepository(db)
	amenityRepo := amenityrepository.NewAmenityRepository(db)

	// service 초기화
	authService := authservice.NewAuthService(authRepo)
	categoryService := categoryservice.NewCategoryService(categoryRepo)
	amenityService := amenityservice.NewAmenityService(amenityRepo)

	// attach router
	v1 := app.Group("/api/v1")
	authroute.AuthRouter(v1, authService)
	categoryroute.CategoryRouter(v1, categoryService)
	amenityroute.AmenityRouter(v1, amenityService)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
