package repositories

import (
	"digimer-api/src/app/medical_records"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MedicalRecord struct {
	gorm.Model
	ID                      uuid.UUID
	DoctorID                uuid.UUID
	PatientID               uuid.UUID
	MedicalRecordCategoryID int
	Doctor                  Doctor
	Patient                 Patient
	MedicalRecordCategory   MedicalRecordCategory
	MedicalRecordDetail     MedicalRecordDetail
	// Prescription            Prescription
}

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
	ID           uuid.UUID
	Name         string
	SIPNumber    string
	Phone        string
	Email        string
	Gender       genderType `gorm:"type:enum('Male', 'Female')"`
	PolyclinicID int
	Polyclinic   Polyclinic
}

type Polyclinic struct {
	ID   int
	Name string
}

type Patient struct {
	ID           uuid.UUID
	MRBookNumber string
	Name         string
	Gender       genderType
	BirthDate    time.Time
}

type MedicalRecordCategory struct {
	ID   int
	Name string
}

type MedicalRecordDetail struct {
	MedicalRecordID uuid.UUID
	ICD             string
	ICDType         string
	Diagnose        string
	Description     string

	// ... will be implemented on next sprint...
	// Prescription    Prescription
}

type ICDResponse struct {
	Name        string   `json:"Name"`
	ICDType     string   `json:"Type"`
	Description string   `json:"Description"`
	Diagnoses   []string `json:"Inclusions"`
	Response    string   `json:"Response"`
}

// ... will be implemented on next sprint...
// type Prescription struct {
// 	gorm.Model
// 	MedicalRecordID uuid.UUID
// 	Medicines       []Medicine `gorm:"many2many:prescription_medicine_pivots"`
// }

// type PrescriptionMedicinePivot struct {
// 	PrescriptionID uuid.UUID
// 	MedicineID     uuid.UUID
// 	Dosage         string
// 	Quantity       int
// 	Preparatory    string
// }

// type Medicine struct {
// 	ID   int
// 	Name string
// }

func mapToDomain(record MedicalRecord) medical_records.Domain {
	return medical_records.Domain{
		ID: record.ID,
		Doctor: medical_records.DoctorReference{
			ID:         record.Doctor.ID,
			Name:       record.Doctor.Name,
			SIPNumber:  record.Doctor.SIPNumber,
			Phone:      record.Doctor.Phone,
			Email:      record.Doctor.Email,
			Gender:     medical_records.GenderType(record.Doctor.Gender),
			Polyclinic: medical_records.PolyclinicReference{},
		},
		Patient: medical_records.PatientReference{
			ID:           record.Patient.ID,
			MRBookNumber: record.Patient.MRBookNumber,
			Name:         record.Patient.Name,
			Gender:       medical_records.GenderType(record.Patient.Gender),
			BirthDate:    record.Patient.BirthDate,
		},
		MRCategory: medical_records.MRCategoryReference{
			ID:   record.MedicalRecordCategory.ID,
			Name: record.MedicalRecordCategory.Name,
		},
		MRDetail: medical_records.MRDetailReference{
			MRID:        record.ID,
			ICD:         record.MedicalRecordDetail.ICD,
			ICDType:     record.MedicalRecordDetail.ICDType,
			Diagnose:    strings.Split(record.MedicalRecordDetail.Diagnose, ";"),
			Description: record.MedicalRecordDetail.Diagnose,
		},
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func mapToNewRecord(domain medical_records.Domain) MedicalRecord {
	currentDataUuid := uuid.Must(uuid.NewRandom())
	return MedicalRecord{
		ID:                      currentDataUuid,
		DoctorID:                domain.Doctor.ID,
		PatientID:               domain.Patient.ID,
		MedicalRecordCategoryID: domain.MRCategory.ID,
		MedicalRecordDetail: MedicalRecordDetail{
			MedicalRecordID: currentDataUuid,
			ICD:             domain.MRDetail.ICD,
			ICDType:         domain.MRDetail.ICDType,
			Diagnose:        strings.Join(domain.MRDetail.Diagnose, ";"),
			Description:     domain.MRDetail.Description,
		},
	}
}

func mapToExistingRecord(domain medical_records.Domain) MedicalRecord {
	return MedicalRecord{
		ID:                      domain.ID,
		DoctorID:                domain.Doctor.ID,
		PatientID:               domain.Patient.ID,
		MedicalRecordCategoryID: domain.MRCategory.ID,
		MedicalRecordDetail: MedicalRecordDetail{
			MedicalRecordID: domain.MRDetail.MRID,
			ICD:             domain.MRDetail.ICD,
			ICDType:         domain.MRDetail.ICDType,
			Diagnose:        strings.Join(domain.MRDetail.Diagnose, ";"),
			Description:     domain.MRDetail.Description,
		},
	}
}

func mapToDomainBatch(records []MedicalRecord) []medical_records.Domain {
	var domains []medical_records.Domain

	for _, record := range records {
		domains = append(domains, mapToDomain(record))
	}
	return domains
}
