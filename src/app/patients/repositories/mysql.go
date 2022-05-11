package repositories

import (
	"digimer-api/src/app/patients"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// CountDataByID implements patients.Repositories
func (repo *repository) CountDataByID(id string) (count int) {
	return int(repo.DB.Where("ID = ?", id).Find(new(Patient)).RowsAffected)
}

// DeleteByID implements patients.Repositories
func (repo *repository) DeleteByID(id string) (err error) {
	return repo.DB.Where("ID = ?", id).Find(new(Patient)).Error
}

// InsertData implements patients.Repositories
func (repo *repository) InsertData(domain patients.Domain) (err error) {
	record := mapToNewRecord(domain)
	return repo.DB.Create(&record).Error
}

// LookupLatestMRBookNumber implements patients.Repositories
func (repo *repository) LookupLatestMRBookNumber() (mrBookNumber string) {
	var patient Patient
	repo.DB.Order("medical_record_book_number desc").First(&patient)
	return patient.MedicalRecordBookNumber
}

// SelectAllData implements patients.Repositories
func (repo *repository) SelectAllData() (data []patients.Domain, err error) {
	var records []Patient

	if err := repo.DB.Find(&records).Error; err != nil {
		return nil, err
	}
	return mapToDomainBatch(records), nil
}

// SelectDataByID implements patients.Repositories
func (repo *repository) SelectDataByID(id string) (selected patients.Domain, err error) {
	var record Patient

	if err := repo.DB.Where("ID = ?", id).First(&record).Error; err != nil {
		return patients.Domain{}, err
	}
	return mapToDomain(record), nil
}

// SelectDataByMRBookNumber implements patients.Repositories
func (repo *repository) SelectDataByMRBookNumber(mrBookNumber string) (selected patients.Domain, err error) {
	var record Patient

	if err := repo.DB.Where("medical_record_book_number = ?", mrBookNumber).First(&record).Error; err != nil {
		return patients.Domain{}, err
	}
	return mapToDomain(record), nil
}

// UpdateByID implements patients.Repositories
func (repo *repository) UpdateByID(id string, domain patients.Domain) (err error) {
	record := mapToExistingRecord(domain)
	return repo.DB.Model(new(Patient)).Where("ID = ?", id).Updates(&record).Error
}

func NewMySQLRepository(DB *gorm.DB) patients.Repositories {
	return &repository{DB}
}
