package database

import (
	doctors "digimer-api/src/app/doctors/repositories"
	medicalRecordCategories "digimer-api/src/app/medical_record_categories/repositories"
	medicalRecords "digimer-api/src/app/medical_records/repositories"
	medicines "digimer-api/src/app/medicines/repositories"
	patients "digimer-api/src/app/patients/repositories"
	polyclinics "digimer-api/src/app/polyclinics/repositories"
	"digimer-api/src/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConf struct{ *gorm.DB }

func (DB *DBConf) InitDB() *DBConf {
	config, _ := configs.LoadServerConfig(".")
	dsn := config.ConnectionString

	conn, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return &DBConf{conn}
}

func (DB *DBConf) Migrate() {
	DB.AutoMigrate(
		&polyclinics.Polyclinic{},
		&medicines.Medicine{},
		&medicalRecordCategories.MedicalRecordCategory{},
		&patients.Patient{},
		&doctors.Doctor{},
		&medicalRecords.MedicalRecord{},
	)
}

func (DB *DBConf) Demigrate() {
	DB.Migrator().DropTable(
		&polyclinics.Polyclinic{},
		&medicines.Medicine{},
		&medicalRecordCategories.MedicalRecordCategory{},
		&patients.Patient{},
		&doctors.Doctor{},
		&medicalRecords.MedicalRecord{},
	)
}
