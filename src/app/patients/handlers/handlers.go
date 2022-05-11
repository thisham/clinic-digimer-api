package handlers

import (
	"digimer-api/src/app/patients"
	"digimer-api/src/app/patients/handlers/request"
	"digimer-api/src/app/patients/handlers/response"
	"digimer-api/src/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services patients.Services
}

func NewHandler(srv patients.Services) *Handler {
	return &Handler{srv}
}

// onCreate
func (h *Handler) CreatePatientHandler(ec echo.Context) error {
	var patientRequest request.Request

	if err := ec.Bind(&patientRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	uid, err := h.services.CreatePatient(patientRequest.MapToDomain())
	if err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	status := http.StatusCreated
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), map[string]string{
		"uid": uid,
	})
}

// onShowAll
func (h *Handler) ShowAllPatientsHandler(ec echo.Context) error {
	data, err := h.services.GetAllPatients()

	if err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	status := http.StatusOK
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), response.MapToBatchResponse(data))
}

// onShowOne
func (h *Handler) ShowPatientByIDHandler(ec echo.Context) error {
	data, err := h.services.GetPatientByID(ec.Param("id"))

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

// onShowOneByMRBookNumber
func (h *Handler) ShowPatientByMRBookNumberHandler(ec echo.Context) error {
	data, err := h.services.GetPatientByMRBookNumber(ec.Param("mrid"))

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

// onUpdate
func (h *Handler) AmendPatientByIDHandler(ec echo.Context) error {
	var patientRequest request.Request

	if err := ec.Bind(&patientRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	if err := h.services.AmendPatientByID(ec.Param("id"), patientRequest.MapToDomain()); err != nil {
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

// onDelete
func (h *Handler) RemovePatientByIDHandler(ec echo.Context) error {
	if err := h.services.RemovePatientByID(ec.Param("id")); err != nil {
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
