package response

import (
	"digimer-api/src/app/patients"
	"time"
)

type Response struct {
	ID                      string    `json:"id"`
	MedicalRecordBookNumber string    `json:"medical_record_book_number"`
	Name                    string    `json:"name"`
	Gender                  string    `json:"gender"`
	BirthDate               time.Time `json:"birthdate"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}

func MapToResponse(domain patients.Domain) Response {
	return Response{
		ID:                      domain.ID.String(),
		MedicalRecordBookNumber: domain.MRBookNumber,
		Name:                    domain.Name,
		Gender:                  string(domain.Gender),
		BirthDate:               domain.BirthDate,
		CreatedAt:               domain.CreatedAt,
		UpdatedAt:               domain.UpdatedAt,
	}
}

func MapToBatchResponse(domains []patients.Domain) []Response {
	responses := []Response{}

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
