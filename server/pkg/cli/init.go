package cli

import "github.com/spf13/cobra"

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "init",
		Short: "Initialize connection to the database",
	}

	return cmd
}
