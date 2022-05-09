package services

import "digimer-api/src/app/medicines"

type usecase struct {
	repo medicines.Repository
}

// CountMedicineByID implements medicines.Services
func (uc *usecase) CountMedicineByID(id int) (count int) {
	return uc.repo.CountDataByID(id)
}

// AmendMedicineByID implements medicines.Services
func (uc *usecase) AmendMedicineByID(id int, medicine medicines.Domain) (err error) {
	err = uc.repo.UpdateByID(id, medicine)
	return
}

// CreateMedicine implements medicines.Services
func (uc *usecase) CreateMedicine(data medicines.Domain) (err error) {
	err = uc.repo.InsertData(data)
	return
}

// GetAllMedicines implements medicines.Services
func (uc *usecase) GetAllMedicines() (medicines []medicines.Domain, err error) {
	medicines, err = uc.repo.SelectAllData()
	return
}

// GetMedicineByID implements medicines.Services
func (uc *usecase) GetMedicineByID(id int) (medicine medicines.Domain, err error) {
	medicine, err = uc.repo.SelectDataByID(id)
	return
}

// RemoveMedicineByID implements medicines.Services
func (uc *usecase) RemoveMedicineByID(id int) (err error) {
	err = uc.repo.DeleteByID(id)
	return
}

func NewService(repo medicines.Repository) medicines.Services {
	return &usecase{repo}
}
