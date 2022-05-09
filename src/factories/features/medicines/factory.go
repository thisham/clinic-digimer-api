package medicines

import (
	"digimer-api/src/app/medicines/handlers"
	"digimer-api/src/app/medicines/repositories"
	"digimer-api/src/app/medicines/services"

	"gorm.io/gorm"
)

func MedicineFactory(conn *gorm.DB) handlers.Handler {
	repo := repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *handlers.NewHandler(serv)
}
