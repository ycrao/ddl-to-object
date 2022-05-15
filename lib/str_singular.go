package lib

import "github.com/jinzhu/inflection"

// Singular using inflection library, may having some bugs, see `str_singular_test.go`
// such as `articles|ARTICLES|Article` -> `article`
func Singular(name string) string {
	snakedName := Snake(name)
	return inflection.Singular(snakedName)
}