package cmd

import (
	"github.com/deifyed/infect/cmd/untaint"
	"github.com/spf13/cobra"
)

// untaintCmd represents the untaint command
var untaintCmd = &cobra.Command{
	Use:   "untaint",
	Short: "Untaint a certain path",
	Long:  "Unmark certain path so infection will spread the path again",
	RunE:  untaint.RunE(fs),
}

func init() {
	rootCmd.AddCommand(untaintCmd)
}
