package handlers

import (
	"digimer-api/src/app/medical_record_categories"
	"digimer-api/src/app/medical_record_categories/handlers/request"
	"digimer-api/src/app/medical_record_categories/handlers/response"
	"digimer-api/src/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	services medical_record_categories.Services
}

func NewHandler(srv medical_record_categories.Services) *Handler {
	return &Handler{srv}
}

// onCreate
func (h *Handler) CreateMedicalRecordCategoryHandler(ec echo.Context) error {
	var medicineRequest request.Request

	if err := ec.Bind(&medicineRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	if err := h.services.CreateMedicalRecordCategory(medicineRequest.MapToDomain()); err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusCreated
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}

// onShowAll
func (h *Handler) ShowAllMedicalRecordCategoriesHandler(ec echo.Context) error {
	data, err := h.services.GetAllMedicalRecordCategories()

	if err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusOK
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), response.MapToBatchResponse(data))
}

// onShowOne
func (h *Handler) ShowMedicalRecordCategoryByIDHandler(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))
	data, err := h.services.GetMedicalRecordCategoryByID(id)

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
func (h *Handler) AmendMedicalRecordCategoryByIDHandler(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))
	var medicineRequest request.Request

	if err := ec.Bind(&medicineRequest); err != nil {
		status := http.StatusBadRequest
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	if count := h.services.CountMedicalRecordCategoryByID(id); count == 0 {
		status := http.StatusNotFound
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	if err := h.services.AmendMedicalRecordCategoryByID(id, medicineRequest.MapToDomain()); err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), err)
	}

	status := http.StatusNoContent
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}

// onDelete
func (h *Handler) RemoveMedicalRecordCategoryByIDHandler(ec echo.Context) error {
	id, _ := strconv.Atoi(ec.Param("id"))

	if count := h.services.CountMedicalRecordCategoryByID(id); count == 0 {
		status := http.StatusNotFound
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	if err := h.services.RemoveMedicalRecordCategoryByID(id); err != nil {
		status := http.StatusInternalServerError
		return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
	}

	status := http.StatusNoContent
	return utils.CreateEchoResponse(ec, status, http.StatusText(status), nil)
}
