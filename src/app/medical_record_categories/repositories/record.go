package repositories

import (
	"digimer-api/src/app/medical_record_categories"
)

// models
type MedicalRecordCategory struct {
	ID   int
	Name string
}

func mapToDomain(record MedicalRecordCategory) medical_record_categories.Domain {
	return medical_record_categories.Domain{
		ID:   record.ID,
		Name: record.Name,
	}
}

func mapToRecord(domain medical_record_categories.Domain) MedicalRecordCategory {
	return MedicalRecordCategory{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func mapToDomainBatch(records []MedicalRecordCategory) []medical_record_categories.Domain {
	domains := []medical_record_categories.Domain{}

	for _, record := range records {
		domains = append(domains, mapToDomain(record))
	}
	return domains
}
