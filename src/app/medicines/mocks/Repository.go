// Code generated by mockery v2.10.6. DO NOT EDIT.

package mocks

import (
	medicines "digimer-api/src/app/medicines"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CountDataByID provides a mock function with given fields: id
func (_m *Repository) CountDataByID(id int) int {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: id
func (_m *Repository) DeleteByID(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertData provides a mock function with given fields: domain
func (_m *Repository) InsertData(domain medicines.Domain) error {
	ret := _m.Called(domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(medicines.Domain) error); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAllData provides a mock function with given fields:
func (_m *Repository) SelectAllData() ([]medicines.Domain, error) {
	ret := _m.Called()

	var r0 []medicines.Domain
	if rf, ok := ret.Get(0).(func() []medicines.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]medicines.Domain)
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

// SelectDataByID provides a mock function with given fields: id
func (_m *Repository) SelectDataByID(id int) (medicines.Domain, error) {
	ret := _m.Called(id)

	var r0 medicines.Domain
	if rf, ok := ret.Get(0).(func(int) medicines.Domain); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(medicines.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: id, domain
func (_m *Repository) UpdateByID(id int, domain medicines.Domain) error {
	ret := _m.Called(id, domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, medicines.Domain) error); ok {
		r0 = rf(id, domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
