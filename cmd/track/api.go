package track

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/deifyed/infect/pkg/config"
	"github.com/deifyed/infect/pkg/storage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunE(fs *afero.Afero) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		targetPath := args[0]

		err := validate(fs, targetPath)
		if err != nil {
			return fmt.Errorf("validating: %w", err)
		}

		err = track(fs, targetPath)
		if err != nil {
			return fmt.Errorf("tracking path %s: %w", targetPath, err)
		}

		return nil
	}
}

const defaultFolderPermissions = 0o700

// track moves target path to dotfilesDir and leaves a symlink
func track(fs *afero.Afero, targetPath string) error {
	dotFilesDir := viper.GetString(config.DotFilesDir)

	dest := path.Join(dotFilesDir, strings.ReplaceAll(targetPath, "/", "-"))

	err := fs.MkdirAll(dotFilesDir, defaultFolderPermissions)
	if err != nil {
		return fmt.Errorf("ensuring dotfiles directory: %w", err)
	}

	db := storage.Store{Fs: fs}

	err = db.Put(storage.Path{OriginalPath: targetPath, DotFilesPath: dest})
	if err != nil {
		return fmt.Errorf("storing path: %w", err)
	}

	err = fs.Rename(targetPath, dest)
	if err != nil {
		return fmt.Errorf("moving directory: %w", err)
	}

	err = os.Symlink(dest, targetPath)
	if err != nil {
		return fmt.Errorf("linking: %w", err)
	}

	return nil
}

func validate(fs *afero.Afero, targetPath string) error {
	exists, err := fs.Exists(targetPath)
	if err != nil {
		return fmt.Errorf("checking existence: %w", err)
	}

	if !exists {
		return fmt.Errorf("path %s does not exist", targetPath)
	}

	return nil
}
