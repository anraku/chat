package main

// package interfaces

import (
	"net/http"

	"github.com/anraku/chat/interfaces"
)

type UserController struct {
	Interactor *UserInteractor
}

func (controller *UserController) LoginMenu(c interfaces.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func (controller *UserController) Login(c interfaces.Context) error {
	// save session in echo.Middleware
	return c.Redirect(http.StatusMovedPermanently, "/index")
}

func (controller *UserController) Logout(c interfaces.Context) error {
	// delete session in echo.Middleware
	return c.Render(http.StatusOK, "logout.html", nil)
}
