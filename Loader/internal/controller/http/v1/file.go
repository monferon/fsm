package v1

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/monferon/fsm/loader/internal/usecase"
	file "github.com/monferon/fsm/loader/pkg/grpc"
	"io"
	"net/http"
	"os"
)

type fileRoutes struct {
	t  *usecase.FileInfoUseCase
	gc file.SenderClient
	//l logger.Interface
}

func newFileRoutes(e *echo.Group, f *usecase.FileInfoUseCase, gc file.SenderClient) {
	r := &fileRoutes{f, gc}

	ee := e.Group("/v1")
	{
		//newFileRoutes(ee, f)
		ee.POST("/upload", r.upload)
	}
}

func (f *fileRoutes) upload(c echo.Context) error {
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully.</p>", file.Filename))
	//translations, err := r.t.History(c.Request.Context())
	//if err != nil {
	//	r.l.Error(err, "http - v1 - history")
	//	errorResponse(c, http.StatusInternalServerError, "database problems")
	//
	//	return
	//}
	//
	//c.JSON(http.StatusOK, historyResponse{translations})
	return nil
}
