package cmd

import (
	"github.com/deifyed/infect/cmd/taint"
	"github.com/spf13/cobra"
)

// taintCmd represents the taint command
var taintCmd = &cobra.Command{
	Use:   "taint",
	Short: "Taint a certain path",
	Long:  "Mark certain paths as to not spread them when infecting filesystem",
	RunE:  taint.RunE(fs),
}

func init() {
	rootCmd.AddCommand(taintCmd)
}
