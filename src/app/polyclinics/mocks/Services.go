// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	polyclinics "digimer-api/src/app/polyclinics"

	mock "github.com/stretchr/testify/mock"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	mock.Mock
}

// AmendPolyclinicByID provides a mock function with given fields: id, polyclinic
func (_m *Services) AmendPolyclinicByID(id int, polyclinic polyclinics.Domain) error {
	ret := _m.Called(id, polyclinic)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, polyclinics.Domain) error); ok {
		r0 = rf(id, polyclinic)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountPolyclinicByID provides a mock function with given fields: id
func (_m *Services) CountPolyclinicByID(id int) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// CreatePolyclinic provides a mock function with given fields: domain
func (_m *Services) CreatePolyclinic(domain polyclinics.Domain) error {
	ret := _m.Called(domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(polyclinics.Domain) error); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPolyclinics provides a mock function with given fields:
func (_m *Services) GetAllPolyclinics() ([]polyclinics.Domain, error) {
	ret := _m.Called()

	var r0 []polyclinics.Domain
	if rf, ok := ret.Get(0).(func() []polyclinics.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]polyclinics.Domain)
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

// GetPolyclinicByID provides a mock function with given fields: id
func (_m *Services) GetPolyclinicByID(id int) (polyclinics.Domain, error) {
	ret := _m.Called(id)

	var r0 polyclinics.Domain
	if rf, ok := ret.Get(0).(func(int) polyclinics.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(polyclinics.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemovePolyclinicByID provides a mock function with given fields: id
func (_m *Services) RemovePolyclinicByID(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
