package repositories

import (
	"digimer-api/src/app/medical_records"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// CountDataByID implements medical_records.Repositories
func (repo *repository) CountDataByID(id string) (count int) {
	return int(repo.DB.Where("ID = ?", id).Find(new(MedicalRecord)).RowsAffected)
}

// DeleteByID implements medical_records.Repositories
func (repo *repository) DeleteByID(id string) (err error) {
	return repo.DB.Where("ID = ?", id).Delete(new(MedicalRecord)).Error
}

// InsertData implements medical_records.Repositories
func (repo *repository) InsertData(domain medical_records.Domain) (id string, err error) {
	record := mapToNewRecord(domain)
	return record.ID.String(), repo.DB.Create(&record).Error
}

// SelectAllData implements medical_records.Repositories
func (repo *repository) SelectAllData() (data []medical_records.Domain, err error) {
	var records []MedicalRecord

	if err = repo.DB.Find(&records).Error; err != nil {
		return nil, err
	}
	return mapToDomainBatch(records), nil
}

// SelectDataByID implements medical_records.Repositories
func (repo *repository) SelectDataByID(id string) (selected medical_records.Domain, err error) {
	var record MedicalRecord

	if err = repo.DB.Where("ID = ?", id).Find(&record).Error; err != nil {
		return medical_records.Domain{}, err
	}
	return mapToDomain(record), nil
}

// UpdateByID implements medical_records.Repositories
func (repo *repository) UpdateByID(id string, domain medical_records.Domain) (err error) {
	record := mapToExistingRecord(domain)
	return repo.DB.Model(new(MedicalRecord)).Where("ID = ?", id).Updates(&record).Error
}

func NewMySQLRepository(DB *gorm.DB) medical_records.Repositories {
	return &repository{DB}
}
