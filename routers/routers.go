package routers

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Invoke(registerRoutes),
)
