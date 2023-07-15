package handler

import (
	"fmt"
	"strings"
	"university-api/handler/response"
	interfacesHandler "university-api/interfaces/handler"
	interfacesService "university-api/interfaces/service"

	"github.com/labstack/echo/v4"
)

type ApiFacultyHandler struct {
	service interfacesService.ApiFacultyServiceInterface
}

func NewApiFacultyHandler(
	service interfacesService.ApiFacultyServiceInterface,
) interfacesHandler.ApiFacultyHandlerInterface {
	return &ApiFacultyHandler{
		service: service,
	}
}

func (handler *ApiFacultyHandler) GetFaculty() echo.HandlerFunc {
	return func(c echo.Context) error {
		faculty_cd := strings.ToUpper(c.Param("faculty_cd"))
		tx, errTx := handler.service.DatabaseRepository().Begin()
		if errTx != nil {
			fmt.Println("Get faculty, begin tx err:", errTx)
			return c.JSON(500, "Something wrong")
		}
		defer handler.service.DatabaseRepository().Close()
		faculty, err := handler.service.GetFaculty(tx, faculty_cd)
		fmt.Println("faculty: ", faculty)
		if faculty == nil {
			return c.JSON(404, "not found")
		}
		if err != nil {
			return c.JSON(500, response.Message{
				Error: "Some thing wrong",
				Data:  nil,
			})
		}
		return c.JSON(200, response.Message{
			Error: "Success",
			Data:  faculty,
		})
	}
}
