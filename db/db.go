package db

import (
	"os"

	"gorm.io/driver/sqlite" // Sqlite driver based on CGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

func open(dbName string) (*gorm.DB, error) {
	// make the temp directory if it doesn't exist
	err := os.MkdirAll("/tmp", 0755)
	if err != nil {
		return nil, err
	}

	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}

func MustOpen(dbName string) *gorm.DB {
	db, err := open(dbName)
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Project{}, &ResumeLink{}, &ResumeEntry{})
	if err != nil {
		panic(err)
	}

	return db
}
