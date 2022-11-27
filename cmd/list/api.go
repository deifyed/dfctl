package list

import (
	"fmt"
	"path"
	"sort"

	"github.com/deifyed/infect/pkg/config"
	"github.com/deifyed/infect/pkg/storage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE(fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		storePath := path.Join(viper.GetString(config.DotFilesDir), "paths.json")
		db := storage.Store{Fs: fs, StorePath: storePath}

		trackedPaths, err := db.GetAll()
		if err != nil {
			return fmt.Errorf("getting tracked paths: %w", err)
		}

		if len(trackedPaths) == 0 {
			fmt.Println("No tracked paths found")

			return nil
		}

		sortedTrackedPaths := sortTrackedPaths(trackedPaths)

		for _, trackedPath := range sortedTrackedPaths {
			taint := ""

			if trackedPath.Taint {
				taint = " (tainted)"
			}

			fmt.Printf("%-50s %s\n", trackedPath.OriginalPath, taint)
		}

		return nil
	}
}

// sortTrackedPaths sorts the tracked paths by their original path.
func sortTrackedPaths(trackedPaths []storage.Path) []storage.Path {
	sort.Slice(trackedPaths, func(i, j int) bool {
		return trackedPaths[i].OriginalPath < trackedPaths[j].OriginalPath
	})

	return trackedPaths
}
