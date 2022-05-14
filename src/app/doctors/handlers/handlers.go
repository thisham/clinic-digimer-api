package handlers

import (
	"digimer-api/src/app/doctors"
	"digimer-api/src/app/doctors/handlers/request"
	"digimer-api/src/app/doctors/handlers/response"
	"digimer-api/src/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services doctors.Services
}

func NewHandler(srv doctors.Services) *Handler {
	return &Handler{srv}
}

// onUpdate
func (h *Handler) AmendDoctorByIDHandler(ec echo.Context) error {
	var doctorRequest request.Request

	if err := ec.Bind(&doctorRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	if err := h.services.AmendDoctorByID(ec.Param("id"), doctorRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			status := http.StatusNotFound
			return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
		}

		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	status := http.StatusNoContent
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}

// onUpdatePassword
func (h *Handler) AmendPasswordByDoctorIDHandler(ec echo.Context) error {
	var doctorRequest request.Request

	if err := ec.Bind(&doctorRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	if err := h.services.AmendDoctorByID(ec.Param("id"), doctorRequest.MapToDomain()); err != nil {
		if strings.Contains(err.Error(), "not found") {
			status := http.StatusNotFound
			return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
		}

		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	status := http.StatusNoContent
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}

// onAttemptLogin
func (h *Handler) AttemptDoctorLoginHandler(ec echo.Context) error {
	var doctorRequest request.LoginRequest

	if err := ec.Bind(&doctorRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	token, err := h.services.AttemptDoctorLogin(doctorRequest.Email, doctorRequest.Password)
	if err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}
	status := http.StatusInternalServerError
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), map[string]string{
		"token": token,
	})
}

// onCreate
func (h *Handler) CreateDoctorHandler(ec echo.Context) error {
	var doctorRequest request.Request

	if err := ec.Bind(&doctorRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	id, err := h.services.CreateDoctor(doctorRequest.MapToDomain())
	if err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}
	status := http.StatusInternalServerError
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), map[string]string{
		"id": id,
	})
}

// onShowAll
func (h *Handler) ShowAllDoctorsHandler(ec echo.Context) error {
	data, err := h.services.GetAllDoctors()

	if err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	status := http.StatusOK
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), response.MapToBatchResponse(data))
}

// onShowOne
func (h *Handler) ShowDoctorByIDHandler(ec echo.Context) error {
	data, err := h.services.GetDoctorByID(ec.Param("id"))

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			status := http.StatusNotFound
			return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
		}

		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	status := http.StatusOK
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), response.MapToResponse(data))
}

// onDelete
func (h *Handler) RemoveDoctorByIDHandler(ec echo.Context) error {
	if err := h.services.RemoveDoctorByID(ec.Param("id")); err != nil {
		if strings.Contains(err.Error(), "not found") {
			status := http.StatusNotFound
			return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
		}

		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	status := http.StatusNoContent
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}
