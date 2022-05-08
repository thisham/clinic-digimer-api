package factories

import (
	polyclinicsHandler "digimer-api/src/app/polyclinics/handlers"
	"digimer-api/src/database"
	"digimer-api/src/factories/features/polyclinics"
)

type apiHandler struct {
	Polyclinic polyclinicsHandler.Handler
}

func ApiInit() apiHandler {
	conn := new(database.DBConf).InitDB()

	return apiHandler{
		Polyclinic: polyclinics.PolyclinicFactory(conn.DB),
	}
}
