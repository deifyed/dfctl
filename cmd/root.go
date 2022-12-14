package cmd

import (
	"os"
	"path"

	"github.com/deifyed/dfctl/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const appName = "dfctl"

var (
	log     = logrus.New()
	fs      = &afero.Afero{Fs: afero.NewOsFs()}
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "Manage your dotfiles",
	Long: `dfctl is a tool for managing dotfiles. This is done by moving relevant files
to a centralized directory and symlinking them to the original path.

This enables easy versioning and sharing of dotfiles.
`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		configureLogger(log)

		return nil
	},
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

	dotfilesDir := path.Join(home, ".config", appName, "dotfiles")

	viper.SetDefault(config.DotFilesDir, dotfilesDir)

	rootCmd.PersistentFlags().StringP(config.LogLevel, "l", "info", "Log level")

	err = viper.BindPFlag(config.LogLevel, rootCmd.PersistentFlags().Lookup(config.LogLevel))
	cobra.CheckErr(err)

	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.config/dfctl/config.yaml)",
	)

	rootCmd.PersistentFlags().StringP(
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

	viper.AddConfigPath(path.Join(home, ".config", "dfctl"))
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		cobra.CheckErr(err)
	}
}

func configureLogger(log *logrus.Logger) {
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	logLevel, err := logrus.ParseLevel(viper.GetString(config.LogLevel))
	if err != nil {
		logLevel = logrus.InfoLevel
		log.Warnf("invalid log level: %s", viper.GetString(config.LogLevel))
	}

	log.SetLevel(logLevel)

	log.WithFields(logrus.Fields{
		"log-level":   logLevel,
		"configFile":  viper.ConfigFileUsed(),
		"dotfilesDir": viper.GetString(config.DotFilesDir),
	}).Debug("logger configured")
}
