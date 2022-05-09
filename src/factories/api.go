package factories

import (
	medicalRecordCategoryHandler "digimer-api/src/app/medical_record_categories/handlers"
	medicinesHandler "digimer-api/src/app/medicines/handlers"
	polyclinicsHandler "digimer-api/src/app/polyclinics/handlers"
	"digimer-api/src/database"
	"digimer-api/src/factories/features"
)

type apiHandler struct {
	Polyclinic            polyclinicsHandler.Handler
	Medicine              medicinesHandler.Handler
	MedicalRecordCategory medicalRecordCategoryHandler.Handler
}

func ApiInit() apiHandler {
	conn := new(database.DBConf).InitDB()

	return apiHandler{
		Polyclinic:            features.PolyclinicFactory(conn.DB),
		Medicine:              features.MedicineFactory(conn.DB),
		MedicalRecordCategory: features.MedicalRecordCategoryFactory(conn.DB),
	}
}
