package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(phpCmd)
}

var phpCmd = &cobra.Command{
	Use:   "php",
	Short: "Generate php target object file",
	Run: func(cmd *cobra.Command, args []string) {
		config := CommandConfig{
			Language:       "php",
			DefaultPackage: "App\\Models",
			FileExtension:  ".php",
		}

		if err := ExecuteLanguageCommand(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}
