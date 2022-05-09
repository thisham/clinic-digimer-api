package factories

import (
	medicinesHandler "digimer-api/src/app/medicines/handlers"
	polyclinicsHandler "digimer-api/src/app/polyclinics/handlers"
	"digimer-api/src/database"
	"digimer-api/src/factories/features"
)

type apiHandler struct {
	Polyclinic polyclinicsHandler.Handler
	Medicine   medicinesHandler.Handler
}

func ApiInit() apiHandler {
	conn := new(database.DBConf).InitDB()

	return apiHandler{
		Polyclinic: features.PolyclinicFactory(conn.DB),
		Medicine:   features.MedicineFactory(conn.DB),
	}
}
