// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	patients "digimer-api/src/app/patients"

	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// AmendPatientByID provides a mock function with given fields: id, domain
func (_m *Services) AmendPatientByID(id string, domain patients.Domain) error {
	ret := _m.Called(id, domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, patients.Domain) error); ok {
		r0 = rf(id, domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountPatientByID provides a mock function with given fields: id
func (_m *Services) CountPatientByID(id string) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// CreatePatient provides a mock function with given fields: domain
func (_m *Services) CreatePatient(domain patients.Domain) (string, error) {
	ret := _m.Called(domain)

	var r0 string
	if rf, ok := ret.Get(0).(func(patients.Domain) string); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(patients.Domain) error); ok {
		r1 = rf(domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPatients provides a mock function with given fields:
func (_m *Services) GetAllPatients() ([]patients.Domain, error) {
	ret := _m.Called()

	var r0 []patients.Domain
	if rf, ok := ret.Get(0).(func() []patients.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]patients.Domain)
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

// GetPatientByID provides a mock function with given fields: id
func (_m *Services) GetPatientByID(id string) (patients.Domain, error) {
	ret := _m.Called(id)

	var r0 patients.Domain
	if rf, ok := ret.Get(0).(func(string) patients.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(patients.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPatientByMRBookNumber provides a mock function with given fields: mrBookNumber
func (_m *Services) GetPatientByMRBookNumber(mrBookNumber string) (patients.Domain, error) {
	ret := _m.Called(mrBookNumber)

	var r0 patients.Domain
	if rf, ok := ret.Get(0).(func(string) patients.Domain); ok {
		r0 = rf(mrBookNumber)
	} else {
		r0 = ret.Get(0).(patients.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(mrBookNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemovePatientByID provides a mock function with given fields: id
func (_m *Services) RemovePatientByID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
