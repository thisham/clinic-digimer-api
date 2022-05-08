package services

import (
	"digimer-api/src/app/polyclinics"
	"digimer-api/src/app/polyclinics/mocks"
	errormessages "digimer-api/src/constants/error_messages"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockRepo         mocks.Repository
	services         polyclinics.Services
	sampleDomainList []polyclinics.Domain
	sampleDomain     polyclinics.Domain
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleDomainList = []polyclinics.Domain{
		{ID: 1, Name: "General"},
		{ID: 2, Name: "Dentistry"},
	}
	sampleDomain = polyclinics.Domain{
		ID: 1, Name: "General",
	}

	os.Exit(m.Run())
}

func TestGetAllPolyclinics(t *testing.T) {
	t.Run("success got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllPolyclinics()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetAllPolyclinics()

		assert.NotNil(t, err)
	})

	t.Run("fetched empty data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetAllPolyclinics()

		assert.NotNil(t, err)
	})
}

func TestGetPolyclinicByID(t *testing.T) {
	t.Run("success got requested data", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID).Return(sampleDomain, nil).Once()
		result, err := services.GetPolyclinicByID(sampleDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, sampleDomain.Name, result.Name)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID).Return(polyclinics.Domain{}, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetPolyclinicByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})

	t.Run("fetched empty data", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID).Return(polyclinics.Domain{}, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetPolyclinicByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestCreatePolyclinic(t *testing.T) {
	t.Run("success create data", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomain).Return(sampleDomain.ID, nil).Once()
		err := services.CreatePolyclinic(sampleDomain)

		assert.Nil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomain).Return(0, errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.CreatePolyclinic(sampleDomain)

		assert.NotNil(t, err)
	})
}

func TestUpdatePolyclinic(t *testing.T) {
	t.Run("success create data", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleDomain.ID, sampleDomain).Return(nil).Once()
		err := services.AmendPolyclinicByID(sampleDomain.ID, sampleDomain)

		assert.Nil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleDomain.ID, sampleDomain).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.AmendPolyclinicByID(sampleDomain.ID, sampleDomain)

		assert.NotNil(t, err)
	})

	t.Run("no data found", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleDomain.ID, sampleDomain).Return(errors.New(errormessages.FoundNoData)).Once()
		err := services.AmendPolyclinicByID(sampleDomain.ID, sampleDomain)

		assert.NotNil(t, err)
	})
}

func TestRemovePolyclinic(t *testing.T) {
	t.Run("success create data", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleDomain.ID).Return(nil).Once()
		err := services.RemovePolyclinicByID(sampleDomain.ID)

		assert.Nil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleDomain.ID).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.RemovePolyclinicByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})

	t.Run("no data found", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleDomain.ID).Return(errors.New(errormessages.FoundNoData)).Once()
		err := services.RemovePolyclinicByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})
}
