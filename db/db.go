package db

import (
	"database/sql"
	"log"
	"time"

	"github.com/holycann/whatsapp-grouping-chat-api/cmd/config"
	_ "github.com/lib/pq"
)

func NewPostgresStorage(address string, maxOpenConns, maxIdleConns, maxIdleTime int64) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Env.DBAddress)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(int(maxOpenConns))
	db.SetMaxIdleConns(int(maxIdleConns))
	db.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Second)

	return db, nil
}
