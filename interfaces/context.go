package interfaces

import (
	"net/http"
)

type Context interface {
	JSON(code int, i interface{}) error
	Param(string) string
	Render(code int, name string, data interface{}) error
	Request() *http.Request
	Redirect(code int, url string) error
}
