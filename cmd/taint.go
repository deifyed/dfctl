package cmd

import (
	"github.com/deifyed/dfctl/cmd/taint"
	"github.com/spf13/cobra"
)

// taintCmd represents the taint command
var taintCmd = &cobra.Command{
	Use:   "taint",
	Short: "Taint a certain path",
	Long:  "Mark certain paths as to ignore them when relinking dotfiles",
	RunE:  taint.RunE(fs),
}

func init() {
	rootCmd.AddCommand(taintCmd)
}
