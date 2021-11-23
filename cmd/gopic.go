package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmdFlag struct {
	files  []string
	dryRun bool
}

var rootCmd = &cobra.Command{
	Use:   "gopic",
	Short: "Gopic is a command line tool which moves images from local markdown files to online storage services.",
	Long:  `Gopic is a command line tool which moves images from local markdown files to online storage services.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	rootCmd.Flags().StringSliceVarP(&rootCmdFlag.files, "files", "f", []string{},
		"give a list of files to be processed")
	rootCmd.Flags().BoolVar(&rootCmdFlag.dryRun, "dry-run", false,
		"just show the effect, won't upload any images or modify the original files")
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println()
		log.Fatal(err)
	}
}
