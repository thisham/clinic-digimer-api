package services

import (
	"digimer-api/src/app/patients"
	errormessages "digimer-api/src/constants/error_messages"
	"errors"
	"fmt"
	"strconv"
)

type usecase struct {
	repo patients.Repositories
}

// AmendPatientByID implements patients.Services
func (uc *usecase) AmendPatientByID(id string, domain patients.Domain) (err error) {
	isFound := uc.repo.CountDataByID(id) != 0

	if isFound {
		return uc.repo.UpdateByID(id, domain)
	}
	return errors.New(errormessages.FoundNoData)
}

// CountPatientByID implements patients.Services
func (uc *usecase) CountPatientByID(id string) (count int) {
	return uc.repo.CountDataByID(id)
}

// CreatePatient implements patients.Services
func (uc *usecase) CreatePatient(domain patients.Domain) (id string, err error) {
	// MR Book Number Generate
	latestMRBook, _ := strconv.Atoi(uc.repo.LookupLatestMRBookNumber())

	domain.MRBookNumber = fmt.Sprintf("%08d", latestMRBook+1)
	return uc.repo.InsertData(domain)
}

// GetAllPatients implements patients.Services
func (uc *usecase) GetAllPatients() (patients []patients.Domain, err error) {
	return uc.repo.SelectAllData()
}

// GetPatientByID implements patients.Services
func (uc *usecase) GetPatientByID(id string) (patient patients.Domain, err error) {
	return uc.repo.SelectDataByID(id)
}

// GetPatientByMRBookNumber implements patients.Services
func (uc *usecase) GetPatientByMRBookNumber(mrBookNumber string) (patient patients.Domain, err error) {
	return uc.repo.SelectDataByMRBookNumber(mrBookNumber)
}

// RemovePatientByID implements patients.Services
func (uc *usecase) RemovePatientByID(id string) (err error) {
	isFound := uc.repo.CountDataByID(id) != 0

	if isFound {
		return uc.repo.DeleteByID(id)
	}
	return errors.New(errormessages.FoundNoData)
}

func NewService(repo patients.Repositories) patients.Services {
	return &usecase{repo}
}
