package db

import "gorm.io/gorm"

type ResumeEntry struct {
	gorm.Model
	Title      string       `json:"name"`
	URL        string       `json:"url"`
	TimePeriod string       `json:"time_period"`
	Position   string       `json:"position"`
	Links      []ResumeLink `json:"links"`
}

type ResumeLink struct {
	gorm.Model
	ResumeEntryID uint
	Title         string
	URL           string
}

type ResumeRepo interface {
	CreateResumeEntry(title, url, time_period, position string, links []ResumeLink) error
	GetAllResumeEntries() ([]ResumeEntry, error)
}

type ResumeStore struct {
	db *gorm.DB
}

func NewResumeStore(db *gorm.DB) *ResumeStore {
	return &ResumeStore{
		db: db,
	}
}

func (s *ResumeStore) CreateResumeEntry(title, url, time_period, position string, links []ResumeLink) error {
	return s.db.Create(&ResumeEntry{
		Title:      title,
		URL:        url,
		TimePeriod: time_period,
		Position:   position,
		Links:      links,
	}).Error
}

func (s *ResumeStore) GetAllResumeEntries() ([]ResumeEntry, error) {
	var entries []ResumeEntry
	err := s.db.Preload("Links").Find(&entries).Error
	if err != nil {
		return nil, err
	}
	return entries, nil
}
