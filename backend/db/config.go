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
	schema := `
	CREATE TABLE IF NOT EXISTS profiles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT UNIQUE,
    is_default BOOLEAN
);

CREATE TABLE IF NOT EXISTS categories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS expenses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT,
    amount REAL,
    profile_id INTEGER,
    FOREIGN KEY (profile_id) REFERENCES profiles (id)
);

CREATE TABLE IF NOT EXISTS earnings (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    description TEXT,
    amount REAL,
    profile_id INTEGER,
    FOREIGN KEY (profile_id) REFERENCES profiles (id)
);

CREATE TABLE IF NOT EXISTS expense_category (
    expense_id INTEGER,
    category_id INTEGER,
    PRIMARY KEY (expense_id, category_id),
    FOREIGN KEY (expense_id) REFERENCES expenses (id),
    FOREIGN KEY (category_id) REFERENCES categories (id)
);

INSERT INTO categories (description) SELECT 'Food' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Food');
INSERT INTO categories (description) SELECT 'Transport' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Transport');
INSERT INTO categories (description) SELECT 'Entertainment' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Entertainment');
INSERT INTO categories (description) SELECT 'Health' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Health');
INSERT INTO categories (description) SELECT 'Education' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Education');
INSERT INTO categories (description) SELECT 'Others' WHERE NOT EXISTS (SELECT 1 FROM categories WHERE description = 'Others');
	`

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}
