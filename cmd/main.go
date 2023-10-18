package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/superhacker2002/shop/internal/config"
	"github.com/superhacker2002/shop/internal/repository"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	db, err := sql.Open("postgres", cfg.Db)
	if err != nil {
		log.Fatalf("failed to open connection with database: %v", err)
	}
	defer db.Close()

	repo := repository.New(db)
	_, err = repo.ShelfProducts()
	if err != nil {
		log.Fatalf("failed to get info from database: %v", err)
	}
}
