package repositories

import (
	"digimer-api/src/app/polyclinics"
)

// models
type Polyclinic struct {
	ID   int
	Name string
}

func mapToDomain(record Polyclinic) polyclinics.Domain {
	return polyclinics.Domain{
		ID:   record.ID,
		Name: record.Name,
	}
}

func mapToRecord(domain polyclinics.Domain) Polyclinic {
	return Polyclinic{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func mapToDomainBatch(records []Polyclinic) []polyclinics.Domain {
	domains := []polyclinics.Domain{}

	for _, record := range records {
		domains = append(domains, mapToDomain(record))
	}
	return domains
}
