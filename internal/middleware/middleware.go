package middleware

import (
	"fmt"
	"net/http"
	"os"
	"selarashomeid/pkg/constant"
	"selarashomeid/pkg/util/validator"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	var APP = constant.APP

	e.Use(Context)
	e.Use(
		echoMiddleware.Recover(),
		echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
			AllowOrigins: []string{"https://selarashome.my.id"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "x-user-id", "ngrok-skip-browser-warning", echo.HeaderAuthorization},
			AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodPatch},
		}),
		echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
			Format:           fmt.Sprintf("\n| %s | Host: ${host} | Time: ${time_custom} | Status: ${status} | LatencyHuman: ${latency_human} | UserAgent: ${user_agent} | RemoteIp: ${remote_ip} | Method: ${method} | Uri: ${uri} |\n", APP),
			CustomTimeFormat: "2006/01/02 15:04:05",
			Output:           os.Stdout,
		}),
		echoMiddleware.SecureWithConfig(echoMiddleware.SecureConfig{
			XFrameOptions: "DENY",
		}),
	)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "https://selarashome.my.id")
			c.Response().Header().Set(echo.HeaderAccessControlAllowMethods, "GET,POST,PUT,DELETE,OPTIONS,PATCH")
			c.Response().Header().Set(echo.HeaderAccessControlAllowHeaders, "Origin,Content-Type,Accept,Authorization")
			return next(c)
		}
	})
	e.HTTPErrorHandler = ErrorHandler
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}
}
