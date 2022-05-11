package services

import (
	"digimer-api/src/app/doctors"
)

type usecase struct {
	repo doctors.Repositories
}

// AmendDoctorByID implements doctors.Services
func (*usecase) AmendDoctorByID(id string, domain doctors.Domain) (err error) {
	panic("unimplemented")
}

// UpdatePassword implements doctors.Services
func (*usecase) AmendPasswordByDoctorID(id string, password string, confirmation string) (err error) {
	panic("unimplemented")
}

// AttemptDoctorLogin implements doctors.Services
func (*usecase) AttemptDoctorLogin(email string, password string) (token string, err error) {
	panic("unimplemented")
}

// CountDoctorByID implements doctors.Services
func (*usecase) CountDoctorByID(id string) (count int) {
	panic("unimplemented")
}

// CountDoctorByMRBookNumber implements doctors.Services
func (*usecase) CountDoctorByMRBookNumber(mrBookNumber string) (count int) {
	panic("unimplemented")
}

// CreateDoctor implements doctors.Services
func (*usecase) CreateDoctor(domain doctors.Domain) (id string, err error) {
	panic("unimplemented")
}

// GetAllDoctors implements doctors.Services
func (*usecase) GetAllDoctors() (doctors []doctors.Domain, err error) {
	panic("unimplemented")
}

// GetDoctorByEmail implements doctors.Services
func (*usecase) GetDoctorByEmail(email string) (password string, err error) {
	panic("unimplemented")
}

// GetDoctorByID implements doctors.Services
func (*usecase) GetDoctorByID(id string) (doctor doctors.Domain, err error) {
	panic("unimplemented")
}

// RemoveDoctorByID implements doctors.Services
func (*usecase) RemoveDoctorByID(id string) (err error) {
	panic("unimplemented")
}

func NewService(repo doctors.Repositories) doctors.Services {
	return &usecase{repo}
}
