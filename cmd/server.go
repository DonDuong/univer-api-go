package cmd

import (
	"university-api/domain/repository"
	interfacesHandler "university-api/interfaces/handler"
	interfacesService "university-api/interfaces/service"

	"github.com/labstack/echo/v4"
)

type ApiServer struct {
	DatabaseRepository repository.DatabaseInterface

	ApiFacultyInterface repository.ApiFacultyInterface

	ApiFacultyServiceInterface interfacesService.ApiFacultyServiceInterface
	ApiFacultyHandlerInterface interfacesHandler.ApiFacultyHandlerInterface

	pgSqlDn     string
	pgSqlUser   string
	pgSqlPass   string
	pgSqlDbname string

	echo *echo.Echo
}

func (server *ApiServer) Start() {
	server.echo = echo.New()
	server.DependenciesInjection()
	server.route()

	server.echo.Logger.Fatal(server.echo.Start(":9090"))
}
