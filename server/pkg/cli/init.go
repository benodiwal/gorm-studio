package cli

import (
	initFunc "github.com/benodiwal/gorm-studio/pkg/init"
	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "init",
		Short: "Initialize connection to the database",
		Run: func(cmd *cobra.Command, args []string) {
			initFunc.CreateEnvFile()
		},
	}

	return cmd
}
