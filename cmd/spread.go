package cmd

import (
	"github.com/deifyed/infect/cmd/spread"
	"github.com/spf13/cobra"
)

// spreadCmd represents the spread command
var spreadCmd = &cobra.Command{
	Use:   "spread",
	Short: "Deploy links of dotfiles to the filesystem",
	RunE:  spread.RunE(log.WithField("command", "spread"), fs),
}

func init() {
	rootCmd.AddCommand(spreadCmd)
}
