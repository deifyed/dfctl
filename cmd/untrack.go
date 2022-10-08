package cmd

import (
	"github.com/deifyed/infect/cmd/untrack"
	"github.com/spf13/cobra"
)

// untrackCmd represents the untrack command
var untrackCmd = &cobra.Command{
	Use:   "untrack",
	Short: "Untrack file or folder",
	Long:  "This will unlink the target and return the source file or folder to this location",
	RunE:  untrack.RunE(fs),
}

func init() {
	rootCmd.AddCommand(untrackCmd)
}
