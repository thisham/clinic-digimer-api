package repositories

import (
	"digimer-api/src/app/polyclinics"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// DeleteByID implements polyclinics.Repository
func (repo *repository) DeleteByID(id int) (err error) {
	// var polyclinic Polyclinic
	err = repo.DB.Where("ID = ?", id).Delete(new(Polyclinic)).Error
	return err
}

// InsertData implements polyclinics.Repository
func (repo *repository) InsertData(domain polyclinics.Domain) (err error) {
	record := mapToRecord(domain)
	err = repo.DB.Create(&record).Error
	return
}

// SelectAllData implements polyclinics.Repository
func (repo *repository) SelectAllData() (data []polyclinics.Domain, err error) {
	var polyclinics []Polyclinic
	if err = repo.DB.Find(&polyclinics).Error; err != nil {
		return nil, err
	}

	return mapToDomainBatch(polyclinics), nil
}

// SelectDataByID implements polyclinics.Repository
func (repo *repository) SelectDataByID(id int) (selected polyclinics.Domain, err error) {
	var polyclinic Polyclinic
	if err = repo.DB.First(&polyclinic, id).Error; err != nil {
		return polyclinics.Domain{}, err
	}

	return mapToDomain(polyclinic), nil
}

// UpdateByID implements polyclinics.Repository
func (repo *repository) UpdateByID(id int, domain polyclinics.Domain) (err error) {
	var polyclinic Polyclinic
	record := mapToRecord(domain)
	err = repo.DB.Model(&polyclinic).Where("ID = ?", id).Updates(&record).Error
	return
}

func NewMySQLRepository(DB *gorm.DB) polyclinics.Repository {
	return &repository{DB}
}
