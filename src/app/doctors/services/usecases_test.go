package services

import (
	"digimer-api/src/app/doctors"
	"digimer-api/src/app/doctors/mocks"
	"digimer-api/src/app/patients"
	errormessages "digimer-api/src/constants/error_messages"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	mockRepo                  mocks.Repositories
	services                  doctors.Services
	sampleDomainList          []doctors.Domain
	samplePassword            string
	sampleDomain              doctors.Domain
	sampleCreateInput         doctors.Domain
	sampleUpdateDataInput     doctors.Domain
	sampleUpdatePasswordInput doctors.Domain
	sampleLoginInput          doctors.Domain
	sampleUUID                uuid.UUID
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	sampleUUID = uuid.Must(uuid.NewRandom())
	samplePassword = "thestrongestpassword"
	sampleDomain = doctors.Domain{
		ID:        sampleUUID,
		Name:      "dr. Damar Danendra",
		Gender:    doctors.Male,
		SIPNumber: "SIP.9999.216.11.12",
		Phone:     "+62 888 8888 8888",
		Email:     "damardanendra@example.com",
		Password:  samplePassword,
		Polyclinic: doctors.PolyclinicReference{
			ID: 1, Name: "General",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	sampleCreateInput = doctors.Domain{
		Name:       "dr. Damar Danendra",
		Gender:     doctors.Male,
		SIPNumber:  "SIP.9999.216.11.12",
		Phone:      "+62 888 8888 8888",
		Email:      "damardanendra@example.com",
		Password:   samplePassword,
		Polyclinic: doctors.PolyclinicReference{ID: 1},
	}

	sampleUpdateDataInput = doctors.Domain{
		Name:       "dr. Damar Danendra",
		Gender:     doctors.Male,
		SIPNumber:  "SIP.9999.216.11.12",
		Phone:      "+62 888 8888 8888",
		Email:      "damardanendra@example.com",
		Polyclinic: doctors.PolyclinicReference{ID: 1},
	}

	sampleUpdatePasswordInput = doctors.Domain{
		Password: samplePassword,
	}

	sampleLoginInput = doctors.Domain{
		Email:    "damardanendra@example.com",
		Password: samplePassword,
	}
}

func TestAmendDoctorByID(t *testing.T) {
	t.Run("should successfully update doctor data", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("UpdateByID", sampleUUID.String(), sampleUpdateDataInput).Return(nil).Once()
		err := services.AmendDoctorByID(sampleUUID.String(), sampleUpdateDataInput)

		assert.Nil(t, err)
	})

	t.Run("should got database error on update", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("UpdateByID", sampleUUID.String(), sampleUpdateDataInput).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.AmendDoctorByID(sampleUUID.String(), sampleUpdateDataInput)

		assert.NotNil(t, err)
	})

	t.Run("should got not found error", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(0).Once()
		err := services.AmendDoctorByID(sampleUUID.String(), sampleUpdateDataInput)

		assert.NotNil(t, err)
	})
}

func TestAmendPasswordByDoctorID(t *testing.T) {
	t.Run("should successfully update the password", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("UpdateByID", sampleUpdatePasswordInput).Return(nil).Once()
		err := services.AmendPasswordByDoctorID(sampleUUID.String(), samplePassword, samplePassword)

		assert.Nil(t, err)
	})

	t.Run("should got not same password error", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		err := services.AmendPasswordByDoctorID(sampleUUID.String(), samplePassword, "weakpassword")

		assert.NotNil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("UpdateByID", sampleUpdatePasswordInput).Return(errors.New(errormessages.CannotConnectDatabase))
		err := services.AmendPasswordByDoctorID(sampleUUID.String(), samplePassword, samplePassword)

		assert.NotNil(t, err)
	})

	t.Run("should got data not found error", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(0).Once()
		err := services.AmendPasswordByDoctorID(sampleUUID.String(), samplePassword, samplePassword)

		assert.NotNil(t, err)
	})
}

func TestAttemptDoctorLogin(t *testing.T) {
	t.Run("should got auth token", func(t *testing.T) {
		mockRepo.On("SelectDataByEmail", sampleDomain.Email).Return(sampleDomain, nil).Once()
		token, err := services.AttemptDoctorLogin(sampleDomain.Email, samplePassword)

		assert.Nil(t, err)
		assert.NotNil(t, token)
	})

	t.Run("should return error while password did not match", func(t *testing.T) {
		mockRepo.On("SelectDataByEmail", sampleDomain.Email).Return(sampleDomain, nil).Once()
		token, err := services.AttemptDoctorLogin(sampleDomain.Email, samplePassword)

		assert.Nil(t, token)
		assert.NotNil(t, err)
	})

	t.Run("should got email not registered error", func(t *testing.T) {
		mockRepo.On("SelectDataByEmail", sampleDomain.Email).Return(patients.Domain{}, errors.New(errormessages.FoundNoData)).Once()
		token, err := services.AttemptDoctorLogin(sampleDomain.Email, samplePassword)

		assert.Nil(t, token)
		assert.NotNil(t, err)
	})
}

func TestCountDoctorByID(t *testing.T) {
	t.Run("should found at least one data", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		count := services.CountDoctorByID(sampleUUID.String())

		assert.Greater(t, count, 0)
	})

	t.Run("should found zero data", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(0).Once()
		count := services.CountDoctorByID(sampleUUID.String())

		assert.Equal(t, count, 0)
	})
}

func TestCreateDoctor(t *testing.T) {
	t.Run("should successfully added doctor data", func(t *testing.T) {
		mockRepo.On("InsertData", sampleCreateInput).Return(nil).Once()
		uid, err := services.CreateDoctor(sampleCreateInput)

		assert.Nil(t, err)
		assert.NotEqual(t, uuid.Nil.String(), uid)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("InsertData", sampleCreateInput).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.CreateDoctor(sampleCreateInput)

		assert.NotNil(t, err)
	})
}

func TestGetAllDoctors(t *testing.T) {
	t.Run("should got all doctors", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllDoctors()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetAllDoctors()

		assert.NotNil(t, err)
	})
}

func TestGetDoctorByID(t *testing.T) {
	t.Run("should successfully got a doctor data", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUID.String()).Return(sampleDomain, nil).Once()
		result, err := services.GetDoctorByID(sampleUUID.String())

		assert.Nil(t, err)
		assert.NotEqual(t, uuid.Nil, result.ID)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUID.String()).Return(patients.Domain{}, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetDoctorByID(sampleUUID.String())

		assert.NotNil(t, err)
	})

	t.Run("should got error data not found", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleUUID.String()).Return(patients.Domain{}, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetDoctorByID(sampleUUID.String())

		assert.NotNil(t, err)
	})
}

func TestRemoveDoctorByID(t *testing.T) {
	t.Run("should successfully remove doctor by id", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("DeleteByID", sampleUUID.String()).Return(nil).Once()
		err := services.RemoveDoctorByID(sampleUUID.String())

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("DeleteByID", sampleUUID.String()).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.RemoveDoctorByID(sampleUUID.String())

		assert.NotNil(t, err)
	})

	t.Run("should got error no data found", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(0).Once()
		err := services.RemoveDoctorByID(sampleUUID.String())

		assert.NotNil(t, err)
	})
}
