package services

import "digimer-api/src/app/medical_record_categories"

type usecase struct {
	repo medical_record_categories.Repository
}

// CountMedicalRecordCategoryByID implements medical_record_categories.Services
func (uc *usecase) CountMedicalRecordCategoryByID(id int) (count int) {
	return uc.repo.CountDataByID(id)
}

// AmendMedicalRecordCategoryByID implements medical_record_categories.Services
func (uc *usecase) AmendMedicalRecordCategoryByID(id int, medical_record_category medical_record_categories.Domain) (err error) {
	err = uc.repo.UpdateByID(id, medical_record_category)
	return
}

// CreateMedicalRecordCategory implements medical_record_categories.Services
func (uc *usecase) CreateMedicalRecordCategory(data medical_record_categories.Domain) (err error) {
	err = uc.repo.InsertData(data)
	return
}

// GetAllMedicalRecordCategories implements medical_record_categories.Services
func (uc *usecase) GetAllMedicalRecordCategories() (medical_record_categories []medical_record_categories.Domain, err error) {
	medical_record_categories, err = uc.repo.SelectAllData()
	return
}

// GetMedicalRecordCategoryByID implements medical_record_categories.Services
func (uc *usecase) GetMedicalRecordCategoryByID(id int) (medical_record_category medical_record_categories.Domain, err error) {
	medical_record_category, err = uc.repo.SelectDataByID(id)
	return
}

// RemoveMedicalRecordCategoryByID implements medical_record_categories.Services
func (uc *usecase) RemoveMedicalRecordCategoryByID(id int) (err error) {
	err = uc.repo.DeleteByID(id)
	return
}

func NewService(repo medical_record_categories.Repository) medical_record_categories.Services {
	return &usecase{repo}
}
