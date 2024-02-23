package database

func (b *Brgt) CreateTask(t *Task) error {
	return b.DB.Create(t).Error
}

func (b *Brgt) GetTaskByID(id int) (*Task, error) {
	var t Task
	if err := b.DB.First(&t, id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (b *Brgt) GetTasks() ([]Task, error) {
	var tasks []Task
	if err := b.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (b *Brgt) UpdateTask(t *Task) error {
	return b.DB.Save(t).Error
}

func (b *Brgt) DeleteTask(id int) error {
	return b.DB.Delete(&Task{}, id).Error
}

func (b *Brgt) GetTasksFromProject(id int) ([]Task, error) {
	var tasks []Task
	if err := b.DB.Model(&Project{ID: id}).Association("Tasks").Find(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (b *Brgt) GetTasksFromUser(id int) ([]Task, error) {
	var tasks []Task
	if err := b.DB.Model(&User{ID: id}).Association("Tasks").Find(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (b *Brgt) GetDocumentsFromTask(id int) ([]Document, error) {
	var documents []Document
	if err := b.DB.Model(&Task{ID: id}).Association("Documents").Find(&documents); err != nil {
		return nil, err
	}
	return documents, nil
}

func (b *Brgt) FindTasksByValue(value string) ([]*Task, error) {
	var t []*Task
	if err := b.DB.Where("name LIKE ? OR description LIKE ?", "%"+value+"%", "%"+value+"%").Find(&t).Error; err != nil {
		return nil, err
	}

	return t, nil
}
