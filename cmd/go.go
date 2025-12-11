package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(goCmd)
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "Generate golang target object file",
	Run: func(cmd *cobra.Command, args []string) {
		config := CommandConfig{
			Language:       "go",
			DefaultPackage: "main",
			FileExtension:  ".go",
		}

		if err := ExecuteLanguageCommand(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}
