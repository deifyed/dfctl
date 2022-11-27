package untrack

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/deifyed/dfctl/pkg/config"
	"github.com/deifyed/dfctl/pkg/storage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE(fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		storePath := path.Join(viper.GetString(config.DotFilesDir), "paths.json")

		targetPath, err := filepath.Abs(args[0])
		if err != nil {
			return fmt.Errorf("getting absolute path: %w", err)
		}

		if err := untrack(fs, storePath, targetPath); err != nil {
			return fmt.Errorf("untracking path: %w", err)
		}

		return nil
	}
}

// untrack will unlink the target and return the source file or folder to this location
func untrack(fs *afero.Afero, storePath string, targetPath string) error {
	db := storage.Store{Fs: fs, StorePath: storePath}

	trackedPath, err := db.Get(targetPath)
	if err != nil {
		return fmt.Errorf("retrieving path data: %w", err)
	}

	if err := fs.Remove(targetPath); err != nil {
		return fmt.Errorf("removing symlink: %w", err)
	}

	err = fs.Rename(trackedPath.DotFilesPath, targetPath)
	if err != nil {
		return fmt.Errorf("moving directory: %w", err)
	}

	if err := db.Delete(targetPath); err != nil {
		return fmt.Errorf("deleting path from database: %w", err)
	}

	return nil
}
