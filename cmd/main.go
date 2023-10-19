package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/superhacker2002/shop/internal/config"
	"github.com/superhacker2002/shop/internal/controller"
	"github.com/superhacker2002/shop/internal/repository/postgresql"
	"log"
	"os"
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

	orders := os.Args[1:]

	repo := postgresql.New(db)
	err = controller.New(repo).ShelfProducts(orders)
	if err != nil {
		log.Fatalf("failed to get orders info: %v", err)
	}
}
