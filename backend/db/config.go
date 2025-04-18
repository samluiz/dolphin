package db

import (
	"dolphin/backend/shared/utils"
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

	Migrate(db)

	return db, nil
}

func createAppDir(homeDir string) error {
	if _, err := os.Stat(filepath.Join(homeDir, ".dolphin")); os.IsNotExist(err) {
			err = os.Mkdir(filepath.Join(homeDir, ".dolphin"), 0755)

			if err != nil {
				return err
			}
		}

	return nil
}

func createCurrentVersionDir(homeDir string, version string) error {
	previousVersionPath, err := getPreviousVersionPath()

	if err != nil {
		return err
	}

	if _, err := os.Stat(filepath.Join(homeDir, ".dolphin", version)); os.IsNotExist(err) {
		err = os.Mkdir(filepath.Join(homeDir, ".dolphin", version), 0755)

		if err != nil {
			return err
		}
	}

	if _, err := os.Stat(filepath.Join(homeDir, ".dolphin", version, "database.db")); os.IsNotExist(err) {
		if previousVersionPath != "" {
			err = copyPreviousDatabase(homeDir, version, previousVersionPath)

			if err != nil {
				return err
			}

			return nil
		}
	}

	return nil
}

func getDbPath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	err = createAppDir(homeDir)

	if err != nil {
		return "", err
	}

	if Version == "" {
		Version = "Build"
	}

	err = createCurrentVersionDir(homeDir, Version)

	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".dolphin", Version, "database.db"), nil
}

func getPreviousVersionPath() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	if _, err := os.Stat(filepath.Join(homeDir, ".dolphin")); os.IsNotExist(err) {
		return "", err
	}

	files, err := os.ReadDir(filepath.Join(homeDir, ".dolphin"))

	if err != nil {
		return "", err
	}

	var versions []string

	for _, file := range files {
		if file.IsDir() {
			versions = append(versions, file.Name())
		}
	}

	if len(versions) == 0 {
		return "", nil
	}

	return filepath.Join(homeDir, ".dolphin", versions[len(versions)-1], "database.db"), nil
}

func copyPreviousDatabase(homeDir string, version string, previousVersionPath string) error {
	_, err := utils.CopyFile(previousVersionPath, filepath.Join(homeDir, ".dolphin", version, "database.db"))

	if err != nil {
		return err
	}

	return nil
}