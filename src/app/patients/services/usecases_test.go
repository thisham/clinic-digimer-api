package services

import (
	"digimer-api/src/app/patients"
	"digimer-api/src/app/patients/mocks"
	errormessages "digimer-api/src/constants/error_messages"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	mockRepo         mocks.Repositories
	services         patients.Services
	sampleDomainList []patients.Domain
	sampleDomain     patients.Domain
	sampleInput      patients.Domain
	sampleUUID       uuid.UUID
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleUUID = uuid.Must(uuid.NewRandom())
	birthdate, _ := time.Parse("2006-01-02", "2000-05-10")
	sampleDomain = patients.Domain{
		ID:           sampleUUID,
		MRBookNumber: "00223198",
		Name:         "Yehezkiel Saragih",
		Gender:       "Male",
		BirthDate:    birthdate,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	sampleInput = patients.Domain{
		MRBookNumber: sampleDomain.MRBookNumber,
		Name:         sampleDomain.Name,
		Gender:       sampleDomain.Gender,
		BirthDate:    sampleDomain.BirthDate,
	}
	sampleDomainList = []patients.Domain{sampleDomain}

	os.Exit(m.Run())
}

func TestGetAllPatients(t *testing.T) {
	t.Run("should got all patients", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllPatients()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.CannotConnectDatabase))
		_, err := services.GetAllPatients()

		assert.NotNil(t, err)
	})

	t.Run("should return nil on empty table", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.FoundNoData))
		_, err := services.GetAllPatients()

		assert.NotNil(t, err)
	})
}

func TestGetPatientByID(t *testing.T) {
	t.Run("should got a patient by id", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID.String()).Return(sampleDomain, nil).Once()
		result, err := services.GetPatientByID(sampleDomain.ID.String())

		assert.Nil(t, err)
		assert.Equal(t, sampleDomain, result)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID.String()).Return(patients.Domain{}, errormessages.CannotConnectDatabase)
		_, err := services.GetPatientByID(sampleDomain.ID.String())

		assert.NotNil(t, err)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID.String()).Return(patients.Domain{}, errormessages.CannotConnectDatabase)
		_, err := services.GetPatientByID(sampleDomain.ID.String())

		assert.NotNil(t, err)
	})
}

func TestGetPatientByMedicalRecordBookNumber(t *testing.T) {
	t.Run("should got a patient by MR Book number", func(t *testing.T) {
		mockRepo.On("SelectDataByMRBookNumber", sampleDomain.MRBookNumber).Return(sampleDomain, nil).Once()
		result, err := services.GetPatientByMRBookNumber(sampleDomain.MRBookNumber)

		assert.Nil(t, err)
		assert.Equal(t, sampleDomain.Name, result.Name)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByMRBookNumber", sampleDomain.MRBookNumber).Return(patients.Domain{}, nil).Once()
		_, err := services.GetPatientByMRBookNumber(sampleDomain.MRBookNumber)

		assert.Nil(t, err)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("SelectDataByMRBookNumber", sampleDomain.MRBookNumber).Return(patients.Domain{}, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetPatientByMRBookNumber(sampleDomain.MRBookNumber)

		assert.Nil(t, err)
	})
}

func TestCreatePatient(t *testing.T) {
	t.Run("should successfully added data", func(t *testing.T) {
		mockRepo.On("LookupLatestMRBookNumber").Return("00223197").Once()
		mockRepo.On("InsertData", sampleInput).Return(sampleDomain.ID.String(), nil).Once()
		uid, err := services.CreatePatient(sampleInput)

		assert.Nil(t, err)
		assert.Equal(t, sampleUUID, uid)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("LookupLatestMRBookNumber").Return("00223197").Once()
		mockRepo.On("InsertData", sampleInput).Return(sampleDomain.ID.String(), errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.CreatePatient(sampleInput)

		assert.Nil(t, err)
	})
}

func TestAmendPatient(t *testing.T) {
	t.Run("should successfully updated data", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("UpdateDataByID", sampleUUID.String(), sampleInput).Return(nil).Once()
		err := services.AmendPatientByID(sampleUUID.String(), sampleInput)

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("UpdateDataByID", sampleUUID.String(), sampleInput).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.AmendPatientByID(sampleUUID.String(), sampleInput)

		assert.NotNil(t, err)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(0).Once()
		err := services.AmendPatientByID(sampleUUID.String(), sampleInput)

		assert.NotNil(t, err)
	})
}

func TestRemovePatient(t *testing.T) {
	t.Run("should successfully updated data", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("DeleteDataByID", sampleUUID.String()).Return(nil).Once()
		err := services.RemovePatientByID(sampleUUID.String())

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("DeleteDataByID", sampleUUID.String()).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.RemovePatientByID(sampleUUID.String())

		assert.NotNil(t, err)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(0).Once()
		err := services.RemovePatientByID(sampleUUID.String())

		assert.NotNil(t, err)
	})
}

func TestCountPatientByID(t *testing.T) {
	t.Run("should got counted patient by MR Book Number", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleDomain.ID.String()).Return(1).Once()
		count := services.CountPatientByID(sampleDomain.ID.String())

		assert.Equal(t, 1, count)
	})

	t.Run("should return zero while not found", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleDomain.ID.String()).Return(0).Once()
		count := services.CountPatientByID(sampleDomain.ID.String())

		assert.Equal(t, 0, count)
	})
}
