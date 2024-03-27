package auditoriumstore

import "gorm.io/gorm"

type sqlStore struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
