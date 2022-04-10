package main

import (
	"github.com/dipeshdulal/clean-gin/bootstrap"
	"github.com/dipeshdulal/clean-gin/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()

	logger := lib.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
