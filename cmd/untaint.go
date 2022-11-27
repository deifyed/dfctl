package cmd

import (
	"github.com/deifyed/dfctl/cmd/untaint"
	"github.com/spf13/cobra"
)

// untaintCmd represents the untaint command
var untaintCmd = &cobra.Command{
	Use:   "untaint",
	Short: "Untaint a certain path",
	Long:  "Unmark certain path to not be ignored anymore when relinking dotfiles",
	RunE:  untaint.RunE(fs),
}

func init() {
	rootCmd.AddCommand(untaintCmd)
}
