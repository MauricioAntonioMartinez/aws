package backend

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

type PostgreStorage struct {
	db  *sql.DB
	uri string
}

func NewPostgreStorage(uri string) (*PostgreStorage, error) {
	ret := &PostgreStorage{
		uri: uri,
	}

	err := ret.Connect()
	return ret, err
}

func (storage *PostgreStorage) Connect() error {
	db, err := sql.Open("postgres", storage.uri)
	if err != nil {
		return err
	}

	storage.db = db
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)
	db.SetConnMaxLifetime(30 * time.Minute)
	return err
}

