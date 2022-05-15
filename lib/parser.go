package lib

import (
	"errors"
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
	Columns              []Column
}

// Column sql field
type Column struct {
	Name             string
	PascalName       string
	CamelName        string
	Comment          string
	CommentTags      map[string]string
	OriginalFieldStr string
	SqlType          string
	SqlProps         string
	JavaType         string
	JavaImport       string
	PhpType          string
	PythonType       string
	GoType           string
	GoTag            string
	GoImport         string
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
var example = "CREATE TABLE `s_blog`.`article` (\n" +
	"  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',\n" +
	"  `user_id` bigint NOT NULL COMMENT '用户id',\n" +
	"  `content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '正文',\n" +
	"  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n" +
	"  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n" +
	"  PRIMARY KEY (`id`)\n" +
	") ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT '文章';"

const (
	// TableNameRegex \x60 for `
	TableNameRegex = `(?im)CREATE\s+TABLE\s+([\x60-zA-Z-_."']+)`
	TableCommentRegex = `(?im)\).*COMMENT\s+(?="|')(.*)(?<="|')`
	// FieldsRegex \x60 for `
	FieldsRegex = `(?im)([\w\x60"']+)\s+([\w\(\),]+).*(\s+COMMENT\s+.*)?`
	// ColumnCommentRegex parse column comment
	ColumnCommentRegex = `(?im)COMMENT\s+(?="|')(.*)(?<="|')`
)

// Parse parse MySQL DDL

func Parse(ddl string) {

}

// parseTable parse table name and comment for ddl string
func parseTable(ddl string) ([]string, error){
	tableRegexp := regexp.MustCompile(TableNameRegex)
	if sections := tableRegexp.FindStringSubmatch(ddl); len(sections) > 0 {
		// `s_blog`.`article`
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
			// '文章'
			oriTableCommentStr := results[1]
			tableCommentStr = normalizedName(oriTableCommentStr)
		} else {
			tableCommentStr = tableNameStr
		}
		return []string {
			tableNameStr,  // article
			tableCommentStr,  // 文章
		}, nil
	}
	return nil, errors.New("parse ddl error, cannot found a valid table")
}

// parseFields parse table fields
// you can test regex online by https://regex101.com/
func parseFields(fieldsStr string) {
	fieldsRegexp := regexp.MustCompile(FieldsRegex)
	matches := fieldsRegexp.FindAllStringSubmatch(fieldsStr, -1)

	len := len(matches)
	/*
	[`id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id', `id` bigint  ]  found at index 0
	[`user_id` bigint NOT NULL COMMENT '用户id', `user_id` bigint  ]  found at index 1
	[`content` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '正文', `content` text  ]  found at index 2
	[`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间', `create_time` datetime  ]  found at index 3
	[`update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间', `update_time` datetime  ]  found at index 4
	[PRIMARY KEY (`id`) PRIMARY KEY  ]
	*/
	for _, match := range matches {
		// `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
		oriFieldStr := match[0]
		// `id`
		oriField := match[1]

		// bigint
		oriFieldType := match[2]
		// BIGINT
		upperFieldType := strings.ToUpper(oriFieldType)
		// unsigned?
		isUnsigned := isUnsigned(oriFieldStr)

	}
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
		return ddl[(openedBracketIndex+1):closedBracketIndex], nil
	}
	return "", errors.New("not a valid ddl string")
}

// isUnsigned is unsigned?
func isUnsigned(field string) bool {
	upperFieldStr := strings.ToUpper(field)
	return strings.Contains(upperFieldStr, "UNSIGNED")
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