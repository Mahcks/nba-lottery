package cmd

import (
	"fmt"

	"github.com/Mahcks/nba-lottery/internal/structures"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Returns the current version of the CLI.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nba-lottery/" + structures.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
