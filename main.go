package main

import (
	"github.com/joho/godotenv"
	"github.com/karim-w/techunicorn-task/config/dbconfig"
	"github.com/karim-w/techunicorn-task/controllers"
	"github.com/karim-w/techunicorn-task/router"
	"github.com/karim-w/techunicorn-task/service"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	app := fx.New(
		dbconfig.DBModule,
		service.AuthServiceModule,
		service.DocServiceModule,
		controllers.DocControllerModule,
		controllers.AuthControllerModule,
		router.Module,
	)
	defer app.Run()
}
