package lib

import "testing"

// ref see: https://github.com/jinzhu/inflection/blob/master/inflections_test.go
var inflections = map[string]string{
	"stars":        "star",
	"STARS":        "star",
	"Stars":        "star",
	"articles":     "article",
	"QRTZ_LOCKS":   "qrtz_lock",
	"dept-users":   "dept_user",
	"departments":  "department",
	"DEPTS":        "dept",
	"persons":      "person",
	"people":       "person",
	"spokesmen":    "spokesman",
	"salespersons": "salesperson",
	"salespeople":  "salesperson",
	"movies":       "movie",
	"colors":       "color",
	"categories":   "category",
	"woman":        "woman",
	"women":        "woman",
	"indices":      "index",
	"analyses":     "analysis",
	"children":     "child",
	"news":         "news",
	// should failed but passed? as merchandise/product meaning, it have no singular case, should be `goods`
	"goods":   "good",
	"queries": "query",
	"heroes":  "hero",
	// error word: `heros`
	"heros": "hero",
	// should passed but failed? got `leafe`
	"leaves":      "leaf",
	"sheep":       "sheep",
	"information": "information",
	// error word: `informations` - information is a uncountable noun
	"informations": "information",
}

// TestSingular testing singular
func TestSingular(t *testing.T) {
	for key, value := range inflections {
		if v := Singular(key); v != value {
			t.Errorf("%v's singluar should be %v, but got %v", key, value, v)
		}
	}
}
