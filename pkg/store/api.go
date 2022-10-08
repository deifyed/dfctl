package store

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/afero"
)

func Add(fs *afero.Afero, targetPath string, dotFilesPath string) error {
	store, err := open(fs)
	if err != nil {
		return fmt.Errorf("opening store: %w", err)
	}

	store.Paths[targetPath] = dotFilesPath

	err = close(fs, store)
	if err != nil {
		return fmt.Errorf("closing store: %w", err)
	}

	return nil
}

func Get(fs *afero.Afero, targetPath string) (string, error) {
	store, err := open(fs)
	if err != nil {
		return "", fmt.Errorf("opening store: %w", err)
	}

	dotFilesPath, ok := store.Paths[targetPath]
	if !ok {
		return "", fmt.Errorf("path not found in store")
	}

	return dotFilesPath, nil
}

func GetAll(fs *afero.Afero) ([]string, error) {
	store, err := open(fs)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []string{}, nil
		}

		return nil, fmt.Errorf("opening store: %w", err)
	}

	var paths []string
	for path := range store.Paths {
		paths = append(paths, path)
	}

	return paths, nil
}
