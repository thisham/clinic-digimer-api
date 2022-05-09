package polyclinics

type Domain struct {
	ID   int
	Name string
}

type Services interface {
	GetAllPolyclinics() (polyclinics []Domain, err error)
	GetPolyclinicByID(id int) (polyclinic Domain, err error)
	CountPolyclinicByID(id int) (count int)
	CreatePolyclinic(domain Domain) (err error)
	AmendPolyclinicByID(id int, polyclinic Domain) (err error)
	RemovePolyclinicByID(id int) (err error)
}

type Repositories interface {
	SelectAllData() (data []Domain, err error)
	SelectDataByID(id int) (selected Domain, err error)
	CountDataByID(id int) (count int)
	InsertData(domain Domain) (err error)
	UpdateByID(id int, domain Domain) (err error)
	DeleteByID(id int) (err error)
}
