package services

import (
	"digimer-api/src/app/medical_records"
	"digimer-api/src/app/medical_records/mocks"
	errormessages "digimer-api/src/constants/error_messages"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	mockRepo          mocks.Repositories
	services          medical_records.Services
	sampleUUID        uuid.UUID
	sampleDoctorUUID  uuid.UUID
	samplePatientUUID uuid.UUID
	sampleDomainList  []medical_records.Domain
	sampleInputDomain medical_records.Domain
	sampleDomain      medical_records.Domain
)

func TestMain(m *testing.M) {
	services = NewService(&mockRepo)
	birthdate, _ := time.Parse("2006-01-02", "2000-05-10")
	sampleUUID = uuid.Must(uuid.NewRandom())
	samplePatientUUID = uuid.Must(uuid.NewRandom())
	sampleDoctorUUID = uuid.Must(uuid.NewRandom())

	sampleInputDomain = medical_records.Domain{
		Doctor: medical_records.DoctorReference{
			ID: sampleDoctorUUID,
		},
		Patient: medical_records.PatientReference{
			ID: samplePatientUUID,
		},
		MRCategory: medical_records.MRCategoryReference{
			ID: 1,
		},
		MRDetail: medical_records.MRDetailReference{
			ICD: "A75.0",
		},
	}

	sampleDomain = medical_records.Domain{
		ID: sampleUUID,
		Doctor: medical_records.DoctorReference{
			ID:        sampleDoctorUUID,
			Name:      "dr. Strange",
			SIPNumber: "SIP.DRHATI.838.111",
			Phone:     "+62 888 8888 8888",
			Email:     "drhati@bucin.net",
			Gender:    medical_records.MALE,
			Polyclinic: medical_records.PolyclinicReference{
				ID:   1,
				Name: "Poli Hati",
			},
		},
		Patient: medical_records.PatientReference{
			ID:           samplePatientUUID,
			MRBookNumber: "12881038",
			Name:         "Damar Danendra",
			Gender:       medical_records.MALE,
			BirthDate:    birthdate,
		},
		MRCategory: medical_records.MRCategoryReference{
			ID:   1,
			Name: "Outpatient",
		},
		MRDetail: medical_records.MRDetailReference{
			MRID: sampleUUID,
			ICD:  "A75.0",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	sampleDomainList = []medical_records.Domain{sampleDomain}

	os.Exit(m.Run())
}

func TestGetAllMedicalRecord(t *testing.T) {
	t.Run("should got all medical_records", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(sampleDomainList, nil).Once()
		result, err := services.GetAllMedicalRecords()

		assert.Nil(t, err)
		assert.Greater(t, len(result), 0)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetAllMedicalRecords()

		assert.NotNil(t, err)
	})

	t.Run("should return nil on empty table", func(t *testing.T) {
		mockRepo.On("SelectAllData").Return(nil, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetAllMedicalRecords()

		assert.NotNil(t, err)
	})
}

func TestGetMedicalRecordByID(t *testing.T) {
	t.Run("should got a medical record by id", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID.String()).Return(sampleDomain, nil).Once()
		result, err := services.GetMedicalRecordByID(sampleDomain.ID.String())

		assert.Nil(t, err)
		assert.Equal(t, sampleDomain, result)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID.String()).Return(medical_records.Domain{}, errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.GetMedicalRecordByID(sampleDomain.ID.String())

		assert.NotNil(t, err)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("SelectDataByID", sampleDomain.ID.String()).Return(medical_records.Domain{}, errors.New(errormessages.FoundNoData)).Once()
		_, err := services.GetMedicalRecordByID(sampleDomain.ID.String())

		assert.NotNil(t, err)
	})
}

