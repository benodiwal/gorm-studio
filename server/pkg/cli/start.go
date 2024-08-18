package cli

import "github.com/spf13/cobra"

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "start",
		Short: "Start the gorm-studio web interface",	
	}

	return cmd
}
