package main

import (
	"database/sql"

	"github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	//"github.com/satori/go.uuid"
)

type Store struct {
	db *sql.DB
}

func NewStore() (*Store, error) {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return nil, errors.Wrap(err, "could not open database")
	}

	return &Store{
		db: db,
	}
}
