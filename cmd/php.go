package cmd

import (
	"ddl-to-object/lib"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(phpCmd)
}

var phpCmd = &cobra.Command{
	Use:   "php",
	Short: "generate php target object file",
	Run: func(cmd *cobra.Command, args []string) {
		// set default package name to `main`
		phpNamespace := "App\\Models"
		if ns != "" {
			phpNamespace = ns
		}
		toDir := "./"
		if to != "" {
			toDir = strings.TrimRight(to, "/") + "/"
			err := lib.VisitLocationInWriteMode(toDir)
			if err != nil {
				fmt.Errorf(err.Error())
			}
		}
		if from != "" {
			content, _ := lib.ReadFile(from)
			result, err := lib.Parse(string(content))
			if err != nil {
				fmt.Errorf(err.Error())
				os.Exit(0)
			}
			// rewrite PhpNamespaceName
			result.PhpNamespaceName = phpNamespace
			targetFile := toDir + result.PascalObjectName + ".php"
			tpl, err := lib.ReadTemplate("php")
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
