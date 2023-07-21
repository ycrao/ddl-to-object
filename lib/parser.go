package lib

import (
	"errors"
	"regexp"
	"strings"
)

// ParsedResult the result of ddl after parsed
type ParsedResult struct {
	PackageName         string
	GoPackageName       string
	JavaPackageName     string
	NamespaceName       string
	PhpNamespaceName    string
	TableName           string
	NormalizedTableName string
	ObjectName          string
	PascalObjectName    string
	CamelObjectName     string
	SnakeObjectName     string
	TableComment        string
	JavaRandomLong      string
	Columns             []Column
}

// Column sql field
type Column struct {
	Name           string
	SnakeName      string
	PascalName     string
	CamelName      string
	Comment        string
	DataType       string
	AdditionalAttr AdditionalAttr
	JavaType       string
	JavaImport     string
	PhpType        string
	// PythonType       string
	GoType   string
	GoImport string
	GoTag    string
}

// AdditionalAttr additional attribute: such as default value, is unsigned, can nullable?
type AdditionalAttr struct {
	OriFieldStr     string
	DataType        string
	IsUnsigned      bool
	IsAutoIncrement bool
	DefaultValue    string
	Nullable        bool
}

const (
	// TableNameRegex \x60 for `
	TableNameRegex    = `(?im)CREATE\s+TABLE\s+([\x60-zA-Z-_."']+)`
	TableCommentRegex = `(?im)\).*COMMENT\s+["|'](.*)["|']`
	// FieldsRegex \x60 for `
	FieldsRegex = `(?im)([\w\x60"']+)\s+([\w]+).*(\s+COMMENT\s+["|'](.*)["|'])?`
	// ColumnCommentRegex parse column comment
	ColumnCommentRegex = `(?im)COMMENT\s+['|"](.*)['|"]`
	NullableValueRegex = `(?im)DEFAULT\s+NULL`
)

// Parse parse MySQL DDL
func Parse(ddl string) (ParsedResult, error) {
	var parsedResult ParsedResult
	table, tableComment, err1 := parseTable(ddl)
	fieldsStr, err2 := getTableFieldsStr(ddl)
	columns, err3 := parseFields(fieldsStr)
	if err1 == nil && err2 == nil && err3 == nil {
		singularName := Singular(table)
		parsedResult = ParsedResult{
			PackageName:         "",
			GoPackageName:       "main",
			JavaPackageName:     "com.example.sample.domain.entity",
			NamespaceName:       "",
			PhpNamespaceName:    "App\\Models",
			TableName:           table,
			NormalizedTableName: singularName,
			ObjectName:          Pascal(singularName),
			PascalObjectName:    Pascal(singularName),
			CamelObjectName:     Camel(singularName),
			SnakeObjectName:     Snake(singularName),
			TableComment:        tableComment,
			JavaRandomLong:      RandomInt64Str(18),
			Columns:             columns,
		}
		return parsedResult, nil
	}
	return parsedResult, errors.New("fail to parse, please check your ddl file")
}

// parseTable parse table name and comment for ddl string
func parseTable(ddl string) (string, string, error) {
	tableRegexp := regexp.MustCompile(TableNameRegex)
	if sections := tableRegexp.FindStringSubmatch(ddl); len(sections) > 0 {
		// value: `s_blog`.`article`
		oriTableNameStr := sections[1]
		tableNameStr := normalizedName(oriTableNameStr)
		tableNameArr := strings.Split(tableNameStr, ".")
		if len(tableNameArr) > 1 {
			tableNameStr = tableNameArr[1]
		} else {
			tableNameStr = tableNameArr[0]
		}
		tableCommentStr := ""
		tableCommentRegexp := regexp.MustCompile(TableCommentRegex)
		optionStr, _ := getTableOptionStr(ddl)
		if results := tableCommentRegexp.FindStringSubmatch(optionStr); len(results) > 0 {
			// value: '文章'
			oriTableCommentStr := results[1]
			tableCommentStr = normalizedName(oriTableCommentStr)
		} else {
			tableCommentStr = tableNameStr
		}
		// article, 文章, nil
		return tableNameStr, tableCommentStr, nil
	}
	return "", "", errors.New("parse ddl error, cannot found a valid table")
}

