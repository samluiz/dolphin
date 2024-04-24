package db

import (
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

func Connect() (*sqlx.DB, error) {

	homeDir, err := os.UserHomeDir()

	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(filepath.Join(homeDir, ".fintrack")); os.IsNotExist(err) {
		err = os.Mkdir(filepath.Join(homeDir, ".fintrack"), 0755)

		if err != nil {
			return nil, err
		}
	}

	dbPath := filepath.Join(homeDir, ".fintrack", "database.db")

	db, err := sqlx.Connect("sqlite", dbPath)

	if err != nil {
		return nil, err
	}

	err = initTables(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initTables(db *sqlx.DB) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	schemaFilePath := filepath.Join(currentDir, "backend", "db", "schema.sql")

	schema, err := os.ReadFile(schemaFilePath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}
