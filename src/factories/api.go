package factories

import (
	medicinesHandler "digimer-api/src/app/medicines/handlers"
	polyclinicsHandler "digimer-api/src/app/polyclinics/handlers"
	"digimer-api/src/database"
	"digimer-api/src/factories/features/medicines"
	"digimer-api/src/factories/features/polyclinics"
)

type apiHandler struct {
	Polyclinic polyclinicsHandler.Handler
	Medicine   medicinesHandler.Handler
}

func ApiInit() apiHandler {
	conn := new(database.DBConf).InitDB()

	return apiHandler{
		Polyclinic: polyclinics.PolyclinicFactory(conn.DB),
		Medicine:   medicines.MedicineFactory(conn.DB),
	}
}
