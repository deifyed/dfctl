package storage

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/deifyed/infect/pkg/config"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

func open(fs *afero.Afero) (store, error) {
	storePath := viper.GetString(config.StorePath)

	content, err := fs.ReadFile(storePath)
	if err != nil {
		return store{}, fmt.Errorf("reading store file: %w", err)
	}

	var db store

	err = json.Unmarshal(content, &db)
	if err != nil {
		return store{}, fmt.Errorf("unmarshalling store file: %w", err)
	}

	return db, nil
}

func close(fs *afero.Afero, db store) error {
	rawStore, err := json.Marshal(db)
	if err != nil {
		return fmt.Errorf("marshalling store: %w", err)
	}

	storePath := viper.GetString(config.StorePath)

	err = fs.WriteReader(storePath, bytes.NewReader(rawStore))
	if err != nil {
		return fmt.Errorf("writing store to disk: %w", err)
	}

	return nil
}
