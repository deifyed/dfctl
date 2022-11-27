package spread

import (
	"fmt"
	"os"
	"path"

	"github.com/deifyed/infect/pkg/config"
	"github.com/deifyed/infect/pkg/storage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type logger interface {
	Debugf(format string, args ...interface{})
}

func RunE(log logger, fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		db := storage.Store{Fs: fs, StorePath: path.Join(viper.GetString(config.DotFilesDir), "paths.json")}

		paths, err := db.GetAll()
		if err != nil {
			return fmt.Errorf("getting all tracked paths: %w", err)
		}

		log.Debugf("found %d tracked paths", len(paths))

		for _, target := range paths {
			log.Debugf("Spreading target %v+", target)

			if target.Taint {
				log.Debugf("Target %v+ is tainted, skipping", target)

				continue
			}

			err = os.Symlink(target.DotFilesPath, target.OriginalPath)
			if err != nil {
				if os.IsExist(err) {
					log.Debugf("Target already exists, skipping")

					continue
				}

				return fmt.Errorf("creating symlink: %w", err)
			}
		}

		return nil
	}
}
