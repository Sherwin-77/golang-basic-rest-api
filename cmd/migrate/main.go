package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sherwin-77/golang-basic-rest-api/configs"
)

var migrationsDir = filepath.Join(configs.GetConfiguration().Database.Migration.Path)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Expected 'up', 'down', or 'make' subcommands")
	}

	db, err := sql.Open("sqlite3", configs.GetConfiguration().Database.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = createMigrationsTable(db)
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "up":
		err = migrateUp(db)
	case "down":
		err = migrateDown(db)
	case "make":
		if len(os.Args) < 3 {
			log.Fatal("Expected migration name")
		}
		err = createMigrationFile(os.Args[2])
	default:
		log.Fatal("Unknown subcommand")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func removeSuffix(name string) string {
	idx := strings.LastIndex(name, "_")
	if idx == -1 {
		return name
	}
	return name[:idx]
}

func createMigrationsTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

func migrateUp(db *sql.DB) error {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if matched, _ := filepath.Match("*_up.sql", file.Name()); matched {
			name := removeSuffix(file.Name())
			applied, err := isMigrationApplied(db, name)
			if err != nil {
				return err
			}
			if !applied {
				err = applyMigration(db, filepath.Join(migrationsDir, file.Name()))
				if err != nil {
					return err
				}
				err = recordMigration(db, name)
				if err != nil {
					return err
				}

				fmt.Printf("Applied migration: %s\n", name)
			}
		}
	}

	return nil
}

func migrateDown(db *sql.DB) error {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return err
	}

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		if matched, _ := filepath.Match("*_down.sql", file.Name()); matched {
			name := removeSuffix(file.Name())
			applied, err := isMigrationApplied(db, name)
			if err != nil {
				return err
			}
			if applied {
				err = applyMigration(db, filepath.Join(migrationsDir, file.Name()))
				if err != nil {
					return err
				}
				err = removeMigrationRecord(db, name)
				if err != nil {
					return err
				}

				fmt.Printf("Rolled back migration: %s\n", name)
			}
		}
	}

	return nil
}

func isMigrationApplied(db *sql.DB, name string) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM migrations WHERE name = ? LIMIT 1", name).Scan(&count)
	return count > 0, err
}

func applyMigration(db *sql.DB, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(content))
	return err
}

func recordMigration(db *sql.DB, name string) error {
	_, err := db.Exec("INSERT INTO migrations (name) VALUES (?)", name)
	return err
}

func removeMigrationRecord(db *sql.DB, name string) error {
	_, err := db.Exec("DELETE FROM migrations WHERE name = ?", name)
	return err
}

func createMigrationFile(name string) error {
	timestamp := time.Now().Format("20060102150405")
	upFileName := fmt.Sprintf("%s_%s_up.sql", timestamp, name)
	downFileName := fmt.Sprintf("%s_%s_down.sql", timestamp, name)

	upFile, err := os.Create(filepath.Join(migrationsDir, upFileName))
	if err != nil {
		return err
	}
	defer upFile.Close()

	downFile, err := os.Create(filepath.Join(migrationsDir, downFileName))
	if err != nil {
		return err
	}
	defer downFile.Close()

	fmt.Printf("Created migration files: %s, %s\n", upFileName, downFileName)
	return nil
}
