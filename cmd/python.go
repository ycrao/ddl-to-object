package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(pythonCmd)
}

var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "Generate python target object file",
	Run: func(cmd *cobra.Command, args []string) {
		config := CommandConfig{
			Language:       "python",
			DefaultPackage: "",
			FileExtension:  ".py",
		}

		if err := ExecuteLanguageCommand(config); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}
