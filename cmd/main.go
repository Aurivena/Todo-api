package main

import (
	"Todo/initializate"
	"Todo/models"
	"Todo/pkg/action"
	"Todo/pkg/domain"
	"Todo/pkg/persistence"
	"Todo/pkg/route"
	"Todo/server"
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	logrus.Info("start init server")
	if err := initializate.Load(); err != nil {
		logrus.Fatal(err.Error())
	}
	logrus.Info("end init server")
}

// @title           TodoService
// @version         1.0.0
// @description     Задачки

// @host      		localhost:1941
// @BasePath  		/api/
func main() {
	var serverInstance server.Server

	businessDatabase := persistence.NewBusinessDatabase(initializate.ConfigService)
	sources := persistence.Sources{
		BusinessDB: businessDatabase,
	}
	persistences := persistence.NewPersistence(&sources)
	domains := domain.NewDomain(persistences, initializate.ConfigService)
	actions := action.NewAction(domains)
	routes := route.NewRoute(actions)
	go run(serverInstance, routes, &initializate.ConfigService.Server)
	stop()
	serverInstance.Stop(context.Background(), businessDatabase)
}

func run(server server.Server, routes *route.Route, config *models.ServerConfig) {
	ginEgine := routes.InitHTTPRoutes(config)

	if err := server.Run(config.Port, ginEgine); err != nil {
		if err.Error() != "http: Server closed" {
			logrus.Fatalf("error occurred while running http server: %s", nil, err.Error())
		}
	}
}

func stop() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGABRT)
	<-quit
}
