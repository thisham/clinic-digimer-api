package response

import "digimer-api/src/app/polyclinics"

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MapToResponse(domain polyclinics.Domain) Response {
	return Response{
		ID: domain.ID, Name: domain.Name,
	}
}

func MapToBatchResponse(domains []polyclinics.Domain) []Response {
	responses := []Response{}

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
