package cmd

import (
	"ddl-to-object/lib"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	rootCmd.AddCommand(goCmd)
}

var goCmd = &cobra.Command{
	Use:   "go",
	Short: "generate golang target object file",
	Run: func(cmd *cobra.Command, args []string) {
		// set default package name to `main`
		goPackage := "main"
		if pk != "" {
			packageArr := strings.Split(pk, ".")
			if len := len(packageArr); len > 0 {
				goPackage = packageArr[len-1]
			}
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
			// rewrite GoPackageName
			result.GoPackageName = goPackage
			targetFile := toDir + result.SnakeObjectName + "_types.go"
			tpl, err := lib.ReadTemplate("go")
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
