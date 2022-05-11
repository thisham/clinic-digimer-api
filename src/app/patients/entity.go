package patients

import (
	"time"

	"github.com/google/uuid"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type Domain struct {
	ID           uuid.UUID
	MRBookNumber string
	Name         string
	Gender       Gender
	BirthDate    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Services interface {
	GetAllPatients() (patients []Domain, err error)
	GetPatientByID(id string) (patient Domain, err error)
	GetPatientByMRBookNumber(mrBookNumber string) (patient Domain, err error)
	CreatePatient(domain Domain) (id string, err error)
	AmendPatientByID(id string, domain Domain) (err error)
	RemovePatientByID(id string) (err error)

	// count
	CountPatientByID(id string) (count int)
}

type Repositories interface {
	SelectAllData() (data []Domain, err error)
	SelectDataByID(id string) (selected Domain, err error)
	SelectDataByMRBookNumber(mrBookNumber string) (selected Domain, err error)
	InsertData(domain Domain) (err error)
	UpdateByID(id string, domain Domain) (err error)
	DeleteByID(id string) (err error)

	// count
	CountDataByID(id string) (count int)

	// utils
	LookupLatestMRBookNumber() (mrBookNumber string)
}
