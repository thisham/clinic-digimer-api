package request

import (
	"digimer-api/src/app/medical_records"

	"github.com/google/uuid"
)

type Request struct {
	PatientID    string `json:"patient_id"`
	MRCategoryID int    `json:"medical_record_catergory_id"`
	ICDCode      string `json:"icd_code"`
}

func MapToDomain(request Request) medical_records.Domain {
	return medical_records.Domain{
		Patient: medical_records.PatientReference{
			ID: uuid.MustParse(request.PatientID),
		},
		MRCategory: medical_records.MRCategoryReference{
			ID: request.MRCategoryID,
		},
		MRDetail: medical_records.MRDetailReference{
			ICD: request.ICDCode,
		},
	}
}
