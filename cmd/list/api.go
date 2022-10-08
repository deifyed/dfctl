package list

import (
	"fmt"
	"sort"

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

			trackedPaths = []storage.Path{
				{
					OriginalPath: "/home/deifyed/.config/nvim",
					DotFilesPath: "/home/deifyed/.dotfiles/nvim",
				},
				{
					OriginalPath: "/home/deifyed/.config/sway",
					DotFilesPath: "/home/deifyed/.dotfiles/sway",
					Taint:        true,
				},
				{
					OriginalPath: "/home/deifyed/.config/alacritty",
					DotFilesPath: "/home/deifyed/.dotfiles/alacritty",
				},
			}
			//return nil
		}

		sortedTrackedPaths := sortTrackedPaths(trackedPaths)

		for index, trackedPath := range sortedTrackedPaths {
			taint := ""

			if trackedPath.Taint {
				taint = " (tainted)"
			}

			fmt.Printf("[%d] %-50s %s\n", index, trackedPath.OriginalPath, taint)
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
