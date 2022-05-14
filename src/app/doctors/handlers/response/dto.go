package response

import (
	"digimer-api/src/app/doctors"
	"time"
)

type Response struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	SIPNumber  string `json:"sip_number"`
	Gender     string `json:"gender"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Polyclinic struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"polyclinic"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MapToResponse(domain doctors.Domain) Response {
	return Response{
		ID:        domain.ID.String(),
		Name:      domain.Name,
		SIPNumber: domain.Name,
		Gender:    string(domain.Gender),
		Phone:     domain.Phone,
		Email:     domain.Email,
		Polyclinic: struct {
			ID   int    "json:\"id\""
			Name string "json:\"name\""
		}{ID: domain.Polyclinic.ID, Name: domain.Polyclinic.Name},
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func MapToBatchResponse(domains []doctors.Domain) []Response {
	var responses []Response

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
