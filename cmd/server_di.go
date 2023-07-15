package cmd

import (
	"university-api/handler"
	"university-api/infrastructure"
	"university-api/services"
)

func (server *ApiServer) DependenciesInjection() {
	server.DatabaseRepository = infrastructure.NewDatabase(
		infrastructure.DATABASE_DRIVER_PGSQL,
		server.pgSqlDn,
		server.pgSqlUser,
		server.pgSqlPass,
		server.pgSqlDbname,
	)

	server.ApiFacultyServiceInterface = services.NewApiFacultyService(server.DatabaseRepository)
	server.ApiFacultyHandlerInterface = handler.NewApiFacultyHandler(server.ApiFacultyServiceInterface)
}
