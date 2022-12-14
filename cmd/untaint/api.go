package untaint

import (
	"fmt"
	"path"

	"github.com/deifyed/dfctl/pkg/config"
	"github.com/deifyed/dfctl/pkg/storage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE(fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]
		storePath := path.Join(viper.GetString(config.DotFilesDir), "paths.json")

		db := storage.Store{Fs: fs, StorePath: storePath}

		trackedPath, err := db.Get(targetPath)
		if err != nil {
			return fmt.Errorf("getting tracked path: %w", err)
		}

		trackedPath.Taint = false

		err = db.Put(trackedPath)
		if err != nil {
			return fmt.Errorf("adding tracked path: %w", err)
		}

		return nil
	}
}
