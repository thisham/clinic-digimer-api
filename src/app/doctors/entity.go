package doctors

import (
	"time"

	"github.com/google/uuid"
)

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

type Domain struct {
	ID         uuid.UUID
	Name       string
	SIPNumber  string
	Gender     Gender
	Phone      string
	Email      string
	Password   string
	BirthDate  time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Polyclinic PolyclinicReference
}

type PolyclinicReference struct {
	ID   int
	Name string
}

type Services interface {
	GetAllDoctors() (doctors []Domain, err error)
	GetDoctorByID(id string) (doctor Domain, err error)
	CreateDoctor(domain Domain) (id string, err error)
	AmendDoctorByID(id string, domain Domain) (err error)
	RemoveDoctorByID(id string) (err error)

	// auth
	AttemptDoctorLogin(email, password string) (token string, err error)
	AmendPasswordByDoctorID(id, password, confirmation string) (err error)

	// count
	CountDoctorByID(id string) (count int)
}

type Repositories interface {
	SelectAllData() (data []Domain, err error)
	SelectDataByID(id string) (selected Domain, err error)
	InsertData(domain Domain) (id string, err error)
	UpdateByID(id string, domain Domain) (err error)
	DeleteByID(id string) (err error)

	// auth
	SelectDataByEmail(email string) (selected Domain, err error)

	// count
	CountDataByID(id string) (count int)
}
