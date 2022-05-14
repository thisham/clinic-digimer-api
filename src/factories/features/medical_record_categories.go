package features

import (
	"digimer-api/src/app/medical_record_categories/handlers"
	"digimer-api/src/app/medical_record_categories/repositories"
	"digimer-api/src/app/medical_record_categories/services"

	"gorm.io/gorm"
)

func MedicalRecordCategoryFactory(conn *gorm.DB) handlers.Handler {
	repo := repositories.NewMySQLRepository(conn)
	serv := services.NewService(repo)
	return *handlers.NewHandler(serv)
}
