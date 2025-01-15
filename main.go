package main

import (
	"camping-backend-with-go/app"
	"camping-backend-with-go/app/core/helper"
	"camping-backend-with-go/app/domain/user"
	"camping-backend-with-go/config"
	_ "camping-backend-with-go/docs"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

// @title Dolphindance App
// @version 1.0
// @description This is an API for Dolphindance Application

// @contact.name ggorockee
// @contact.email woohyeon88@daolcompany.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	fx.New(
		config.Module,
		helper.Module,

		user.ControllerModule,

		fx.Provide(
			app.NewFiber,
			fx.Annotate(
				app.NewRouter,
				fx.ParamTags(``, `group:"routes"`),
			),
		),

		fx.Invoke(
			func(fiber.Router) {},
			func(*fiber.App) {},
		),
	).Run()
}
