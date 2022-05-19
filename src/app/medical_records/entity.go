package medical_records

import (
	"time"

	"github.com/google/uuid"
)

type GenderType string

const (
	MALE   GenderType = "Male"
	FEMALE GenderType = "Female"
)

type Domain struct {
	ID         uuid.UUID
	Doctor     DoctorReference
	Patient    PatientReference
	MRCategory MRCategoryReference
	MRDetail   MRDetailReference
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DoctorReference struct {
	ID         uuid.UUID
	Name       string
	SIPNumber  string
	Phone      string
	Email      string
	Gender     GenderType
	Polyclinic PolyclinicReference
}

type PolyclinicReference struct {
	ID   int
	Name string
}

type PatientReference struct {
	ID           uuid.UUID
	MRBookNumber string
	Name         string
	Gender       GenderType
	BirthDate    time.Time
}

type MRCategoryReference struct {
	ID   int
	Name string
}

type MRDetailReference struct {
	MRID        uuid.UUID
	ICD         string
	ICDType     string
	Diagnose    []string
	Description string

	// ... will be implemented on the next sprint...
	// Prescription PrescriptionReference
}

// ... will be implemented on next sprint...
// type PrescriptionReference struct {
// 	MRID      uuid.UUID
// 	Medicines []MedicineReference
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

// ... will be implemented on next sprint...
// type MedicineReference struct {
// 	ID          int
// 	Name        string
// 	Dosage      string
// 	Quantity    string
// 	Preparatory string
// }

type Services interface {
	GetAllMedicalRecords() (medicalRecords []Domain, err error)
	GetMedicalRecordByID(id string) (medicalRecord Domain, err error)
	CountMedicalRecordByID(id string) (count int)
	CreateMedicalRecord(domain Domain) (id string, err error)
	AmendMedicalRecordByID(id string, domain Domain) (err error)
	DeleteMedicalRecordByID(id string) (err error)
}

type Repositories interface {
	SelectAllData() (data []Domain, err error)
	SelectDataByID(id string) (selected Domain, err error)
	FindICDData(icdCode string) (diagnose MRDetailReference, err error)
	CountDataByID(id string) (count int)
	InsertData(domain Domain) (id string, err error)
	UpdateByID(id string, domain Domain) (err error)
	DeleteByID(id string) (err error)
}
