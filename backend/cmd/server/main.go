package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"backend/pkg/handler"

	_ "github.com/lib/pq"
)

func main() {
	// Try to connect using the handler helper which sets pool params.
	if err := handler.ConnectDB(); err != nil {
		log.Println("ConnectDB:", err)
		// leave db nil; handler will fallback to in-memory
	} else {
		log.Println("connected to database")
	}

	// Optionally apply migrations on start (set MIGRATE_ON_START=1)
	if os.Getenv("MIGRATE_ON_START") == "1" {
		log.Println("running migrations on start")
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			dsn = "postgres://sa:sa@localhost:5432/safood?sslmode=disable"
		}
		if err := runMigrations(dsn); err != nil {
			log.Println("migrations error:", err)
		} else {
			log.Println("migrations applied")
		}
	}

	http.HandleFunc("/api/orders", handler.Handler)
	addr := ":8080"
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func runMigrations(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	dir := "../migrations"
	files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
	if err != nil {
		return err
	}
	for _, f := range files {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			return err
		}
		if _, err := db.Exec(string(data)); err != nil {
			return err
		}
	}
	return nil
}
