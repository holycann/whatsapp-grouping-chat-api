package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/holycann/whatsapp-grouping-chat-api/cmd/api"
	"github.com/holycann/whatsapp-grouping-chat-api/cmd/config"
	"github.com/holycann/whatsapp-grouping-chat-api/db"
)

func main() {
	db, err := db.NewPostgresStorage(config.Env.DBAddress, config.Env.MaxOpenConns, config.Env.MaxIdleConns, config.Env.MaxIdleTime)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
}
