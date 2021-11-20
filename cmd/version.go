package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Check the version number of Gopic you used",
	Long:  "Check the version number of Gopic you used",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gopic v0.01")
	},
}
