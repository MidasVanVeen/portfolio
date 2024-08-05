package db

import "gorm.io/gorm"

type Project struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type ProjectRepo interface {
	CreateProject(name, description, url string) error
	GetAllProjects() ([]Project, error)
}

type ProjectStore struct {
	db *gorm.DB
}

func NewProjectStore(db *gorm.DB) *ProjectStore {
	return &ProjectStore{
		db: db,
	}
}

func (s *ProjectStore) CreateProject(name, description, url string) error {
	return s.db.Create(&Project{
		Name:        name,
		Description: description,
		URL:         url,
	}).Error
}

func (s *ProjectStore) GetAllProjects() ([]Project, error) {
	var projects []Project
	err := s.db.Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}
