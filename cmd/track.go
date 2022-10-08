package cmd

import (
	"github.com/deifyed/infect/cmd/track"
	"github.com/spf13/cobra"
)

// trackCmd represents the store command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track a file or folder",
	Long:  "Replace a directory with a link and place the original in a configured dotfiles directory",
	Args:  cobra.ExactArgs(1),
	RunE:  track.RunE(fs),
}

func init() {
	rootCmd.AddCommand(trackCmd)
}
