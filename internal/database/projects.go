package database

func (b *Brgt) CreateProject(p *Project) error {
	return b.DB.Create(p).Error
}

func (b *Brgt) GetProjectByID(id int) (*Project, error) {
	var p Project
	if err := b.DB.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (b *Brgt) GetProjects() ([]Project, error) {
	var projects []Project
	if err := b.DB.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (b *Brgt) GetProjectByName(name string) (*Project, error) {
	var p Project
	if err := b.DB.Where("name = ?", name).First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (b *Brgt) UpdateProject(p *Project) error {
	return b.DB.Save(p).Error
}

func (b *Brgt) DeleteProject(id int) error {
	return b.DB.Delete(&Project{}, id).Error
}

func (b *Brgt) GetUsersFromProject(id int) ([]User, error) {
	var users []User
	if err := b.DB.Model(&Project{ID: id}).Association("Users").Find(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (b *Brgt) FindProjectByValue(value string) (*Project, error) {
	var p Project
	if err := b.DB.Where("name LIKE ? OR description LIKE ?", "%"+value+"%", "%"+value+"%").First(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
