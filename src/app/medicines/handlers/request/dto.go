package request

import "digimer-api/src/app/medicines"

type Request struct {
	Name string `json:"name"`
}

func (req *Request) MapToDomain() medicines.Domain {
	return medicines.Domain{
		Name: req.Name,
	}
}
