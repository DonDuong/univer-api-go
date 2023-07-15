package cmd

func (server *ApiServer) route() {
	server.echo.GET("/faculty/:faculty_cd", server.ApiFacultyHandlerInterface.GetFaculty())
}
