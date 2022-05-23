package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/monferon/fsm/loader/internal/usecase"
	file "github.com/monferon/fsm/loader/pkg/grpc"
)

func NewRouter(e *echo.Echo, f *usecase.FileInfoUseCase, c file.SenderClient) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ee := e.Group("/api")
	{
		newFileRoutes(ee, f, c)
	}
}
