package store

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/deifyed/infect/pkg/config"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

type storage struct {
	Paths map[string]string
}

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

func open(fs *afero.Afero) (storage, error) {
	storePath := viper.GetString(config.StorePath)

	content, err := fs.ReadFile(storePath)
	if err != nil {
		return storage{}, fmt.Errorf("reading store file: %w", err)
	}

	var store storage

	err = json.Unmarshal(content, &store)
	if err != nil {
		return storage{}, fmt.Errorf("unmarshalling store file: %w", err)
	}

	return store, nil
}

// close saves the store to disk
func close(fs *afero.Afero, store storage) error {
	rawStore, err := json.Marshal(store)
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
