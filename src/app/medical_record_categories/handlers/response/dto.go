package response

import "digimer-api/src/app/medical_record_categories"

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MapToResponse(domain medical_record_categories.Domain) Response {
	return Response{
		ID: domain.ID, Name: domain.Name,
	}
}

func MapToBatchResponse(domains []medical_record_categories.Domain) []Response {
	responses := []Response{}

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
