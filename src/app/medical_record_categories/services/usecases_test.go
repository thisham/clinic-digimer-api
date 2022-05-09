package services

import (
	"digimer-api/src/app/medical_record_categories"
	"digimer-api/src/app/medical_record_categories/mocks"
	errormessages "digimer-api/src/constants/error_messages"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mockRepo         mocks.Repositories
	services         medical_record_categories.Services
	sampleDomainList []medical_record_categories.Domain
	sampleDomain     medical_record_categories.Domain
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleDomainList = []medical_record_categories.Domain{
		{ID: 1, Name: "Paracetamol 500 mg tab"},
		{ID: 2, Name: "Piracetam 300 mg tab"},
	}
	sampleDomain = medical_record_categories.Domain{
		ID: 1, Name: "Paracetamol 500 mg Tab",
	}

	os.Exit(m.Run())
}

func TestGetAllMedicalRecordCategories(t *testing.T) {
	t.Run("success got all data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllMedicalRecordCategories()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetAllMedicalRecordCategories()

		assert.NotNil(t, err)
	})

	t.Run("fetched empty data", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetAllMedicalRecordCategories()

		assert.NotNil(t, err)
	})
}

func TestGetMedicalRecordCategoryByID(t *testing.T) {
	t.Run("success got requested data", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID).Return(sampleDomain, nil).Once()
		result, err := services.GetMedicalRecordCategoryByID(sampleDomain.ID)

		assert.Nil(t, err)
		assert.Equal(t, sampleDomain.Name, result.Name)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID).Return(medical_record_categories.Domain{}, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetMedicalRecordCategoryByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})

	t.Run("fetched empty data", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID).Return(medical_record_categories.Domain{}, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetMedicalRecordCategoryByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})
}

func TestCountMedicalRecordCategoryByID(t *testing.T) {
	t.Run("success got requested data", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleDomain.ID).Return(1).Once()
		result := services.CountMedicalRecordCategoryByID(sampleDomain.ID)

		assert.Equal(t, 1, result)
	})

	t.Run("cannot connect database or data not found", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleDomain.ID).Return(0).Once()
		result := services.CountMedicalRecordCategoryByID(sampleDomain.ID)

		assert.Equal(t, 0, result)
	})
}

func TestCreateMedicalRecordCategory(t *testing.T) {
	t.Run("success create data", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomain).Return(nil).Once()
		err := services.CreateMedicalRecordCategory(sampleDomain)

		assert.Nil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("InsertData", sampleDomain).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.CreateMedicalRecordCategory(sampleDomain)

		assert.NotNil(t, err)
	})
}

func TestUpdateMedicalRecordCategory(t *testing.T) {
	t.Run("success create data", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleDomain.ID, sampleDomain).Return(nil).Once()
		err := services.AmendMedicalRecordCategoryByID(sampleDomain.ID, sampleDomain)

		assert.Nil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleDomain.ID, sampleDomain).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.AmendMedicalRecordCategoryByID(sampleDomain.ID, sampleDomain)

		assert.NotNil(t, err)
	})

	t.Run("no data found", func(t *testing.T) {
		mockRepo.On("UpdateByID", sampleDomain.ID, sampleDomain).Return(errors.New(errormessages.FoundNoData)).Once()
		err := services.AmendMedicalRecordCategoryByID(sampleDomain.ID, sampleDomain)

		assert.NotNil(t, err)
	})
}

func TestRemoveMedicalRecordCategory(t *testing.T) {
	t.Run("success create data", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleDomain.ID).Return(nil).Once()
		err := services.RemoveMedicalRecordCategoryByID(sampleDomain.ID)

		assert.Nil(t, err)
	})

	t.Run("cannot connect database", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleDomain.ID).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.RemoveMedicalRecordCategoryByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})

	t.Run("no data found", func(t *testing.T) {
		mockRepo.On("DeleteByID", sampleDomain.ID).Return(errors.New(errormessages.FoundNoData)).Once()
		err := services.RemoveMedicalRecordCategoryByID(sampleDomain.ID)

		assert.NotNil(t, err)
	})
}
