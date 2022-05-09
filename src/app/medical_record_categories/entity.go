package medical_record_categories

type Domain struct {
	ID   int
	Name string
}

type Services interface {
	GetAllMedicalRecordCategories() (medical_record_categories []Domain, err error)
	GetMedicalRecordCategoryByID(id int) (medical_record_category Domain, err error)
	CountMedicalRecordCategoryByID(id int) (count int)
	CreateMedicalRecordCategory(domain Domain) (err error)
	AmendMedicalRecordCategoryByID(id int, medical_record_category Domain) (err error)
	RemoveMedicalRecordCategoryByID(id int) (err error)
}

type Repository interface {
	SelectAllData() (data []Domain, err error)
	SelectDataByID(id int) (selected Domain, err error)
	CountDataByID(id int) (count int)
	InsertData(domain Domain) (err error)
	UpdateByID(id int, domain Domain) (err error)
	DeleteByID(id int) (err error)
}
