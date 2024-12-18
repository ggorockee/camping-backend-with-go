package server

import (
	amenityrepository "camping-backend-with-go/internal/domain/repository/amenity"
	authrepository "camping-backend-with-go/internal/domain/repository/auth"
	categoryrepository "camping-backend-with-go/internal/domain/repository/category"
	spotrepository "camping-backend-with-go/internal/domain/repository/spot"
	userrepository "camping-backend-with-go/internal/domain/repository/user"
	amenityservice "camping-backend-with-go/internal/domain/service/amenity"
	authservice "camping-backend-with-go/internal/domain/service/auth"
	categoryservice "camping-backend-with-go/internal/domain/service/category"
	healthcheckservice "camping-backend-with-go/internal/domain/service/healthcheck"
	spotservice "camping-backend-with-go/internal/domain/service/spot"
	userservice "camping-backend-with-go/internal/domain/service/user"
	"camping-backend-with-go/internal/infrastructure/database"
	amenityroute "camping-backend-with-go/pkg/api/route/amenity"
	authroute "camping-backend-with-go/pkg/api/route/auth"
	categoryroute "camping-backend-with-go/pkg/api/route/category"
	healthcheckroute "camping-backend-with-go/pkg/api/route/healthcheck"
	spotroute "camping-backend-with-go/pkg/api/route/spot"
	swaggerroute "camping-backend-with-go/pkg/api/route/swagger"
	userroute "camping-backend-with-go/pkg/api/route/user"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	userRepo := userrepository.NewUserRepository(db)
	categoryRepo := categoryrepository.NewCategoryRepository(db)
	amenityRepo := amenityrepository.NewAmenityRepository(db)
	spotRepo := spotrepository.NewSpotRepository(db, userRepo, categoryRepo, amenityRepo)

	// service 초기화
	authService := authservice.NewAuthService(authRepo)
	userService := userservice.NewUserService(userRepo)
	categoryService := categoryservice.NewCategoryService(categoryRepo)
	amenityService := amenityservice.NewAmenityService(amenityRepo)
	spotService := spotservice.NewSpotService(spotRepo)

	healthcheckService := healthcheckservice.NewHealthCheckService()

	// attach router
	v1 := app.Group("/api/v1")
	authroute.AuthRouter(v1, authService)
	userroute.UserRouter(v1, userService)
	categoryroute.CategoryRouter(v1, categoryService)
	amenityroute.AmenityRouter(v1, amenityService)
	spotroute.SpotRouter(v1, spotService)
	swaggerroute.SwaggerRouter(v1)

	healthcheckroute.HealthCheckRouter(v1, healthcheckService)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
