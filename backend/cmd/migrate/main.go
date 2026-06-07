package main

import (
	"database/sql"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://sa:sa@localhost:5432/safood?sslmode=disable"
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	migrationsDir := "../migrations"
	entries := []string{}
	err = filepath.WalkDir(migrationsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".sql" {
			entries = append(entries, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("reading migrations: %v", err)
	}
	if len(entries) == 0 {
		fmt.Println("no migrations found")
		return
	}
	for _, f := range entries {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			log.Fatalf("read %s: %v", f, err)
		}
		if _, err := db.Exec(string(data)); err != nil {
			log.Fatalf("exec %s: %v", f, err)
		}
		fmt.Println("applied:", f)
	}
	fmt.Println("migrations applied")
}
