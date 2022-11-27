package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/deifyed/infect/pkg/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	fs      = &afero.Afero{Fs: afero.NewOsFs()}
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:          "infect",
	Short:        "Infect a filesystem with dotfiles",
	SilenceUsage: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	dotfilesDir := path.Join(home, ".config", "infect", "dotfiles")

	viper.SetDefault(config.DotFilesDir, dotfilesDir)

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.config/infect/infect.yaml)",
	)

	rootCmd.Flags().StringP(
		config.DotFilesDir,
		"d",
		viper.GetString(config.DotFilesDir),
		"directory where dotfiles are stored",
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".infect" (without extension).
	viper.AddConfigPath(path.Join(home, ".config", "infect"))
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".infect")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
