package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbose bool

var rootCmd = &cobra.Command{
	Use:   "mxd",
	Short: "Mx Deploy",
	Long:  `mxd is a gcloud wrapper that allows options to be JSON.`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func initConfig() {
	viper.SetEnvPrefix("MXD")
	viper.AutomaticEnv()
}
