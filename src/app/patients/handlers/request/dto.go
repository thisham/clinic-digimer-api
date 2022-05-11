package request

import (
	"digimer-api/src/app/patients"
	"time"
)

type Request struct {
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birthdate"`
}

func (req *Request) MapToDomain() patients.Domain {
	birthdate, _ := time.Parse("2006-01-02", req.BirthDate)
	return patients.Domain{
		Name:      req.Name,
		Gender:    patients.Gender(req.Gender),
		BirthDate: birthdate,
	}
}
