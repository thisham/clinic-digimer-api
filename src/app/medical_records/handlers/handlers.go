package handlers

import (
	"digimer-api/src/app/medical_records"
	"digimer-api/src/app/medical_records/handlers/request"
	"digimer-api/src/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	services medical_records.Services
}

func NewHandler(srv medical_records.Services) *Handler {
	return &Handler{srv}
}

func (h *Handler) AmendMedicalRecordByIDHandler(ec echo.Context) error {
	var medicalRecordRequest request.Request

	if err := ec.Bind(&medicalRecordRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err.Error())
	}

	mrRequest := request.MapToDomain(medicalRecordRequest)
	cookie, _ := ec.Cookie("token")
	jwtData, _ := utils.ExtractClaims(cookie.Value)
	mrRequest.Doctor.ID = uuid.MustParse(jwtData.Id)

	if err := h.services.AmendMedicalRecordByID(ec.Param("id"), mrRequest); err != nil {
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

func (h *Handler) RemoveMedicalRecordByIDHandler(ec echo.Context) error {
	if err := h.services.DeleteMedicalRecordByID(ec.Param("id")); err != nil {
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

func (h *Handler) MapJWTHandler(ec echo.Context) error {
	token, _ := ec.Cookie("token")
	jwt, _ := utils.ExtractClaims(token.Value)
	status := http.StatusOK
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), jwt)
}
