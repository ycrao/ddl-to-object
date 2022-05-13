package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ddl-to-object",
	Short: "ddl-to-object help to generate object files in different languages from sql ddl file. lang mysql ddl to target language Object <like: java or php class/golang struct/python object>",
	Long: `As name, a object generator. support languages below:
 - java: generate entity class with snake-style to camelStyle, comments, and package directory support
 - golang: generate to struct with tag and comments
 - php: generate to simple class with namespace and comments support
 - python: generate to simple object with comments support
Complete documentation is available at https://github.com/ycrao/ddl-to-object`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
