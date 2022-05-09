package services

import (
	"digimer-api/src/app/medicines"
	"digimer-api/src/app/medicines/mocks"
	errormessages "digimer-api/src/constants/error_messages"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockRepo         mocks.Repository
	services         medicines.Services
	sampleDomainList []medicines.Domain
	sampleDomain     medicines.Domain
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleDomainList = []medicines.Domain{
		{ID: 1, Name: "Paracetamol 500 mg tab"},
		{ID: 2, Name: "Piracetam 300 mg tab"},
	}
	sampleDomain = medicines.Domain{
		ID: 1, Name: "Paracetamol 500 mg Tab",
	}

	os.Exit(m.Run())
}

func TestGetAllMedicines(t *testing.T) {
	t.Run("success got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllMedicines()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetAllMedicines()

		assert.NotNil(t, err)
	})

	t.Run("fetched empty data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetAllMedicines()

		assert.NotNil(t, err)
	})
}

func TestGetMedicineByID(t *testing.T) {
	t.Run("success got requested data", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID).Return(sampleDomain, nil).Once()
		result, err := services.GetMedicineByID(sampleDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, sampleDomain.Name, result.Name)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID).Return(medicines.Domain{}, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetMedicineByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})

	t.Run("fetched empty data", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID).Return(medicines.Domain{}, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetMedicineByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestCountMedicineByID(t *testing.T) {
	t.Run("success got requested data", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleDomain.ID).Return(1).Once()
		result := services.CountMedicineByID(sampleDomain.ID)

		assert.Equal(t, 1, result)
	})

	t.Run("cannot connect database or data not found", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleDomain.ID).Return(0).Once()
		result := services.CountMedicineByID(sampleDomain.ID)

		assert.Equal(t, 0, result)
	})
}

func TestCreateMedicine(t *testing.T) {
	t.Run("success create data", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomain).Return(nil).Once()
		err := services.CreateMedicine(sampleDomain)

		assert.Nil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomain).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.CreateMedicine(sampleDomain)

		assert.NotNil(t, err)
	})
}

func TestUpdateMedicine(t *testing.T) {
	t.Run("success create data", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleDomain.ID, sampleDomain).Return(nil).Once()
		err := services.AmendMedicineByID(sampleDomain.ID, sampleDomain)

		assert.Nil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleDomain.ID, sampleDomain).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.AmendMedicineByID(sampleDomain.ID, sampleDomain)

		assert.NotNil(t, err)
	})

	t.Run("no data found", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleDomain.ID, sampleDomain).Return(errors.New(errormessages.FoundNoData)).Once()
		err := services.AmendMedicineByID(sampleDomain.ID, sampleDomain)

		assert.NotNil(t, err)
	})
}

func TestRemoveMedicine(t *testing.T) {
	t.Run("success create data", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleDomain.ID).Return(nil).Once()
		err := services.RemoveMedicineByID(sampleDomain.ID)

		assert.Nil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleDomain.ID).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.RemoveMedicineByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})

	t.Run("no data found", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleDomain.ID).Return(errors.New(errormessages.FoundNoData)).Once()
		err := services.RemoveMedicineByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})
}
