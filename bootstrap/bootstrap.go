package bootstrap

import (
	"context"
	"fmt"

	"github.com/dipeshdulal/clean-gin/controllers"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/routers"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	routers.Module,
	lib.Module,
	services.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routers.Routes,
	env lib.Env,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting Application")
			fmt.Println("---------------------")
			fmt.Println("------- CLEAN -------")
			fmt.Println("---------------------")

			godotenv.Load()
			env.LoadEnv()

			go func() {
				routes.Setup()
				handler.Gin.Run(env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			fmt.Println("Stopping Application")
			return nil
		},
	})
}
