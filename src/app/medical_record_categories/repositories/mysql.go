package repositories

import (
	"digimer-api/src/app/medical_record_categories"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// CountDataByID implements medical_record_categories.Repository
func (repo *repository) CountDataByID(id int) (count int) {
	var medical_record_category MedicalRecordCategory
	return int(repo.DB.Where("ID = ?", id).Find(&medical_record_category).RowsAffected)
}

// DeleteByID implements medical_record_categories.Repository
func (repo *repository) DeleteByID(id int) (err error) {
	err = repo.DB.Where("ID = ?", id).Delete(new(MedicalRecordCategory)).Error
	return err
}

// InsertData implements medical_record_categories.Repository
func (repo *repository) InsertData(domain medical_record_categories.Domain) (err error) {
	record := mapToRecord(domain)
	err = repo.DB.Create(&record).Error
	return
}

// SelectAllData implements medical_record_categories.Repository
func (repo *repository) SelectAllData() (data []medical_record_categories.Domain, err error) {
	var medical_record_categories []MedicalRecordCategory
	if err = repo.DB.Find(&medical_record_categories).Error; err != nil {
		return nil, err
	}

	return mapToDomainBatch(medical_record_categories), nil
}

// SelectDataByID implements medical_record_categories.Repository
func (repo *repository) SelectDataByID(id int) (selected medical_record_categories.Domain, err error) {
	var medical_record_category MedicalRecordCategory
	if err = repo.DB.First(&medical_record_category, id).Error; err != nil {
		return medical_record_categories.Domain{}, err
	}

	return mapToDomain(medical_record_category), nil
}

// UpdateByID implements medical_record_categories.Repository
func (repo *repository) UpdateByID(id int, domain medical_record_categories.Domain) (err error) {
	var medical_record_category MedicalRecordCategory
	record := mapToRecord(domain)
	err = repo.DB.Model(&medical_record_category).Where("ID = ?", id).Updates(&record).Error
	return
}

func NewMySQLRepository(DB *gorm.DB) medical_record_categories.Repository {
	return &repository{DB}
}
