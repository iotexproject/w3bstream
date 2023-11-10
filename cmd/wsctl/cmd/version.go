package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is set during build time using -ldflags
var Version string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of wsctl",
	Long:  `All software has versions. This is wsctl's version.`,
	Run: func(cmd *cobra.Command, args []string) {
		if Version == "" {
			Version = "development" // default version if not set during build
		}
		fmt.Printf("wsctl version: %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
