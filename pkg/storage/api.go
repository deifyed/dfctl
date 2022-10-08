package storage

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/afero"
)

func Add(fs *afero.Afero, targetPath string, dotFilesPath string) error {
	db, err := open(fs)
	if err != nil {
		return fmt.Errorf("opening store: %w", err)
	}

	db.Paths[targetPath] = dotFilesPath

	err = close(fs, db)
	if err != nil {
		return fmt.Errorf("closing store: %w", err)
	}

	return nil
}

func Get(fs *afero.Afero, targetPath string) (string, error) {
	db, err := open(fs)
	if err != nil {
		return "", fmt.Errorf("opening store: %w", err)
	}

	dotFilesPath, ok := db.Paths[targetPath]
	if !ok {
		return "", fmt.Errorf("path not found in store")
	}

	return dotFilesPath, nil
}

func GetAll(fs *afero.Afero) ([]string, error) {
	db, err := open(fs)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []string{}, nil
		}

		return nil, fmt.Errorf("opening store: %w", err)
	}

	var paths []string
	for path := range db.Paths {
		paths = append(paths, path)
	}

	return paths, nil
}
