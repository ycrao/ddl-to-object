package cmd


import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(javaCmd)
}

var javaCmd = &cobra.Command{
	Use: "java",
	Short: "generate java target object file",
}