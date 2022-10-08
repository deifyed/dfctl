package cmd

import (
	"github.com/deifyed/infect/cmd/list"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List infected paths",
	Long:  "This will list all tracked files and folders",
	RunE:  list.RunE(fs),
}

func init() {
	rootCmd.AddCommand(listCmd)
}
