// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	doctors "digimer-api/src/app/doctors"

	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// AmendDoctorByID provides a mock function with given fields: id, domain
func (_m *Services) AmendDoctorByID(id string, domain doctors.Domain) error {
	ret := _m.Called(id, domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, doctors.Domain) error); ok {
		r0 = rf(id, domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AmendPasswordByDoctorID provides a mock function with given fields: id, password, confirmation
func (_m *Services) AmendPasswordByDoctorID(id string, password string, confirmation string) error {
	ret := _m.Called(id, password, confirmation)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(id, password, confirmation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AttemptDoctorLogin provides a mock function with given fields: email, password
func (_m *Services) AttemptDoctorLogin(email string, password string) (string, error) {
	ret := _m.Called(email, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDoctorByID provides a mock function with given fields: id
func (_m *Services) CountDoctorByID(id string) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// CreateDoctor provides a mock function with given fields: domain
func (_m *Services) CreateDoctor(domain doctors.Domain) (string, error) {
	ret := _m.Called(domain)

	var r0 string
	if rf, ok := ret.Get(0).(func(doctors.Domain) string); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(doctors.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllDoctors provides a mock function with given fields:
func (_m *Services) GetAllDoctors() ([]doctors.Domain, error) {
	ret := _m.Called()

	var r0 []doctors.Domain
	if rf, ok := ret.Get(0).(func() []doctors.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]doctors.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDoctorByID provides a mock function with given fields: id
func (_m *Services) GetDoctorByID(id string) (doctors.Domain, error) {
	ret := _m.Called(id)

	var r0 doctors.Domain
	if rf, ok := ret.Get(0).(func(string) doctors.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(doctors.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveDoctorByID provides a mock function with given fields: id
func (_m *Services) RemoveDoctorByID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
