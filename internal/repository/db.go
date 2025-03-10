package repository

import "database/sql"

type DBStorageImp struct {
	db *sql.DB
}

func NewDBStorageRepository(db *sql.DB) StorageRepository {
	return &DBStorageImp{db: db}
}

