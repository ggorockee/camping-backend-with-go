package app

import (
	"camping-backend-with-go/config"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewFiber(lc fx.Lifecycle, c *config.Config) *fiber.App {
	app := initializeFiber()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			addr := fmt.Sprintf(":%s", c.Server.Port)
			go func() {
				log.Println("Starting Fiber Server", zap.String("address", addr))
				if err := app.Listen(addr); err != nil {
					log.Println("fiber server error", zap.Error(err))
				}
			}()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(100 * time.Millisecond):
				log.Println("Fiber server started successfully")
			}

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down Fiber server")
			return app.Shutdown()
		},
	})

	return app
}

func initializeFiber() *fiber.App {
	app := fiber.New()
	app.Get("/api/v1/healthcheck", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	app.Get("/api/v1/docs/*", swagger.HandlerDefault)

	app = setMiddleware(app)

	return app
}

func setMiddleware(app *fiber.App) *fiber.App {
	app.Use(cors.New())
	return app
}
