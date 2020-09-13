package main

import (
	"github.com/dipeshdulal/clean-gin/bootstrap"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	_ "github.com/go-sql-driver/mysql"

)

func main() {
	godotenv.Load()
	fx.New(bootstrap.Module).Run()
}
