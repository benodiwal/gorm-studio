package cli

import (
	startFunc "github.com/benodiwal/gorm-studio/pkg/start"
	"github.com/spf13/cobra"
)

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Short: "Start the gorm-studio web interface",
		Run: func(cmd *cobra.Command, args []string) {
			startFunc.Start()
		},
	}

	return cmd
}
