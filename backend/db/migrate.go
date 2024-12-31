package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"path"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

//go:embed migrations/*
var migrations embed.FS

func Migrate(db *sqlx.DB) {
	fmt.Println("Starting migration...")

	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		log.Fatalf("Failed to start transactions: %v", err)
	}

	_, err = tx.Exec(`
	CREATE TABLE IF NOT EXISTS db_version (
		version INTEGER PRIMARY KEY
	)`)

	if err != nil {
		log.Fatalf("Failed to create versioning table: %v", err)
	}

	files, err := fs.ReadDir(migrations, "migrations")

	if err != nil {
			log.Fatalf("Failed to read migrations: %v", err)
	}

	var version int
	var fileVersion int
	var migration []byte

	for _, v := range files {
		file, err := v.Info()

		if err != nil {
			log.Fatalf("Failed to get file information: %v", err)
		}

		log.Default().Printf("Reading migration: %s", file.Name())

		fileVersion, err = strconv.Atoi(strings.Split(file.Name(), "__")[0])

		if err != nil {
			log.Fatalf("Failed to get migration version: %v", err)
		}

		err = tx.QueryRow("SELECT version FROM db_version").Scan(&version)

		if err != nil {
			if err != sql.ErrNoRows {
				log.Fatalf("Failed to get latest migration version: %v", err)
			}
		}

		if fileVersion < version {
			continue
		}

		migration, err = fs.ReadFile(migrations, path.Join("migrations", file.Name()))

		log.Default().Printf("%s", string(migration))

		if err != nil {
			log.Fatalf("Failed to read migration file: %v", err)
		}
	}

	if string(migration) != "" {
		_, err := tx.Exec(string(migration))

		if err != nil {
			log.Fatalf("Failed to execute migration: %v", err)
		}

		if fileVersion != 0 && fileVersion > version {
			_, err = tx.Exec("INSERT INTO db_version (version) VALUES (?)", fileVersion)
	
			if err != nil {
				log.Fatalf("Failed to update migration version: %v", err)
			}
		}
	}

	err = tx.Commit()

	if err != nil {
			log.Fatalf("Failed to commit migration transaction: %v", err)
	}

	log.Default().Printf("End of migration.")
}