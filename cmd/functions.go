package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"mxd/gcloudutils"
	"os"
	"path/filepath"
	"strings"
)

var (
	functionName  string
	configPath    string
	sourcePath    string
	gCloudCommand *gcloudutils.GcloudCommand
)

func init() {

	rootCmd.AddCommand(functionsCmd)

	functionsCmd.AddCommand(listCommand)
	functionsCmd.AddCommand(deployCommand)

	deployCommand.Flags().StringVarP(&functionName, "name", "n", "", "name of the function")
	deployCommand.Flags().StringVarP(&configPath, "config", "c", "", "path to the configuration")
	deployCommand.Flags().StringVarP(&sourcePath, "source", "s", "", "path to the source")

	_ = deployCommand.MarkFlagRequired("name")
	_ = deployCommand.MarkFlagRequired("config")
}

var functionsCmd = &cobra.Command{
	Use:   "functions",
	Short: "Operations for cloud functions",
	Long: `Operations for cloud functions, eg;
mxd functions list
mxd functions deploy
`,
}

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "List functions",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		err := gcloudutils.NewGcloudCommand("functions", "list").Run()

		if err != nil {
			fmt.Printf("an error occurred %s\n", err.Error())
			os.Exit(1)
		}
	},
}

var deployCommand = &cobra.Command{
	Use:   "deploy <function-name> <function-config> <source>",
	Short: "Deploy a function",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		if configPath != "" {

			abs, err := filepath.Abs(configPath)

			if err != nil {
				panic(err)
			}

			base := filepath.Base(abs)
			path := filepath.Dir(abs)

			viper.SetConfigName(strings.Split(base, ".")[0])
			viper.AddConfigPath(path)

			if err := viper.ReadInConfig(); err != nil {
				panic(err)
			}

			gCloudCommand = gcloudutils.NewGcloudCommand("functions", "deploy", functionName)
			gCloudCommand.Verbose = Verbose
			gCloudCommand.AddListMapping("opts")
			gCloudCommand.AddMapMapping("update-labels", "set-build-env-vars", "update-build-env-vars")
			gCloudCommand.AddMapListMapping("remove-labels", "remove-env-vars")

			gCloudCommand.ViperBuild()

			if Verbose {
				gCloudCommand.Debug()
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		var a []string

		if sourcePath != "" {
			if Verbose {
				fmt.Printf("Adding source: %s\n", sourcePath)
			}
			a = append(args, fmt.Sprintf("--source=%s", sourcePath))
		}

		err := gCloudCommand.Run(a...)

		if err != nil {
			fmt.Printf("an error occurred %s\n", err.Error())
			os.Exit(1)
		}
	},
}
