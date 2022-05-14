package repositories

import (
	"digimer-api/src/app/doctors"

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

type Doctor struct {
	gorm.Model
	ID           uuid.UUID
	Name         string
	SIPNumber    string
	Gender       genderType `gorm:"type:enum('Male', 'Female')"`
	Phone        string
	Email        string
	Password     string
	PolyclinicID int
	Polyclinic   Polyclinic
}

type Polyclinic struct {
	ID   int
	Name string
}

func mapToDomain(record Doctor) doctors.Domain {
	return doctors.Domain{
		ID:        record.ID,
		Name:      record.Name,
		SIPNumber: record.SIPNumber,
		Gender:    doctors.Gender(record.Gender),
		Phone:     record.Phone,
		Email:     record.Email,
		Password:  record.Password,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		Polyclinic: doctors.PolyclinicReference{
			ID:   record.Polyclinic.ID,
			Name: record.Polyclinic.Name,
		},
	}
}

func mapToNewRecord(domain doctors.Domain) Doctor {
	return Doctor{
		ID:        uuid.Must(uuid.NewRandom()),
		Name:      domain.Name,
		SIPNumber: domain.SIPNumber,
		Gender:    genderType(domain.Gender),
		Phone:     domain.Phone,
		Email:     domain.Email,
		Password:  domain.Password,
	}
}

func mapToExistingRecord(domain doctors.Domain) Doctor {
	return Doctor{
		ID:           domain.ID,
		Name:         domain.Name,
		SIPNumber:    domain.SIPNumber,
		Gender:       genderType(domain.Gender),
		Phone:        domain.Phone,
		Email:        domain.Email,
		Password:     domain.Password,
		PolyclinicID: domain.Polyclinic.ID,
		Polyclinic: Polyclinic{
			domain.Polyclinic.ID, domain.Polyclinic.Name,
		},
	}
}

func mapToDomainBatch(records []Doctor) []doctors.Domain {
	var domains []doctors.Domain

	for _, record := range records {
		domains = append(domains, mapToDomain(record))
	}
	return domains
}
