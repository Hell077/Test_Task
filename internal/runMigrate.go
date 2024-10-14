package internal

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Hell077/Test_Task/internal/database"
	_ "github.com/lib/pq"
)

func RunMigrations() error {
	db := database.Connect()
	defer db.Close()

	if err := runMigrationFile(db, "migrations/000001_music_library.up.sql"); err != nil {
		return err
	}

	log.Println("Migrations have been run successfully.")
	return nil
}

func runMigrationFile(db *sql.DB, migrationFilePath string) error {
	file, err := os.Open(migrationFilePath)
	if err != nil {
		return fmt.Errorf("failed to open migration file: %w", err)
	}
	defer file.Close()

	sqlContent, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	if _, err := db.Exec(string(sqlContent)); err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
}

func CheckExist() bool {
	database.Connect()
	var tableName string
	err := database.DB.QueryRow("SELECT to_regclass('public.songs')").Scan(&tableName)

	if err != nil {
		log.Println("Error querying the database:", err)
		return false
	}
	if tableName == "" {
		return false
	}

	return true
}
