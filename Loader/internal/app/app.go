package app

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/monferon/fsm/loader/config"
	v1 "github.com/monferon/fsm/loader/internal/controller/http/v1"
	"github.com/monferon/fsm/loader/internal/usecase"
	"github.com/monferon/fsm/loader/internal/usecase/repo"
	"github.com/monferon/fsm/loader/internal/usecase/webapi"
	"github.com/monferon/fsm/loader/pkg/postgres"
)

func Run(cfg *config.Config) {

	//database
	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pg.DB)
	//use case
	// Use case
	fileUseCases := usecase.New(
		repo.New(pg),
		webapi.New(),
	)
	//rmq

	//Echo
	handler := echo.New()
	v1.NewRouter(handler, fileUseCases)
	handler.Logger.Fatal(handler.Start(":8081"))

	//handler := gin.New()
	//v1.NewRouter(handler, l, translationUseCase)
	//httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}
