package repositories

import (
	"digimer-api/src/app/medicines"
)

// models
type Medicine struct {
	ID   int
	Name string
}

func mapToDomain(record Medicine) medicines.Domain {
	return medicines.Domain{
		ID:   record.ID,
		Name: record.Name,
	}
}

func mapToRecord(domain medicines.Domain) Medicine {
	return Medicine{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func mapToDomainBatch(records []Medicine) []medicines.Domain {
	domains := []medicines.Domain{}

	for _, record := range records {
		domains = append(domains, mapToDomain(record))
	}
	return domains
}
