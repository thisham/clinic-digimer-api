package request

import "digimer-api/src/app/polyclinics"

type Request struct {
	Name string `json:"name"`
}

func (req *Request) MapToDomain() polyclinics.Domain {
	return polyclinics.Domain{
		Name: req.Name,
	}
}
