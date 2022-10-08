package storage

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/afero"
)

func Put(fs *afero.Afero, trackedPath Path) error {
	db, err := open(fs)
	if err != nil {
		return fmt.Errorf("opening store: %w", err)
	}

	upsert(&db, trackedPath)

	err = close(fs, db)
	if err != nil {
		return fmt.Errorf("closing store: %w", err)
	}

	return nil
}

func Get(fs *afero.Afero, targetPath string) (Path, error) {
	db, err := open(fs)
	if err != nil {
		return Path{}, fmt.Errorf("opening store: %w", err)
	}

	for _, path := range db.Paths {
		if path.OriginalPath == targetPath {
			return path, nil
		}
	}

	return Path{}, fmt.Errorf("path not found")
}

func GetAll(fs *afero.Afero) ([]Path, error) {
	db, err := open(fs)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Path{}, nil
		}

		return nil, fmt.Errorf("opening store: %w", err)
	}

	return db.Paths, nil
}
