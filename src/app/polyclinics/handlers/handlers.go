package handlers

import (
	"digimer-api/src/app/polyclinics"
	"digimer-api/src/app/polyclinics/handlers/request"
	"digimer-api/src/app/polyclinics/handlers/response"
	"digimer-api/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services polyclinics.Services
}

func NewHandler(srv polyclinics.Services) *Handler {
	return &Handler{srv}
}

// onCreate
func (h *Handler) CreatePolyclinicHandler(ec echo.Context) error {
	var polyclinicRequest request.Request

	if err := ec.Bind(&polyclinicRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	if err := h.services.CreatePolyclinic(polyclinicRequest.MapToDomain()); err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusCreated
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}

// onShowAll
func (h *Handler) ShowAllPolyclinicsHandler(ec echo.Context) error {
	data, err := h.services.GetAllPolyclinics()

	if err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusOK
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), response.MapToBatchResponse(data))
}

// onShowOne
func (h *Handler) ShowPolyclinicByIDHandler(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))
	data, err := h.services.GetPolyclinicByID(id)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			status := http.StatusNotFound
			return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
		}

		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusOK
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), data)
}

// onUpdate
func (h *Handler) AmendPolyclinicByIDHandler(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))
	var polyclinicRequest request.Request

	if err := ec.Bind(&polyclinicRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	if count := h.services.CountPolyclinicByID(id); count == 0 {
		status := http.StatusNotFound
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	if err := h.services.AmendPolyclinicByID(id, polyclinicRequest.MapToDomain()); err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusNoContent
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}

// onDelete
func (h *Handler) RemovePolyclinicByIDHandler(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))

	if count := h.services.CountPolyclinicByID(id); count == 0 {
		status := http.StatusNotFound
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	if err := h.services.RemovePolyclinicByID(id); err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	status := http.StatusNoContent
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}
