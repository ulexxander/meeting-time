package storage

type OrganizationsStore struct {
	*Store
}

type OrganizationInsertParams struct {
	Name string
}

func (os *OrganizationsStore) GetByID(id uint) (*Organization, error) {
	var item Organization
	if err := os.DB.First(&item, id).Error; err != nil {
		return nil, os.mapError(err)
	}
	return &item, nil
}

func (os *OrganizationsStore) Create(params OrganizationInsertParams) (*Organization, error) {
	item := Organization{Name: params.Name}
	if err := os.DB.Create(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
