package list

import (
	"fmt"

	"github.com/deifyed/infect/pkg/storage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func RunE(fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		trackedPaths, err := storage.GetAll(fs)
		if err != nil {
			return fmt.Errorf("getting tracked paths: %w", err)
		}

		if len(trackedPaths) == 0 {
			fmt.Println("No tracked paths found")

			return nil
		}

		for index, path := range trackedPaths {
			fmt.Printf("[%d] %s", index, path)
		}

		return nil
	}
}
