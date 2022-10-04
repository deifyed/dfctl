package cmd

import (
	"github.com/deifyed/infect/cmd/track"
	"github.com/spf13/cobra"
)

// trackCmd represents the store command
var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(1),
	RunE:  track.RunE(fs),
}

func init() {
	rootCmd.AddCommand(trackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
