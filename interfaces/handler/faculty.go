package interfaces

import "github.com/labstack/echo/v4"

type ApiFacultyHandlerInterface interface {
	GetFaculty() echo.HandlerFunc
}
