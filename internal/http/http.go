package http

import (
	"fmt"
	"net/http"

	_ "selarashomeid/docs"
	"selarashomeid/internal/app/test"
	"selarashomeid/internal/factory"
	"selarashomeid/pkg/constant"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(e *echo.Echo, f *factory.Factory) {
	var (
		APP     = constant.APP
		VERSION = constant.VERSION
	)

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Hello there, welcome to app %s version %s ", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// routes
	api_v1 := e.Group("/selarashomeid/api/v1")
	test.NewHandler(f).Route(api_v1)
}
