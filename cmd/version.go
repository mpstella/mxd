package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print MxD version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.1")
	},
}
