package services

import (
	"digimer-api/src/app/doctors"
	errormessages "digimer-api/src/constants/error_messages"
	"digimer-api/src/utils"
	"errors"
)

type usecase struct {
	repo doctors.Repositories
}

// AmendDoctorByID implements doctors.Services
func (uc *usecase) AmendDoctorByID(id string, domain doctors.Domain) (err error) {
	isFound := uc.repo.CountDataByID(id) > 0

	if !isFound {
		return errors.New(errormessages.FoundNoData)
	}
	return uc.repo.UpdateByID(id, domain)
}

// UpdatePassword implements doctors.Services
func (uc *usecase) AmendPasswordByDoctorID(id string, password string, confirmation string) (err error) {
	isFound := uc.repo.CountDataByID(id) > 0
	if !isFound {
		return errors.New(errormessages.FoundNoData)
	}

	if password != confirmation {
		return errors.New(errormessages.PasswordNotMatch)
	}

	return uc.repo.UpdateByID(id, doctors.Domain{Password: password})
}

// AttemptDoctorLogin implements doctors.Services
func (uc *usecase) AttemptDoctorLogin(email string, password string) (token string, err error) {
	user, err := uc.repo.SelectDataByEmail(email)

	if err != nil {
		return "", err
	}

	isMatch := utils.ValidateHash(password, user.Password)
	if !isMatch {
		return "", errors.New(errormessages.PasswordNotMatch)
	}
	return utils.GenerateJwt(user.ID.String(), utils.DOCTOR)
}

// CountDoctorByID implements doctors.Services
func (uc *usecase) CountDoctorByID(id string) (count int) {
	return uc.repo.CountDataByID(id)
}

// CreateDoctor implements doctors.Services
func (uc *usecase) CreateDoctor(domain doctors.Domain) (id string, err error) {
	return uc.repo.InsertData(domain)
}

// GetAllDoctors implements doctors.Services
func (uc *usecase) GetAllDoctors() (doctors []doctors.Domain, err error) {
	return uc.repo.SelectAllData()
}

// GetDoctorByID implements doctors.Services
func (uc *usecase) GetDoctorByID(id string) (doctor doctors.Domain, err error) {
	return uc.repo.SelectDataByID(id)
}

// RemoveDoctorByID implements doctors.Services
func (uc *usecase) RemoveDoctorByID(id string) (err error) {
	isFound := uc.repo.CountDataByID(id) > 0

	if !isFound {
		return errors.New(errormessages.FoundNoData)
	}
	return uc.repo.DeleteByID(id)
}

func NewService(repo doctors.Repositories) doctors.Services {
	return &usecase{repo}
}
