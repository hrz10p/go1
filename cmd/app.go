package main

import (
	"database/sql"
	"fmt"
	"main/pkg/services"
	"main/pkg/utils/logger"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Application struct {
	Service *services.Service
	Router  *http.ServeMux
	Logger  *logger.Logger
}

func NewApplication() *Application {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Print(err)
	}
	router := http.NewServeMux()

	return &Application{
		Service: services.NewService(db),
		Router:  router,
		Logger:  logger.GetLogger(),
	}
}

func (app *Application) Start(addr string) error {
	app.Logger.Info("starting server... on localhost:8080")
	app.InitializeRoutes()
	return http.ListenAndServe(addr, app.Router)
}
