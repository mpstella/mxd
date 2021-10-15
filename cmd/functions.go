package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"mxd/internal/gcloud"
	"os"
)

var (
	functionName  string
	configPath    string
	sourcePath    string
	gCloudCommand *gcloud.Command
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
	Use:     "list",
	Short:   "List functions",
	Aliases: []string{"ls"},
	Args:    cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		return gcloud.NewCommand("functions", "list").Run()
	},
}

var deployCommand = &cobra.Command{
	Use:     "deploy <function-name> <function-config> <source>",
	Short:   "Deploy a function",
	Aliases: []string{"dep"},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		gcloud.Verbose = Verbose

		if Beta { // preference beta over alpha if user has defined both
			gcloud.SetBeta()
		} else if Alpha {
			gcloud.SetAlpha()
		}

		gCloudCommand = gcloud.NewCommand("functions", "deploy", functionName)

		gCloudCommand.AddListMapping("opts").
			AddMapMapping("update-labels", "set-build-env-vars", "update-build-env-vars").
			AddMapListMapping("remove-labels", "remove-env-vars")

		err := gCloudCommand.ReadConfig(configPath)

		if err != nil {
			fmt.Printf("An error occurred: %s\n", err.Error())
			os.Exit(1)
		}

		if Verbose {
			gCloudCommand.Debug()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		var a []string

		if sourcePath != "" {
			a = append(args, fmt.Sprintf("--source=%s", sourcePath))
		}

		err := gCloudCommand.Run(a...)

		if err != nil {
			fmt.Printf("an error occurred %s\n", err.Error())
			os.Exit(1)
		}
	},
}
