package storage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
)

func (s *Store) open() error {
	if s.StorePath == "" {
		return ErrStorePathEmpty
	}

	content, err := s.Fs.ReadFile(s.StorePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			s.paths = make([]Path, 0)

			return nil
		}

		return fmt.Errorf("reading store file: %w", err)
	}

	err = json.Unmarshal(content, &s.paths)
	if err != nil {
		return fmt.Errorf("unmarshalling store file: %w", err)
	}

	return nil
}

func (s *Store) close() error {
	if s.StorePath == "" {
		return ErrStorePathEmpty
	}

	rawStore, err := json.Marshal(s.paths)
	if err != nil {
		return fmt.Errorf("marshalling store: %w", err)
	}

	err = s.Fs.WriteReader(s.StorePath, bytes.NewReader(rawStore))
	if err != nil {
		return fmt.Errorf("writing store to disk: %w", err)
	}

	return nil
}
