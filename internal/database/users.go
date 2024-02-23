package database

func (b *Brgt) CreateUser(u *User) error {
	return b.DB.Create(u).Error
}

func (b *Brgt) GetUserByID(id int) (*User, error) {
	var u User
	if err := b.DB.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (b *Brgt) GetUserByUsername(username string) (*User, error) {
	var u User
	if err := b.DB.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (b *Brgt) GetUsers() ([]User, error) {
	var users []User
	if err := b.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (b *Brgt) UpdateUser(u *User) error {
	return b.DB.Save(u).Error
}

func (b *Brgt) DeleteUser(id int) error {
	return b.DB.Delete(&User{}, id).Error
}

func (b *Brgt) AddProjectToUser(userID, projectID int) error {
	return b.DB.Model(&User{ID: userID}).Association("Projects").Append(&Project{ID: projectID})
}

func (b *Brgt) RemoveProjectFromUser(userID, projectID int) error {
	return b.DB.Model(&User{ID: userID}).Association("Projects").Delete(&Project{ID: projectID})
}

func (b *Brgt) GetProjectsFromUser(id int) ([]Project, error) {
	var projects []Project
	if err := b.DB.Model(&User{ID: id}).Association("Projects").Find(&projects); err != nil {
		return nil, err
	}
	return projects, nil
}
