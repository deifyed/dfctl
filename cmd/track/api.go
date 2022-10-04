package track

import (
	"fmt"

	"github.com/deifyed/infect/pkg/store"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func RunE(fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]

		exists, err := fs.Exists(targetPath)
		if err != nil {
			return fmt.Errorf("checking existence: %w", err)
		}

		if !exists {
			return fmt.Errorf("path %s does not exist", targetPath)
		}

		err = store.Add(targetPath)
		if err != nil {
			return fmt.Errorf("adding path to store: %w", err)
		}

		return nil
	}
}
