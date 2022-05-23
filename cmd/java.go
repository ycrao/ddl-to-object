package cmd

import (
	"ddl-to-object/lib"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(javaCmd)
}

var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "generate java target object file",
	Run: func(cmd *cobra.Command, args []string) {
		// set default package name to `main`
		javaPackage := "com.example.sample.domain.entity"
		if pk != "" {
			javaPackage = pk
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
			// rewrite JavaPackageName
			result.JavaPackageName = javaPackage
			targetFile := toDir + result.PascalObjectName + ".java"
			tpl, err := lib.ReadTemplate("java")
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
