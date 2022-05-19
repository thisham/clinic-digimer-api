package routes

import (
	"digimer-api/src/factories"
	"digimer-api/src/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()
	api := route.Group("/api")

	webFactory := factories.WebInit()
	apiFactory := factories.ApiInit()

	authApi := api.Group("")
	authApi.Use(middlewares.VerifyAuthentication())

	// main
	route.GET("/", webFactory.InitialPageHandler.InitialPage)

	// api
	// polyclinics
	authApi.GET("/polyclinics", apiFactory.Polyclinic.ShowAllPolyclinicsHandler)
	authApi.GET("/polyclinics/:id", apiFactory.Polyclinic.ShowPolyclinicByIDHandler)
	authApi.POST("/polyclinics", apiFactory.Polyclinic.CreatePolyclinicHandler)
	authApi.PUT("/polyclinics/:id", apiFactory.Polyclinic.AmendPolyclinicByIDHandler)
	authApi.DELETE("/polyclinics/:id", apiFactory.Polyclinic.RemovePolyclinicByIDHandler)

	// medicines
	authApi.GET("/medicines", apiFactory.Medicine.ShowAllMedicinesHandler)
	authApi.GET("/medicines/:id", apiFactory.Medicine.ShowMedicineByIDHandler)
	authApi.POST("/medicines", apiFactory.Medicine.CreateMedicineHandler)
	authApi.PUT("/medicines/:id", apiFactory.Medicine.AmendMedicineByIDHandler)
	authApi.DELETE("/medicines/:id", apiFactory.Medicine.RemoveMedicineByIDHandler)

	// medicines
	authApi.GET("/medical-record-categories", apiFactory.MedicalRecordCategory.ShowAllMedicalRecordCategoriesHandler)
	authApi.GET("/medical-record-categories/:id", apiFactory.MedicalRecordCategory.ShowMedicalRecordCategoryByIDHandler)
	authApi.POST("/medical-record-categories", apiFactory.MedicalRecordCategory.CreateMedicalRecordCategoryHandler)
	authApi.PUT("/medical-record-categories/:id", apiFactory.MedicalRecordCategory.AmendMedicalRecordCategoryByIDHandler)
	authApi.DELETE("/medical-record-categories/:id", apiFactory.MedicalRecordCategory.RemoveMedicalRecordCategoryByIDHandler)

	// patients
	authApi.GET("/patients", apiFactory.Patient.ShowAllPatientsHandler)
	authApi.GET("/patients/id/:id", apiFactory.Patient.ShowPatientByIDHandler)
	authApi.GET("/patients/mr/:mrid", apiFactory.Patient.ShowPatientByMRBookNumberHandler)
	authApi.POST("/patients", apiFactory.Patient.CreatePatientHandler)
	authApi.PUT("/patients/id/:id", apiFactory.Patient.AmendPatientByIDHandler)
	authApi.DELETE("/patients/id/:id", apiFactory.Patient.RemovePatientByIDHandler)

	// doctors
	authApi.GET("/doctors", apiFactory.Doctor.ShowAllDoctorsHandler)
	authApi.GET("/doctors/:id", apiFactory.Doctor.ShowDoctorByIDHandler)
	authApi.POST("/doctors", apiFactory.Doctor.CreateDoctorHandler)
	authApi.PUT("/doctors/profile/:id", apiFactory.Doctor.AmendDoctorByIDHandler)
	authApi.PUT("/doctors/password/:id", apiFactory.Doctor.AmendPasswordByDoctorIDHandler)
	authApi.DELETE("/doctors/:id", apiFactory.Doctor.RemoveDoctorByIDHandler)

	// auths
	api.POST("/doctors/auth/login", apiFactory.Doctor.AttemptDoctorLoginHandler)

	return route
}
