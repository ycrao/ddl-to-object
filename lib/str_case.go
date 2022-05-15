package lib

// see https://github.com/douyasi/common/tree/master/str
// test see https://github.com/douyasi/common/blob/master/str/case_test.go
import (
	"regexp"
	"strings"
)

var numberSequence = regexp.MustCompile(`([a-zA-Z])(\d+)([a-zA-Z]?)`)
var numberReplacement = []byte(`$1 $2 $3`)

// addWordBoundariesToNumbers helper
func addWordBoundariesToNumbers(s string) string {
	b := []byte(s)
	b = numberSequence.ReplaceAll(b, numberReplacement)
	return string(b)
}

/*--SNAKE STYLE START--*/

// Snake Convert a string to snake case : fooBar --> foo_bar
func Snake(s string) string {
	head, tail := headTailCount(s, '_')
	return strings.Repeat("_", head) + strings.Join(words(s), "_") + strings.Repeat("_", tail)
}

// CamelSnake case is a variant of Snake case with each element's first letter uppercased : fooBar --> Foo_Bar
func CamelSnake(s string) string {
	head, tail := headTailCount(s, '_')
	return strings.Repeat("_", head) + strings.Join(camel(words(s), 0), "_") + strings.Repeat("_", tail)
}

// ScreamingSnake case is a variant of snake case with all letters uppercased : fooBar --> FOO_BAR
func ScreamingSnake(s string) string {
	head, tail := headTailCount(s, '_')
	return strings.Repeat("_", head) + strings.Join(scream(words(s)), "_") + strings.Repeat("_", tail)
}

/*--SNAKE STYLE END--*/

/*--KEBAB STYLE START--*/

// Kebab Covert a string to kebab case : fooBar --> foo-bar
func Kebab(s string) string {
	head, tail := headTailCount(s, '-')
	return strings.Repeat("-", head) + strings.Join(words(s), "-") + strings.Repeat("-", tail)
}

// CamelKebab case is a variant of Kebab case with each element's first letter uppercased : fooBar --> Foo-Bar
func CamelKebab(s string) string {
	head, tail := headTailCount(s, '-')
	return strings.Repeat("-", head) + strings.Join(camel(words(s), 0), "-") + strings.Repeat("-", tail)
}

// ScreamingKebab case is a variant of Kebab case with  with all letters uppercased : fooBar --> FOO-BAR
func ScreamingKebab(s string) string {
	head, tail := headTailCount(s, '-')
	return strings.Repeat("-", head) + strings.Join(scream(words(s)), "-") + strings.Repeat("-", tail)
}

/*--KEBAB STYLE END--*/

/*--Camel STYLE START--*/

// LowerCamel see Camel()
func LowerCamel(s string) string {
	return Camel(s)
}

// Camel Covert a string to (lower) camel case : foo_bar --> fooBar
func Camel(s string) string {
	return strings.Join(camel(words(s), 1), "")
}

// UpperCamel see Pascal()
func UpperCamel(s string) string {
	return Pascal(s)
}

// Pascal Covert a string to pascal case also call upper camel case : foo Bar --> FooBar
func Pascal(s string) string {
	return strings.Join(camel(words(s), 0), "")
}

/*--Camel STYLE END--*/

/**--OTHER STYLE START--*/

// Delimited delimited word
func Delimited(s string, d uint8) string {
	delimited := string(d)
	head, tail := headTailCount(s, d)
	return strings.Repeat(delimited, head) + strings.Join(words(s), delimited) + strings.Repeat(delimited, tail)
}

/*--OTHER STYLE END--*/

// words helper
func words(s string) (w []string) {
	s = addWordBoundariesToNumbers(s)
	start := 0
	l := len(s)
	var prevLower, prevUpper bool
Loop:
	for i, c := range s {
		switch c {
		case '-', '_', ' ', '.':
			if start != i {
				w = append(w, strings.ToLower(s[start:i]))
			}
			start = i + 1
			prevLower = false
			prevUpper = false
			continue Loop
		}
		cs := s[i : i+1]
		if strings.ToUpper(cs) == cs {
			prevUpper = true
			if prevLower {
				if i != start {
					w = append(w, strings.ToLower(s[start:i]))
				}
				start = i
				prevLower = false
			}
		} else {
			prevLower = true
			if prevUpper {
				if i-1 != start {
					w = append(w, strings.ToLower(s[start:i-1]))
				}
				start = i - 1
				prevUpper = false
			}
		}
		if i == l-1 {
			w = append(w, strings.ToLower(s[start:]))
		}
	}
	return
}

// camel helper
func camel(s []string, start int) []string {
	for i := start; i < len(s); i++ {
		switch len(s[i]) {
		case 0:
		case 1:
			s[i] = strings.ToUpper(s[i][0:1])
		default:
			s[i] = strings.ToUpper(s[i][0:1]) + s[i][1:]
		}
	}
	return s
}

// scream helper
func scream(s []string) []string {
	for i := 0; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i])
	}
	return s
}

// headTailCount helper
func headTailCount(s string, sub byte) (head, tail int) {
	for i := 0; i < len(s); i++ {
		if s[i] != sub {
			head = i
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != sub {
			tail = len(s) - i - 1
			break
		}
	}
	return
}