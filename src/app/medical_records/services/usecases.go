package services

import (
	"digimer-api/src/app/medical_records"
	errormessages "digimer-api/src/constants/error_messages"
	"errors"

	"github.com/google/uuid"
)

type usecase struct {
	repo medical_records.Repositories
}

// AmendMedicalRecordByID implements medical_records.Services
func (uc *usecase) AmendMedicalRecordByID(id string, domain medical_records.Domain) (err error) {
	if count := uc.repo.CountDataByID(id); count <= 0 {
		return errors.New(errormessages.FoundNoData)
	}

	diagnose, err := uc.repo.FindICDData(domain.MRDetail.ICD)
	if err != nil {
		return
	}

	domain.MRDetail.MRID = uuid.MustParse(id)
	domain.MRDetail = diagnose
	return uc.repo.UpdateByID(id, domain)
}

// CountMedicalRecordByID implements medical_records.Services
func (uc *usecase) CountMedicalRecordByID(id string) (count int) {
	return uc.repo.CountDataByID(id)
}

// CreateMedicalRecord implements medical_records.Services
func (uc *usecase) CreateMedicalRecord(domain medical_records.Domain) (id string, err error) {
	diagnose, err := uc.repo.FindICDData(domain.MRDetail.ICD)
	if err != nil {
		return
	}

	domain.MRDetail = diagnose
	return uc.repo.InsertData(domain)
}

// DeleteMedicalRecordByID implements medical_records.Services
func (uc *usecase) DeleteMedicalRecordByID(id string) (err error) {
	if count := uc.repo.CountDataByID(id); count <= 0 {
		return errors.New(errormessages.FoundNoData)
	}

	return uc.repo.DeleteByID(id)
}

// GetAllMedicalRecords implements medical_records.Services
func (uc *usecase) GetAllMedicalRecords() (medicalRecords []medical_records.Domain, err error) {
	return uc.repo.SelectAllData()
}

// GetMedicalRecordByID implements medical_records.Services
func (uc *usecase) GetMedicalRecordByID(id string) (medicalRecord medical_records.Domain, err error) {
	return uc.repo.SelectDataByID(id)
}

func NewService(repo medical_records.Repositories) medical_records.Services {
	return &usecase{repo}
}
