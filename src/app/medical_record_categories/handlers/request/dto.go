package request

import "digimer-api/src/app/medical_record_categories"

type Request struct {
	Name string `json:"name"`
}

func (req *Request) MapToDomain() medical_record_categories.Domain {
	return medical_record_categories.Domain{
		Name: req.Name,
	}
}
