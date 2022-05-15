package cmd


import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(goCmd)
}

var goCmd = &cobra.Command{
	Use: "go",
	Short: "generate golang target object file",
}