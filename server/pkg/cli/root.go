package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

func init() {
	rootCmd.AddCommand(initCmd())
	rootCmd.AddCommand(startCmd())

	rootCmd.Version = version
	rootCmd.SetVersionTemplate("gorm-studio version {{.Version}}\n")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

var rootCmd = &cobra.Command{
	Use: "gorm-studio",
	Short: "Gorm Studio is a CLI tool for visualizing and interacting with your Gorm-managed databases.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
