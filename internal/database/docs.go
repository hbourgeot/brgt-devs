package database

func (b *Brgt) CreateDocumet(d *Document) error {
	return b.DB.Create(d).Error
}

func (b *Brgt) GetDocumentByID(id int) (*Document, error) {
	var d Document
	if err := b.DB.First(&d, id).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (b *Brgt) GetDocuments() ([]Document, error) {
	var documents []Document
	if err := b.DB.Find(&documents).Error; err != nil {
		return nil, err
	}
	return documents, nil
}

func (b *Brgt) UpdateDocument(d *Document) error {
	return b.DB.Save(d).Error
}

func (b *Brgt) DeleteDocument(id int) error {
	return b.DB.Delete(&Document{}, id).Error
}

func (b *Brgt) GetVersionsFromDocument(id int) ([]Version, error) {
	var versions []Version
	if err := b.DB.Model(&Document{ID: id}).Association("Versions").Find(&versions); err != nil {
		return nil, err
	}
	return versions, nil
}
