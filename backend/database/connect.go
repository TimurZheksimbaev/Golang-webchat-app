package database

import (
	"database/sql"

	"github.com/TimurZheksimbaev/Golang-webchat/config"
	_ "github.com/lib/pq"
)


type Database struct {
	db *sql.DB

}

func NewDatabase(config *config.AppConfig) (*Database, error) {
	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}