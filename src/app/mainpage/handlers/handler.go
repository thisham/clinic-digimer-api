package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct{}

func NewMainPageHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitialPage(ec echo.Context) error {
	return ec.HTML(http.StatusOK, "Server runs perfectly!")
}
