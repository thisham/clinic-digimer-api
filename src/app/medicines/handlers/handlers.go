package handlers

import (
	"digimer-api/src/app/medicines"
	"digimer-api/src/app/medicines/handlers/request"
	"digimer-api/src/app/medicines/handlers/response"
	"digimer-api/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services medicines.Services
}

func NewHandler(srv medicines.Services) *Handler {
	return &Handler{srv}
}

// onCreate
func (h *Handler) CreateMedicineHandler(ec echo.Context) error {
	var medicineRequest request.Request

	if err := ec.Bind(&medicineRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	if err := h.services.CreateMedicine(medicineRequest.MapToDomain()); err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusCreated
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}

// onShowAll
func (h *Handler) ShowAllMedicinesHandler(ec echo.Context) error {
	data, err := h.services.GetAllMedicines()

	if err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusOK
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), response.MapToBatchResponse(data))
}

// onShowOne
func (h *Handler) ShowMedicineByIDHandler(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))
	data, err := h.services.GetMedicineByID(id)

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
func (h *Handler) AmendMedicineByIDHandler(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))
	var medicineRequest request.Request

	if err := ec.Bind(&medicineRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	if count := h.services.CountMedicineByID(id); count == 0 {
		status := http.StatusNotFound
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	if err := h.services.AmendMedicineByID(id, medicineRequest.MapToDomain()); err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusNoContent
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}

// onDelete
func (h *Handler) RemoveMedicineByIDHandler(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))

	if count := h.services.CountMedicineByID(id); count == 0 {
		status := http.StatusNotFound
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	if err := h.services.RemoveMedicineByID(id); err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	status := http.StatusNoContent
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}
