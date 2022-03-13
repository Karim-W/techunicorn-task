package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/karim-w/techunicorn-task/controllers"
	"go.uber.org/fx"
)

func SetupRoutes(ac *controllers.AuthController, dc *controllers.DocController) *gin.Engine {
	r := gin.Default()
	base := r.Group("/api/v1")
	ac.SetupRoutes(base)
	docs := r.Group("/doctors")
	dc.SetupRoutes(docs)
	r.Run()
	return r
}
func registerHooks(lifecycle fx.Lifecycle, ginRouter *gin.Engine) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			println("Initializing server")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			println("Terminating server")
			return nil
		},
	})
}

var Module = fx.Options(fx.Provide(SetupRoutes), fx.Invoke(registerHooks))
