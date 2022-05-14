package features

import (
	"digimer-api/src/app/patients/handlers"
	"digimer-api/src/app/patients/repositories"
	"digimer-api/src/app/patients/services"

	"gorm.io/gorm"
)

func PatientFactory(conn *gorm.DB) handlers.Handler {
	repo := repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *handlers.NewHandler(serv)
}
