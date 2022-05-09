package repositories

import (
	"digimer-api/src/app/medicines"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// CountDataByID implements medicines.Repository
func (repo *repository) CountDataByID(id int) (count int) {
	var medicine Medicine
	return int(repo.DB.Where("ID = ?", id).Find(&medicine).RowsAffected)
}

// DeleteByID implements medicines.Repository
func (repo *repository) DeleteByID(id int) (err error) {
	err = repo.DB.Where("ID = ?", id).Delete(new(Medicine)).Error
	return err
}

// InsertData implements medicines.Repository
func (repo *repository) InsertData(domain medicines.Domain) (err error) {
	record := mapToRecord(domain)
	err = repo.DB.Create(&record).Error
	return
}

// SelectAllData implements medicines.Repository
func (repo *repository) SelectAllData() (data []medicines.Domain, err error) {
	var medicines []Medicine
	if err = repo.DB.Find(&medicines).Error; err != nil {
		return nil, err
	}

	return mapToDomainBatch(medicines), nil
}

// SelectDataByID implements medicines.Repository
func (repo *repository) SelectDataByID(id int) (selected medicines.Domain, err error) {
	var medicine Medicine
	if err = repo.DB.First(&medicine, id).Error; err != nil {
		return medicines.Domain{}, err
	}

	return mapToDomain(medicine), nil
}

// UpdateByID implements medicines.Repository
func (repo *repository) UpdateByID(id int, domain medicines.Domain) (err error) {
	var medicine Medicine
	record := mapToRecord(domain)
	err = repo.DB.Model(&medicine).Where("ID = ?", id).Updates(&record).Error
	return
}

func NewMySQLRepository(DB *gorm.DB) medicines.Repository {
	return &repository{DB}
}
