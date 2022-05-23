package lib

import "testing"

var exampleTableName = "article"
var exampleTableComment = "文章"
var exampleColumns = []string{
	"id",
	"user_id",
	"content",
	"create_time",
	"update_time",
}

var exampleJavaTypeMaps = map[string]string{
	"id":         "Long",
	"userId":     "Long",
	"content":    "String",
	"createTime": "Timestamp",
	"updateTime": "Timestamp",
}

var exampleGoTypeMaps = map[string]string{
	"Id":         "uint64",
	"UserId":     "int64",
	"Content":    "string",
	"CreateTime": "time.Time",
	"UpdateTime": "time.Time",
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

func TestExample(t *testing.T) {
	result, err := Parse(Example)
	if err != nil {
		panic(err)
	}

	if exampleTableName != result.TableName {
		t.Errorf("table name should be %s, but got %s", exampleTableName, result.TableName)
	}

	if exampleTableComment != result.TableComment {
		t.Errorf("table comment should be %s, but got %s", exampleTableComment, result.TableComment)
	}

	for _, column := range result.Columns {
		goType := column.GoType
		goProperty := column.PascalName
		if exampleGoTypeMaps[goProperty] != goType {
			t.Errorf("data type in golang for %s property should be %s, but got %s", goProperty, exampleGoTypeMaps[goProperty], goType)
		}
		javaType := column.JavaType
		javaProperty := column.CamelName
		if exampleJavaTypeMaps[javaProperty] != javaType {
			t.Errorf("data type in java for %s property should be %s, but got %s", javaProperty, exampleJavaTypeMaps[javaProperty], javaType)
		}
	}
}
