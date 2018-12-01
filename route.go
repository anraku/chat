// package infrastracture
package main

import (
	"html/template"
	"io"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

// Template used for creating HTML
type Template struct {
	templates *template.Template
}

// Render create HTML with template file
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewRouter(ui *UserInteractor) *echo.Echo {
	userController := UserController{
		Interactor: ui,
	}
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	// setting static filgs[es
	e.Static("/vendor", "vendor")
	e.Static("/css", "css")
	e.Static("/avatars", "avatars")

	// setting tmeplates
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	// Routing
	e.GET("/login", func(c echo.Context) error { return userController.LoginMenu(c) })
	e.POST("/login", func(c echo.Context) error { return userController.Login(c) }, saveSession)
	e.GET("/logout", func(c echo.Context) error { return userController.Logout(c) }, deleteSession)
	e.GET("/", Index, CheckLogin)
	e.GET("/index", Index, CheckLogin)
	e.GET("/chat/:id", EnterRoom)
	e.GET("/room/:id", Chat)
	e.GET("/room/new", NewRoom, CheckLogin)
	e.POST("/room/create", CreateRoom, CheckLogin)

	return e
}

func saveSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.FormValue("name")
		if userName == "" {
			userName = "名無しさん"
		}
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		sess.Values["username"] = userName
		sess.Save(c.Request(), c.Response())
		return next(c)
	}
}

func deleteSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		delete(sess.Values, "username")
		sess.Save(c.Request(), c.Response())
		return next(c)
	}
}
