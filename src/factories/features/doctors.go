package features

import (
	"digimer-api/src/app/doctors/handlers"
	"digimer-api/src/app/doctors/repositories"
	"digimer-api/src/app/doctors/services"

	"gorm.io/gorm"
)

func DoctorFactory(conn *gorm.DB) handlers.Handler {
	repo := repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *handlers.NewHandler(serv)
}
