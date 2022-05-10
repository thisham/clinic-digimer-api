package services

import "digimer-api/src/app/patients"

type usecase struct {
	repo patients.Repositories
}

// AmendPatientByID implements patients.Services
func (uc *usecase) AmendPatientByID(id string, domain patients.Domain) (err error) {
	panic("unimplemented")
}

// CountPatientByID implements patients.Services
func (uc *usecase) CountPatientByID(id string) (count int) {
	panic("unimplemented")
}

// CreatePatient implements patients.Services
func (uc *usecase) CreatePatient(domain patients.Domain) (id string, err error) {
	panic("unimplemented")
}

// GetAllPatients implements patients.Services
func (uc *usecase) GetAllPatients() (patients []patients.Domain, err error) {
	panic("unimplemented")
}

// GetPatientByID implements patients.Services
func (uc *usecase) GetPatientByID(id string) (patient patients.Domain, err error) {
	panic("unimplemented")
}

// GetPatientByMRBookNumber implements patients.Services
func (uc *usecase) GetPatientByMRBookNumber(mrBookNumber string) (patient patients.Domain, err error) {
	panic("unimplemented")
}

// RemovePatientByID implements patients.Services
func (uc *usecase) RemovePatientByID(id string) (err error) {
	panic("unimplemented")
}

func NewService(repo patients.Repositories) patients.Services {
	return &usecase{repo}
}
