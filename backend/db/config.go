package db

import (
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var Version string

func Connect() (*sqlx.DB, error) {

	dbPath, err := getDbPath()

	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("sqlite", dbPath + "?cache=shared")

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	err = initTables(db)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func initTables(db *sqlx.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func getDbPath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	if _, err := os.Stat(filepath.Join(homeDir, ".dolphin")); os.IsNotExist(err) {
		err = os.Mkdir(filepath.Join(homeDir, ".dolphin"), 0755)

		if err != nil {
			return "", err
		}
	}

	if Version == "" {
		Version = "Build"
	}

	if _, err := os.Stat(filepath.Join(homeDir, ".dolphin", Version)); os.IsNotExist(err) {
		err = os.Mkdir(filepath.Join(homeDir, ".dolphin", Version), 0755)

		if err != nil {
			return "", err
		}
	}

	return filepath.Join(homeDir, ".dolphin", Version, "database.db"), nil
}