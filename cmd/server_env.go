package cmd

import (
	"fmt"
	"os"
)

func (server *ApiServer) env() []string {
	var erros []string
	server.pgSqlUser = os.Getenv("DBUSERNAME")
	fmt.Println(server.pgSqlUser)
	if server.pgSqlUser == "" {
		erros = append(erros, "Username not found")
	}
	server.pgSqlPass = os.Getenv("PASSWORD")
	fmt.Println(server.pgSqlPass)

	if server.pgSqlUser == "" {
		erros = append(erros, "Password not found")
	}

	server.pgSqlDbname = os.Getenv("DBNAME")

	fmt.Println(server.pgSqlDbname)

	if server.pgSqlUser == "" {
		erros = append(erros, "DBname not found")
	}

	return erros
}
