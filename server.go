package main

import (
	"github.com/dipeshdulal/clean-gin/bootstrap"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
