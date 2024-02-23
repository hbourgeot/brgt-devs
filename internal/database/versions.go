package database

func (b *Brgt) CreateVersion(v *Version) error {
	return b.DB.Create(v).Error
}

func (b *Brgt) GetVersionByID(id int) (*Version, error) {
	var v Version
	if err := b.DB.First(&v, id).Error; err != nil {
		return nil, err
	}
	return &v, nil
}

func (b *Brgt) GetVersions() ([]Version, error) {
	var versions []Version
	if err := b.DB.Find(&versions).Error; err != nil {
		return nil, err
	}
	return versions, nil
}

func (b *Brgt) UpdateVersion(v *Version) error {
	return b.DB.Save(v).Error
}

func (b *Brgt) DeleteVersion(id int) error {
	return b.DB.Delete(&Version{}, id).Error
}

func (b *Brgt) GetDocumentsFromVersion(id int) ([]Document, error) {
	var documents []Document
	if err := b.DB.Model(&Version{ID: id}).Association("Documents").Find(&documents); err != nil {
		return nil, err
	}
	return documents, nil
}

func (b *Brgt) GetVersionsByValue(value string) ([]*Version, error) {
	var v []*Version
	if err := b.DB.Where("name LIKE ? OR description LIKE ? OR path LIKE ? OR changes LIKE ?", "%"+value+"%", "%"+value+"%", "%"+value+"%", "%"+value+"%").Find(&v).Error; err != nil {
		return nil, err
	}

	return v, nil
}
