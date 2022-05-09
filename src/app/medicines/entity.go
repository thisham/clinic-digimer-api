package medicines

type Domain struct {
	ID   int
	Name string
}

type Services interface {
	GetAllMedicines() (medicines []Domain, err error)
	GetMedicineByID(id int) (medicine Domain, err error)
	CountMedicineByID(id int) (count int)
	CreateMedicine(domain Domain) (err error)
	AmendMedicineByID(id int, medicine Domain) (err error)
	RemoveMedicineByID(id int) (err error)
}

type Repositories interface {
	SelectAllData() (data []Domain, err error)
	SelectDataByID(id int) (selected Domain, err error)
	CountDataByID(id int) (count int)
	InsertData(domain Domain) (err error)
	UpdateByID(id int, domain Domain) (err error)
	DeleteByID(id int) (err error)
}
