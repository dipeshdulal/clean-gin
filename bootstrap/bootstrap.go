package bootstrap

import (
	"context"
	"fmt"

	"github.com/dipeshdulal/clean-gin/controllers"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/routers"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	routers.Module,
	lib.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(lifecycle fx.Lifecycle, handler lib.RequestHandler) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting Application")
			fmt.Println("---------------------")
			fmt.Println("------- CLEAN -------")
			fmt.Println("---------------------")
			go handler.Gin.Run(":5000")
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Stopping Application")
			return nil
		},
	})
}
