package services

import "digimer-api/src/app/polyclinics"

type usecase struct {
	repo polyclinics.Repositories
}

// CountPolyclinicByID implements polyclinics.Services
func (uc *usecase) CountPolyclinicByID(id int) (count int) {
	return uc.repo.CountDataByID(id)
}

// AmendPolyclinicByID implements polyclinics.Services
func (uc *usecase) AmendPolyclinicByID(id int, polyclinic polyclinics.Domain) (err error) {
	err = uc.repo.UpdateByID(id, polyclinic)
	return
}

// CreatePolyclinic implements polyclinics.Services
func (uc *usecase) CreatePolyclinic(data polyclinics.Domain) (err error) {
	err = uc.repo.InsertData(data)
	return
}

// GetAllPolyclinics implements polyclinics.Services
func (uc *usecase) GetAllPolyclinics() (polyclinics []polyclinics.Domain, err error) {
	polyclinics, err = uc.repo.SelectAllData()
	return
}

// GetPolyclinicByID implements polyclinics.Services
func (uc *usecase) GetPolyclinicByID(id int) (polyclinic polyclinics.Domain, err error) {
	polyclinic, err = uc.repo.SelectDataByID(id)
	return
}

// RemovePolyclinicByID implements polyclinics.Services
func (uc *usecase) RemovePolyclinicByID(id int) (err error) {
	err = uc.repo.DeleteByID(id)
	return
}

func NewService(repo polyclinics.Repositories) polyclinics.Services {
	return &usecase{repo}
}
