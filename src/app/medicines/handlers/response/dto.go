package response

import "digimer-api/src/app/medicines"

type Response struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func MapToResponse(domain medicines.Domain) Response {
	return Response{
		ID: domain.ID, Name: domain.Name,
	}
}

func MapToBatchResponse(domains []medicines.Domain) []Response {
	responses := []Response{}

	for _, domain := range domains {
		responses = append(responses, MapToResponse(domain))
	}
	return responses
}
