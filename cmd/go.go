package cmd

import (
	"ddl-to-object/lib"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(goCmd)
}

var goCmd = &cobra.Command{
	Use: "go",
	Short: "generate golang target object file",
	Run: func(cmd *cobra.Command, args []string) {
		// set default package name to `main`
		goPackage := "main"
		if pk != "" {
			packageArr := strings.Split(pk, ".")
			if len := len(packageArr); len > 0 {
				goPackage = packageArr[len - 1]
			}
		}
		toDir := "./"
		if to != "" {
			err := lib.VisitLocationInWriteMode(to)
			if err != nil {
				panic(err)
			}
			toDir = to
		}
		if from != "" {
			content, err := lib.ReadFile(from)
			if err != nil {
				panic(err)
			}
			lib.Parse(string(content))
		}

	},
}