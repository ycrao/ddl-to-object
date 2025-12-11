package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(javaCmd)
}

var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "Generate java target object file",
	Run: func(cmd *cobra.Command, args []string) {
		config := CommandConfig{
			Language:       "java",
			DefaultPackage: "com.example.sample.domain.entity",
			FileExtension:  ".java",
		}

		if err := ExecuteLanguageCommand(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}
