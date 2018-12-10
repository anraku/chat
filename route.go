// package infrastracture
package main

import (
	"html/template"
	"io"
	"net/http"

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

func NewRouter(ui *UserInteractor, ri *RoomInteractor, mi *MessageInteractor) *echo.Echo {
	// create user controller
	userController := UserController{
		UserInteractor: ui,
	}

	// create room controller
	roomController := RoomController{
		RoomInteractor:    ri,
		MessageInteractor: mi,
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
	// UserController
	e.GET("/login", func(c echo.Context) error { return userController.LoginMenu(c) })
	e.POST("/login", func(c echo.Context) error { return userController.Login(c) }, saveSession)
	e.GET("/logout", func(c echo.Context) error { return userController.Logout(c) }, deleteSession)

	// RoomController
	e.GET("/", func(c echo.Context) error { return roomController.Index(c) }, checkLogin)
	e.GET("/index", func(c echo.Context) error { return roomController.Index(c) }, checkLogin)
	e.GET("/chat/:id", func(c echo.Context) error { return roomController.EnterRoom(c) })
	e.GET("/room/:id", func(c echo.Context) error { return roomController.Chat(c) }, checkLogin)
	e.GET("/room/new", func(c echo.Context) error { return roomController.NewRoom(c) }, checkLogin)
	e.POST("/room/create", func(c echo.Context) error { return roomController.CreateRoom(c) }, checkLogin)

	return e
}

func checkLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get username from session
		sess, err := session.Get("session", c)
		if err != nil {
			return err
		}
		username := sess.Values["username"]
		if username == nil || username == "" {
			return c.Redirect(http.StatusMovedPermanently, "/login")
		}
		c.Set("username", username)
		return next(c)
	}
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
