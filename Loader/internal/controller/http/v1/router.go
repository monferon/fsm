package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/monferon/fsm/loader/internal/usecase"
)

//func upload(c echo.Context) error {
//
//}

func NewRouter(e *echo.Echo, f *usecase.FileInfoUseCase) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ee := e.Group("/api")
	{
		newFileRoutes(ee, f)
		//ee.POST("/upload", upload)
	}
}
