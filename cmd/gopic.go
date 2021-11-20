package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gopic",
	Short: "Gopic is a command line tool which moves images from local markdown files to online storage services.",
	Long:  `Gopic is a command line tool which moves images from local markdown files to online storage services.`,
}

func Execute() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
