package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	// --stdout/-s: enable or disable stdout
	stdout bool
	// --from/-f: from a path which a single-table ddl file located
	from string
	// --ns/-n: namespace, alias to --pk/-p
	ns string
	// --pk/-p: package, alias to --ns/-n
	pk string
	// --to/-t: output to a directory, if directory not existed will create
	to string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&stdout, "stdout", "s", false, "enable stdout or not, default set false to disable")
	rootCmd.PersistentFlags().StringVarP(&from, "from", "f", "","from `path` which a single-table DDL file located")
	rootCmd.PersistentFlags().StringVarP(&ns, "ns", "n", "App\\Models", "`namespace` name for php, only in php command")
	rootCmd.PersistentFlags().StringVarP(&pk, "pk", "p", "com.example.sample.domain.entity", "`package` name, only in java or go command" )
	rootCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "output to target `path` or location, create directory automatically if it not existed")
}
var rootCmd = &cobra.Command{
	Use:   "ddl-to-object",
	Short: "ddl-to-object help to generate object files in different languages from sql ddl file.",
	Long: `As name, a object generator. support languages below:
 - java: generate entity class with snake-style to camelStyle, comments, and package directory support
 - golang: generate to struct with tag and comments
 - php: generate to simple class with namespace and comments support
 - python: generate to simple object with comments support
Complete documentation is available at https://github.com/ycrao/ddl-to-object`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) == 0 {
			_ = cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
