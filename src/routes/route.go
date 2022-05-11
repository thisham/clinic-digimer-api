package routes

import (
	"digimer-api/src/factories"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()
	api := route.Group("/api")

	webFactory := factories.WebInit()
	apiFactory := factories.ApiInit()

	// main
	route.GET("/", webFactory.InitialPageHandler.InitialPage)

	// api
	// polyclinics
	api.GET("/polyclinics", apiFactory.Polyclinic.ShowAllPolyclinicsHandler)
	api.GET("/polyclinics/:id", apiFactory.Polyclinic.ShowPolyclinicByIDHandler)
	api.POST("/polyclinics", apiFactory.Polyclinic.CreatePolyclinicHandler)
	api.PUT("/polyclinics/:id", apiFactory.Polyclinic.AmendPolyclinicByIDHandler)
	api.DELETE("/polyclinics/:id", apiFactory.Polyclinic.RemovePolyclinicByIDHandler)

	// medicines
	api.GET("/medicines", apiFactory.Medicine.ShowAllMedicinesHandler)
	api.GET("/medicines/:id", apiFactory.Medicine.ShowMedicineByIDHandler)
	api.POST("/medicines", apiFactory.Medicine.CreateMedicineHandler)
	api.PUT("/medicines/:id", apiFactory.Medicine.AmendMedicineByIDHandler)
	api.DELETE("/medicines/:id", apiFactory.Medicine.RemoveMedicineByIDHandler)

	// medicines
	api.GET("/medical-record-categories", apiFactory.MedicalRecordCategory.ShowAllMedicalRecordCategoriesHandler)
	api.GET("/medical-record-categories/:id", apiFactory.MedicalRecordCategory.ShowMedicalRecordCategoryByIDHandler)
	api.POST("/medical-record-categories", apiFactory.MedicalRecordCategory.CreateMedicalRecordCategoryHandler)
	api.PUT("/medical-record-categories/:id", apiFactory.MedicalRecordCategory.AmendMedicalRecordCategoryByIDHandler)
	api.DELETE("/medical-record-categories/:id", apiFactory.MedicalRecordCategory.RemoveMedicalRecordCategoryByIDHandler)

	// patients
	api.GET("/patients", apiFactory.Patient.ShowAllPatientsHandler)
	api.GET("/patients/id/:id", apiFactory.Patient.ShowPatientByIDHandler)
	api.GET("/patients/mr/:mrid", apiFactory.Patient.ShowPatientByMRBookNumberHandler)
	api.POST("/patients", apiFactory.Patient.CreatePatientHandler)
	api.PUT("/patients/id/:id", apiFactory.Patient.AmendPatientByIDHandler)
	api.DELETE("/patients/id/:id", apiFactory.Patient.RemovePatientByIDHandler)

	return route
}
