package cmd


import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(pythonCmd)
}

var pythonCmd = &cobra.Command{
	Use: "python",
	Short: "generate python target object file",
}