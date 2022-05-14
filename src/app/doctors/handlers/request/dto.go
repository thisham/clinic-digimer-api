package request

import "digimer-api/src/app/doctors"

type Request struct {
	Name         string `json:"name"`
	SIPNumber    string `json:"sip_number"`
	Gender       string `json:"gender"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PolyclinicID int    `json:"polyclinic_id"`
}

type UpdateDataRequest struct {
	Name         string `json:"name"`
	SIPNumber    string `json:"sip_number"`
	Gender       string `json:"gender"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	PolyclinicID int    `json:"polyclinic_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdatePasswordRequest struct {
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func (req *Request) MapToDomain() doctors.Domain {
	return doctors.Domain{
		Name:       req.Name,
		SIPNumber:  req.SIPNumber,
		Gender:     doctors.Gender(req.Gender),
		Phone:      req.Phone,
		Email:      req.Email,
		Password:   req.Password,
		Polyclinic: doctors.PolyclinicReference{ID: req.PolyclinicID},
	}
}
