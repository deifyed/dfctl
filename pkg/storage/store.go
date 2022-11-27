package storage

import (
	"fmt"

	"github.com/spf13/afero"
)

type Store struct {
	Fs        *afero.Afero
	StorePath string

	paths []Path
}

func (s *Store) Put(trackedPath Path) error {
	err := s.open()
	if err != nil {
		return fmt.Errorf("opening store: %w", err)
	}

	s.upsert(trackedPath)

	err = s.close()
	if err != nil {
		return fmt.Errorf("closing store: %w", err)
	}

	return nil
}

func (s *Store) Get(targetPath string) (Path, error) {
	err := s.open()
	if err != nil {
		return Path{}, fmt.Errorf("opening store: %w", err)
	}

	for _, path := range s.paths {
		if path.OriginalPath == targetPath {
			return path, nil
		}
	}

	return Path{}, fmt.Errorf("path not found")
}

func (s *Store) GetAll() ([]Path, error) {
	err := s.open()
	if err != nil {
		return nil, fmt.Errorf("opening store: %w", err)
	}

	return s.paths, nil
}

func (s *Store) Delete(targetPath string) error {
	err := s.open()
	if err != nil {
		return fmt.Errorf("opening store: %w", err)
	}

	for index, path := range s.paths {
		if path.OriginalPath == targetPath {
			s.paths = append(s.paths[:index], s.paths[index+1:]...)
		}
	}

	err = s.close()
	if err != nil {
		return fmt.Errorf("closing store: %w", err)
	}

	return nil
}
