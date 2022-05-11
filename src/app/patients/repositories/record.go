package repositories

import (
	"digimer-api/src/app/patients"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type genderType string

const (
	MALE   genderType = "Male"
	FEMALE genderType = "Female"
)

func (gt *genderType) Scan(value interface{}) error {
	*gt = genderType(value.([]byte))
	return nil
}

func (gt genderType) Value() string {
	return string(gt)
}

type Patient struct {
	gorm.Model
	ID                      uuid.UUID
	MedicalRecordBookNumber string `gorm:"unique"`
	Name                    string
	Gender                  genderType `gorm:"sql:gender_type"`
	BirthDate               time.Time
}

func mapToDomain(record Patient) patients.Domain {
	return patients.Domain{
		ID:           record.ID,
		MRBookNumber: record.MedicalRecordBookNumber,
		Name:         record.Name,
		Gender:       patients.Gender(record.Gender.Value()),
		BirthDate:    record.BirthDate,
		CreatedAt:    record.CreatedAt,
		UpdatedAt:    record.UpdatedAt,
	}
}

func mapToDomainBatch(records []Patient) []patients.Domain {
	var domains []patients.Domain

	for _, record := range records {
		domains = append(domains, mapToDomain(record))
	}
	return domains
}

func mapToRecord(domain patients.Domain) Patient {
	return Patient{
		ID:        domain.ID,
		Name:      domain.Name,
		Gender:    genderType(domain.Gender),
		BirthDate: domain.BirthDate,
	}
}
