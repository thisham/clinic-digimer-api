package factories

import (
	doctorsHandler "digimer-api/src/app/doctors/handlers"
	medicalRecordCategoryHandler "digimer-api/src/app/medical_record_categories/handlers"
	medicinesHandler "digimer-api/src/app/medicines/handlers"
	patientsHandler "digimer-api/src/app/patients/handlers"
	polyclinicsHandler "digimer-api/src/app/polyclinics/handlers"
	"digimer-api/src/database"
	"digimer-api/src/factories/features"
)

type apiHandler struct {
	Polyclinic            polyclinicsHandler.Handler
	Medicine              medicinesHandler.Handler
	MedicalRecordCategory medicalRecordCategoryHandler.Handler
	Patient               patientsHandler.Handler
	Doctor                doctorsHandler.Handler
}

func ApiInit() apiHandler {
	conn := new(database.DBConf).InitDB()

	return apiHandler{
		Polyclinic:            features.PolyclinicFactory(conn.DB),
		Medicine:              features.MedicineFactory(conn.DB),
		MedicalRecordCategory: features.MedicalRecordCategoryFactory(conn.DB),
		Patient:               features.PatientFactory(conn.DB),
		Doctor:                features.DoctorFactory(conn.DB),
	}
}
