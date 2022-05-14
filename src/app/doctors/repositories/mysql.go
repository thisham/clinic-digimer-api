package repositories

import (
	"digimer-api/src/app/doctors"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// CountDataByID implements doctors.Repositories
func (repo *repository) CountDataByID(id string) (count int) {
	return int(repo.DB.Where("ID = ?", id).Find(new(Doctor)).RowsAffected)
}

// DeleteByID implements doctors.Repositories
func (repo *repository) DeleteByID(id string) (err error) {
	return repo.DB.Where("ID = ?", id).Delete(new(Doctor)).Error
}

// InsertData implements doctors.Repositories
func (repo *repository) InsertData(domain doctors.Domain) (id string, err error) {
	record := mapToNewRecord(domain)
	return record.ID.String(), repo.DB.Create(&record).Error
}

// SelectAllData implements doctors.Repositories
func (repo *repository) SelectAllData() (data []doctors.Domain, err error) {
	var records []Doctor

	if err = repo.DB.Find(&records).Error; err != nil {
		return nil, err
	}
	return mapToDomainBatch(records), nil
}

// SelectDataByEmail implements doctors.Repositories
func (repo *repository) SelectDataByEmail(email string) (selected doctors.Domain, err error) {
	var record Doctor

	if err = repo.DB.Where("email = ?", email).First(&record).Error; err != nil {
		return doctors.Domain{}, err
	}
	return mapToDomain(record), nil
}

// SelectDataByID implements doctors.Repositories
func (repo *repository) SelectDataByID(id string) (selected doctors.Domain, err error) {
	var record Doctor

	if err = repo.DB.Where("ID = ?", id).First(&record).Error; err != nil {
		return doctors.Domain{}, err
	}
	return mapToDomain(record), nil
}

// UpdateByID implements doctors.Repositories
func (repo *repository) UpdateByID(id string, domain doctors.Domain) (err error) {
	record := mapToExistingRecord(domain)
	return repo.DB.Model(new(Doctor)).Where("ID = ?", id).Updates(&record).Error
}

func NewMySQLRepository(DB *gorm.DB) doctors.Repositories {
	return &repository{DB}
}
