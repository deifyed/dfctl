package untrack

import (
	"fmt"

	"github.com/deifyed/infect/pkg/storage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func RunE(fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]

		if err := untrack(fs, targetPath); err != nil {
			return fmt.Errorf("untracking path: %w", err)
		}

		return nil
	}
}

// untrack will unlink the target and return the source file or folder to this location
func untrack(fs *afero.Afero, targetPath string) error {
	trackedPath, err := storage.Get(fs, targetPath)
	if err != nil {
		return fmt.Errorf("storing path: %w", err)
	}

	if err := fs.Remove(targetPath); err != nil {
		return fmt.Errorf("removing symlink: %w", err)
	}

	err = fs.Rename(trackedPath.DotFilesPath, targetPath)
	if err != nil {
		return fmt.Errorf("moving directory: %w", err)
	}

	return nil
}