func TestCreateMedicalRecord(t *testing.T) {
	sampleInputWithICD := sampleInputDomain
	sampleInputWithICD.MRDetail = medical_records.MRDetailReference{
		ICD:         "A75.0",
		ICDType:     "ICD-10-CM",
		Diagnose:    []string{"Classical typhus (fever)", "Epidemic (louse-borne) typhus"},
		Description: "Epidemic louse-borne typhus fever due to Rickettsia prowazekii",
	}

	t.Run("should successfully added data", func(t *testing.T) {
		mockRepo.On("FindICDData", sampleInputDomain.MRDetail.ICD).Return(sampleInputWithICD.MRDetail, nil).Once()
		mockRepo.On("InsertData", sampleInputWithICD).Return(sampleUUID.String(), nil).Once()
		id, err := services.CreateMedicalRecord(sampleInputDomain)

		assert.Nil(t, err)
		assert.NotNil(t, id)
	})

	t.Run("should got error fetching icd server", func(t *testing.T) {
		mockRepo.On("FindICDData", sampleInputDomain.MRDetail.ICD).Return(medical_records.MRDetailReference{}, errors.New("external server didn't reached")).Once()
		_, err := services.CreateMedicalRecord(sampleInputDomain)

		assert.NotNil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("FindICDData", sampleInputDomain.MRDetail.ICD).Return(sampleInputWithICD.MRDetail, nil).Once()
		mockRepo.On("InsertData", sampleInputWithICD).Return(uuid.Nil.String(), errors.New(errormessages.CannotConnectDatabase)).Once()
		_, err := services.CreateMedicalRecord(sampleInputDomain)

		assert.NotNil(t, err)
	})
}

func TestAmendMedicalRecord(t *testing.T) {
	ICDDetail := medical_records.MRDetailReference{
		ICD:         "A75.0",
		ICDType:     "ICD-10-CM",
		Diagnose:    []string{"Classical typhus (fever)", "Epidemic (louse-borne) typhus"},
		Description: "Epidemic louse-borne typhus fever due to Rickettsia prowazekii",
	}
	sampleInputWithICD := sampleInputDomain
	sampleInputWithICD.MRDetail = ICDDetail
	sampleInputWithICD.MRDetail.MRID = sampleInputDomain.ID

	t.Run("should successfully updated data", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("FindICDData", sampleInputDomain.MRDetail.ICD).Return(ICDDetail, nil).Once()
		mockRepo.On("UpdateByID", sampleUUID.String(), sampleInputWithICD).Return(nil).Once()
		err := services.AmendMedicalRecordByID(sampleUUID.String(), sampleInputDomain)

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("FindICDData", sampleInputDomain.MRDetail.ICD).Return(ICDDetail, nil).Once()
		mockRepo.On("UpdateByID", sampleUUID.String(), sampleInputWithICD).Return(errors.New(errormessages.CannotConnectDatabase))
		err := services.AmendMedicalRecordByID(sampleUUID.String(), sampleInputDomain)

		assert.NotNil(t, err)
	})

	t.Run("should got error fetching icd server", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("FindICDData", sampleInputDomain.MRDetail.ICD).Return(medical_records.MRDetailReference{}, errors.New("external server didn't reached")).Once()
		err := services.AmendMedicalRecordByID(sampleUUID.String(), sampleInputDomain)

		assert.NotNil(t, err)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(0).Once()
		err := services.AmendMedicalRecordByID(sampleUUID.String(), sampleInputDomain)

		assert.NotNil(t, err)
	})
}

func TestRemoveMedicalRecord(t *testing.T) {
	t.Run("should successfully updated data", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("DeleteByID", sampleUUID.String()).Return(nil).Once()
		err := services.DeleteMedicalRecordByID(sampleUUID.String())

		assert.Nil(t, err)
	})

	t.Run("should got database error", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		mockRepo.On("DeleteByID", sampleUUID.String()).Return(errors.New(errormessages.CannotConnectDatabase)).Once()
		err := services.DeleteMedicalRecordByID(sampleUUID.String())

		assert.NotNil(t, err)
	})

	t.Run("should got error while data not found", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(0).Once()
		err := services.DeleteMedicalRecordByID(sampleUUID.String())

		assert.NotNil(t, err)
	})
}

func TestCountMedicalRecordByID(t *testing.T) {
	t.Run("should got counted mrecord by id", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(1).Once()
		count := services.CountMedicalRecordByID(sampleUUID.String())

		assert.Equal(t, 1, count)
	})

	t.Run("should return zero while not found", func(t *testing.T) {
		mockRepo.On("CountDataByID", sampleUUID.String()).Return(0).Once()
		count := services.CountMedicalRecordByID(sampleUUID.String())

		assert.Equal(t, 0, count)
	})
}
