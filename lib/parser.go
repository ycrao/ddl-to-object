package lib

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// ParsedResult the result of ddl after parsed
type ParsedResult struct {
	PackageName         string
	NamespaceName       string
	TableName           string
	NormalizedTableName string
	ObjectName          string
	PascalObjectName    string
	CamelObjectName     string
	TableComment        string
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
	GoType string
	GoTag  string
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

// example ddl:
/*
CREATE TABLE `s_blog`.`article` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '正文',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '文章';
*/
const Example = "CREATE TABLE `s_blog`.`article` (\n" +
	"  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',\n" +
	"  `user_id` bigint NOT NULL COMMENT '用户id',\n" +
	"  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '正文',\n" +
	"  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n" +
	"  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n" +
	"  PRIMARY KEY (`id`)\n" +
	") ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '文章';"

const (
	// TableNameRegex \x60 for `
	TableNameRegex    = `(?im)CREATE\s+TABLE\s+([\x60-zA-Z-_."']+)`
	TableCommentRegex = `(?im)\).*COMMENT\s+["|'](.*)["|']`
	// FieldsRegex \x60 for `
	FieldsRegex = `(?im)([\w\x60"']+)\s+([\w\(\),]+).*(\s+COMMENT\s+["|'](.*)["|'])?`
	// ColumnCommentRegex parse column comment
	ColumnCommentRegex = `(?im)COMMENT\s+['|"](.*)['|"]`
	NullableValueRegex = `(?im)DEFAULT\s+NULL`
)

// Parse parse MySQL DDL
func Parse(ddl string) {
	tables, _ := parseTable(ddl)
	fieldsStr, _ := getTableFieldsStr(ddl)
	columns, _ := parseFields(fieldsStr)
	bytes, _ := json.Marshal(columns)
	fmt.Printf("%v", tables)
	// fmt.Printf("%v", columns)
	fmt.Println("")
	fmt.Println(string(bytes))
	filePath := "./json.json"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(string(bytes))
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

// parseTable parse table name and comment for ddl string
func parseTable(ddl string) ([]string, error) {
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
		if results := tableCommentRegexp.FindStringSubmatch(ddl); len(results) > 0 {
			// value: '文章'
			oriTableCommentStr := results[1]
			tableCommentStr = normalizedName(oriTableCommentStr)
		} else {
			tableCommentStr = tableNameStr
		}
		return []string{
			tableNameStr,    // article
			tableCommentStr, // 文章
		}, nil
	}
	return nil, errors.New("parse ddl error, cannot found a valid table")
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
	for index, match := range matches {
		javaType := ""
		javaImport := ""
		goType := ""
		phpType := ""
		fmt.Println(index)
		fmt.Println("-------")
		fmt.Println(match)
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
			goType, _ = MapToGoType(additionalAttr)
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
				// PythonType:       "",
				GoType: goType,
				GoTag:  goTag,
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

// getTableFieldsStr get table fields str cut off ( and )
func getTableFieldsStr(ddl string) (string, error) {
	openedBracketIndex := strings.IndexAny(ddl, `(`)
	closedBracketIndex := strings.LastIndexAny(ddl, `)`)
	if openedBracketIndex != -1 && closedBracketIndex != -1 {
		return ddl[(openedBracketIndex + 1):closedBracketIndex], nil
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
