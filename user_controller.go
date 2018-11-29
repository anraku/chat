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
