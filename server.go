package main

import (
	"github.com/dipeshdulal/clean-gin/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
