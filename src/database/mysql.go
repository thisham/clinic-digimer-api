package database

import (
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
	DB.AutoMigrate(&polyclinics.Polyclinic{})
}

func (DB *DBConf) Demigrate() {
	DB.Migrator().DropTable(&polyclinics.Polyclinic{})
}
