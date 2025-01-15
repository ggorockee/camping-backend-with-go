package user

import (
	"camping-backend-with-go/app"

	"go.uber.org/fx"
)

var ControllerModule = fx.Module(
	"controller",
	fx.Provide(
		NewRepository,
		app.AsRoute(NewController),
		NewService,
	),
)
