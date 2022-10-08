package untaint

import (
	"fmt"

	"github.com/deifyed/infect/pkg/storage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func RunE(fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]

		trackedPath, err := storage.Get(fs, targetPath)
		if err != nil {
			return fmt.Errorf("getting tracked path: %w", err)
		}

		trackedPath.Taint = false

		err = storage.Put(fs, trackedPath)
		if err != nil {
			return fmt.Errorf("adding tracked path: %w", err)
		}

		return nil
	}
}
