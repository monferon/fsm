package app

import (
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/monferon/fsm/loader/config"
	v1 "github.com/monferon/fsm/loader/internal/controller/http/v1"
	"github.com/monferon/fsm/loader/internal/usecase"
	"github.com/monferon/fsm/loader/internal/usecase/repo"
	"github.com/monferon/fsm/loader/internal/usecase/webapi"
	file "github.com/monferon/fsm/loader/pkg/grpc"
	"github.com/monferon/fsm/loader/pkg/logger"
	"github.com/monferon/fsm/loader/pkg/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//TODO вынести в server
type server struct {
	file.SenderServer
}

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "test", "Name to greet")
)

func Run(cfg *config.Config) {
	//logger
	l := logger.New()
	//database
	l.Info("dtatuti", "test-test-test")
	pg, err := postgres.New(cfg.PG.URL)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pg.DB)

	// Use case
	fileUseCases := usecase.New(
		repo.New(pg),
		webapi.New(),
	)
	//grpc
	//server
	//s := grpc.NewServer()
	//file.RegisterSenderServer(s, &server{})

	//client
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		//TODO error
		//log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := file.NewSenderClient(conn)

	//Echo
	handler := echo.New()
	v1.NewRouter(handler, fileUseCases, c)
	handler.Logger.Fatal(handler.Start(":8081"))

	//handler := gin.New()
	//v1.NewRouter(handler, l, translationUseCase)
	//httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
}
