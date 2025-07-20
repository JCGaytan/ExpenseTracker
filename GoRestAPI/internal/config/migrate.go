package config

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// RunMigrations applies all SQL files in the migrations folder to the SQLite DB
func RunMigrations(dbPath string) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("failed to open db for migration: %v", err)
	}
	defer db.Close()

	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatalf("failed to read migrations dir: %v", err)
	}
	for _, f := range files {
		if f.IsDir() { continue }
		content, err := os.ReadFile("migrations/" + f.Name())
		if err != nil {
			log.Fatalf("failed to read migration %s: %v", f.Name(), err)
		}
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatalf("migration %s failed: %v", f.Name(), err)
		}
		log.Printf("applied migration: %s", f.Name())
	}
}
