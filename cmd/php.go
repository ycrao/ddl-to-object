package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(phpCmd)
}

var phpCmd = &cobra.Command{
	Use: "php",
	Short: "generate php target object file",
}