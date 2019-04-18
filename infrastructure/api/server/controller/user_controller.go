package controller

// package interfaces

import (
	"net/http"

	"github.com/anraku/chat/usecase"
	"github.com/labstack/echo"
)

type UserController struct {
	UserInteractor usecase.UserInputBoundary
}

func (controller *UserController) LoginMenu(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func (controller *UserController) Login(c echo.Context) error {
	// execute after save session in echo.Middleware
	return c.Redirect(http.StatusMovedPermanently, "/index")
}

func (controller *UserController) Logout(c echo.Context) error {
	// execute after delete session in echo.Middleware
	return c.Render(http.StatusOK, "logout.html", nil)
}
