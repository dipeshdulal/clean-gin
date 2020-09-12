package routers

import "github.com/dipeshdulal/clean-gin/lib"

func registerRoutes(handler lib.RequestHandler) {
	api := handler.Gin.Group("/api")
	{
		api.GET("/")
	}
}
