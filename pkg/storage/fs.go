package storage

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/deifyed/infect/pkg/config"
	"github.com/spf13/viper"
)

func (s *Store) open() error {
	storePath := viper.GetString(config.StorePath)

	content, err := s.Fs.ReadFile(storePath)
	if err != nil {
		return fmt.Errorf("reading store file: %w", err)
	}

	err = json.Unmarshal(content, &s.paths)
	if err != nil {
		return fmt.Errorf("unmarshalling store file: %w", err)
	}

	return nil
}

func (s *Store) close() error {
	rawStore, err := json.Marshal(s.paths)
	if err != nil {
		return fmt.Errorf("marshalling store: %w", err)
	}

	storePath := viper.GetString(config.StorePath)

	err = s.Fs.WriteReader(storePath, bytes.NewReader(rawStore))
	if err != nil {
		return fmt.Errorf("writing store to disk: %w", err)
	}

	return nil
}