// parseFields parse table fields
// you can test regex online by https://regex101.com/
func parseFields(fieldsStr string) ([]Column, error) {
	fieldsRegexp := regexp.MustCompile(FieldsRegex)
	matches := fieldsRegexp.FindAllStringSubmatch(fieldsStr, -1)
	/*
		[`id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id', `id` bigint  ]  found at index 0
		[`user_id` bigint NOT NULL COMMENT '用户id', `user_id` bigint  ]  found at index 1
		[`content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '正文', `content` text  ]  found at index 2
		[`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', `create_time` datetime  ]  found at index 3
		[`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', `update_time` datetime  ]  found at index 4
		[PRIMARY KEY (`id`) PRIMARY KEY  ]
	*/
	var columns []Column
outerLoop:
	for _, match := range matches {
		javaType := ""
		javaImport := ""
		goType := ""
		goImport := ""
		phpType := ""
		// using the first element of matches by example
		// match[0] value:
		// `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
		oriFieldStr := match[0]
		if isValidColumn(oriFieldStr) == true {
			// value: `id`
			oriField := match[1]
			columnName := normalizedName(oriField)
			// value: bigint
			oriFieldType := match[2]
			// value: BIGINT
			upperFieldType := strings.ToUpper(oriFieldType)
			// unsigned?
			isUnsigned := isUnsigned(oriFieldStr)
			// auto increment?
			isAutoIncrement := IsAutoIncrement(oriFieldStr)
			columnComment := getColumnComment(oriFieldStr)
			if columnComment == "" {
				columnComment = columnName
			}
			isNullable := IsNullable(oriFieldStr)
			defaultValue := "UNKNOWN"
			if isNullable == true {
				defaultValue = "NULL"
			}
			additionalAttr := AdditionalAttr{
				DataType:        upperFieldType,
				OriFieldStr:     oriFieldStr,
				IsUnsigned:      isUnsigned,
				IsAutoIncrement: isAutoIncrement,
				DefaultValue:    defaultValue,
				Nullable:        isNullable,
			}
			// java
			javaType, javaImport = MapToJavaType(additionalAttr)
			// php
			phpType, _ = MapToPhpType(additionalAttr)
			// python ignore
			//// ...
			// golang
			goType, goImport = MapToGoType(additionalAttr)
			snakeName := Snake(columnName)
			goTag := "`json:\"" + snakeName + "\" db:\"" + columnName + "\"`"
			column := Column{
				Name:           columnName,
				PascalName:     Pascal(columnName),
				SnakeName:      snakeName,
				CamelName:      Camel(columnName),
				Comment:        columnComment,
				DataType:       upperFieldType,
				AdditionalAttr: additionalAttr,
				JavaType:       javaType,
				JavaImport:     javaImport,
				PhpType:        phpType,
				GoType:         goType,
				GoImport:       goImport,
				GoTag:          goTag,
				// PythonType:       "",
			}
			columns = append(columns, column)
		} else {
			continue outerLoop
		}
	}
	return columns, nil
}

// normalizedName remove " or ' or `
func normalizedName(name string) string {
	name = strings.Replace(name, "`", "", -1)
	name = strings.Replace(name, `"`, "", -1)
	name = strings.Replace(name, `'`, "", -1)
	return name
}

// getTableFieldsStr get table fields str cut off
func getTableFieldsStr(ddl string) (string, error) {
	openedBracketIndex := strings.IndexAny(ddl, `(`)
	closedBracketIndex := strings.LastIndexAny(ddl, `)`)
	if openedBracketIndex != -1 && closedBracketIndex != -1 {
		return ddl[(openedBracketIndex + 1):closedBracketIndex], nil
	}
	return "", errors.New("not a valid ddl string")
}

// getTableOptionStr get table option str cut off
func getTableOptionStr(ddl string) (string, error) {
	closedBracketIndex := strings.LastIndexAny(ddl, `)`)
	if closedBracketIndex != -1 {
		return ddl[(closedBracketIndex + 1):], nil
	}
	return "", errors.New("not a valid ddl string")
}

// isUnsigned is unsigned?
func isUnsigned(field string) bool {
	upperFieldStr := strings.ToUpper(field)
	return strings.Contains(upperFieldStr, "UNSIGNED")
}

// IsAutoIncrement is auto increment?
func IsAutoIncrement(field string) bool {
	upperFieldStr := strings.ToUpper(field)
	return strings.Contains(upperFieldStr, "AUTO_INCREMENT")
}

// IsNullable default value is null?
func IsNullable(field string) bool {
	nullableValueRegexp := regexp.MustCompile(NullableValueRegex)
	if results := nullableValueRegexp.FindStringSubmatch(field); len(results) > 0 {
		return true
	}
	return false
}

// getColumnComment
func getColumnComment(field string) string {
	columnCommentRegexp := regexp.MustCompile(ColumnCommentRegex)
	columnCommentStr := ""
	if results := columnCommentRegexp.FindStringSubmatch(field); len(results) > 0 {
		// '文章'
		oriColumnCommentStr := results[1]
		columnCommentStr = normalizedName(oriColumnCommentStr)
	}
	return columnCommentStr
}

// isValidColumn is a valid column
func isValidColumn(field string) bool {
	upperField := strings.ToUpper(field)
	checkedTexts := "PRIMARY|KEY|FOREIGN|CHECK|UNIQUE|CONSTRAINT|INDEX"
	checkedTextsArr := strings.Split(checkedTexts, "|")
	for _, checkedText := range checkedTextsArr {
		existed := strings.Contains(upperField, checkedText)
		if existed {
			return false
		}
	}
	return true
}
