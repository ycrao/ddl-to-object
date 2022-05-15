package cmd

import (
	"ddl-to-object/lib"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(pythonCmd)
}

var pythonCmd = &cobra.Command{
	Use:   "python",
	Short: "generate python target object file",
	Run: func(cmd *cobra.Command, args []string) {
		toDir := "./"
		if to != "" {
			err := lib.VisitLocationInWriteMode(to)
			if err != nil {
				fmt.Errorf(err.Error())
			}
			toDir = to
		}
		if from != "" {
			content, _ := lib.ReadFile(from)
			result, err := lib.Parse(string(content))
			if err != nil {
				fmt.Errorf(err.Error())
				os.Exit(0)
			}
			targetFile := toDir + result.SnakeObjectName + ".py"
			tpl, err := lib.ReadTemplate("python")
			if err != nil {
				fmt.Errorf(err.Error())
			}
			file, err := os.Create(targetFile)
			if stdout {
				tpl.Execute(os.Stdout, result)
			}
			if err != nil && !stdout {
				tpl.Execute(os.Stdout, result)
			}
			tpl.Execute(file, result)
		}
	},
}
