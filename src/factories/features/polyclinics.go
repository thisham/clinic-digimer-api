package features

import (
	"digimer-api/src/app/polyclinics/handlers"
	"digimer-api/src/app/polyclinics/repositories"
	"digimer-api/src/app/polyclinics/services"

	"gorm.io/gorm"
)

func PolyclinicFactory(conn *gorm.DB) handlers.Handler {
	repo := repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *handlers.NewHandler(serv)
}
