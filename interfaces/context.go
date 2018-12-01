package interfaces

import (
	"net/http"

	"github.com/labstack/echo"
)

type Context interface {
	Get(key string) interface{}
	JSON(code int, i interface{}) error
	Param(string) string
	FormValue(name string) string
	Render(code int, name string, data interface{}) error
	Request() *http.Request
	Redirect(code int, url string) error
	Response() *echo.Response
}
