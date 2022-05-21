package repositories

import (
	"digimer-api/src/app/medical_records"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// FindICDData implements medical_records.Repositories
func (repo *repository) FindICDData(icdCode string) (diagnose medical_records.MRDetailReference, err error) {
	endpoint := fmt.Sprintf("http://icd10api.com/?code=%s&desc=long&r=json", icdCode)
	resp, err := http.Get(endpoint)
	if err != nil {
		return medical_records.MRDetailReference{}, err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var body ICDResponse
	json.Unmarshal(bodyBytes, &body)

	diagnose = medical_records.MRDetailReference{
		ICD:         icdCode,
		ICDType:     body.ICDType,
		Diagnose:    body.Diagnoses,
		Description: body.Description,
	}
	return diagnose, nil
}
