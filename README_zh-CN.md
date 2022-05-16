# ddl-to-object

[![GoVersion](https://img.shields.io/github/go-mod/go-version/ycrao/ddl-to-object)](https://github.com/ycrao/ddl-to-object/blob/master/go.mod)
[![Release](https://img.shields.io/github/v/release/ycrao/ddl-to-object)](https://github.com/ycrao/ddl-to-object/releases)
![Stars](https://img.shields.io/github/stars/ycrao/ddl-to-object)
[![MIT license](https://img.shields.io/github/license/ycrao/ddl-to-object)](https://opensource.org/licenses/MIT)
[![OpenIssue](https://img.shields.io/github/issues/ycrao/ddl-to-object)](https://github.com/ycrao/ddl-to-object/issues?q=is%3Aopen+is%3Aissue)

[ENGLISH README/英文读我](./README.md)

>   ddl-to-object: 一个工具，帮助从 `SQL DDL` 文件生成不同语言的相关类文件。

### 数据库支持

- 仅针对 MySQL/MariaDB DDL SQL 进行了测试。

### 语言支持

>   欢迎 PR ！您可以做一些编码工作来让它支持另一种（编程）语言。

- java: 生成实体类，对属性自动转换（数据库中蛇形 `snake_style` 字段）到驼峰 `camelStyle` 风格，引入注释，并使用 lombok 插件减少 `getter/setter` 等代码、并支持包目录；
- golang: 生成带有标签 `tag` 和注释 `comment` 的结构体；
- php: 生成带有命名空间和注释支持的简单类；
- python: 生成带有注释的简单类；
- 想要其它编程语言支持？：`Pull Request` 是欢迎的。

### 最佳实践

- 在 MySQL DDL 中使用良好的涉及模式，例如使用单数名词作为表名和列名，以蛇形 `snake_case` 样式命名、更多的注释、不要使用表前缀和表有主键等。
- 剩下地，就是使用这个工具帮你生成目标语言类文件。

#### 安装

下载目标操作系统 ZIP 压缩文件，解压它，把二进制文件（即 `ddl-to-object` 或 `ddl-to-object.exe`）移动到 `/usr/local/bin/` 或其它可自动加载的环境变量路径下。

默认情况下，您需要复制项目 `template` 文件到 `~/.dto/template` 目录下（注：`~` 代表当前用户家目录）。

然后您即可从任何位置在终端 `terminal` 运行它。从它打印帮助或下文中获取帮助。

#### 命令帮助

```bash
  ddl-to-object go         generate golang target object file
  ddl-to-object java       generate java target object file
  ddl-to-object php        generate php target object file
  ddl-to-object python     generate python target object file
  -f, --from path      from path which a single-table DDL file located
  -n, --ns namespace   namespace name for php, only in php command (default "App\\Models")
  -p, --pk package     package name, only in java or go command (default "com.example.sample.domain.entity")
  -s, --stdout         enable stdout or not, default set false to disable
  -t, --to path        output to target path or location, create directory automatically if it not existed
```

#### 使用示例

```bash
$ ddl-to-object php -f ./output/samples/example_3.ddl.txt -n Modules\\Blog\\Models -t ./output/php/
$ ddl-to-object java -f ./output/samples/example_2.ddl.txt -p com.douyasi.sample.domain.entity -t ./output/java/
$ ddl-to-object go -f ./output/sampls/example_3.ddl.txt -p models -t ./output/go/
```

#### 输出示例

查看 output 文件夹。

- [java](./output/java/Article.java)
- [golang](./output/go/article_types.go)
- [php](./output/php/Article.php)
- [python](./output/python/article.py)

#### 如何修改模板

如安装那节所说，默认模板 `template` 文件位于 `~/.dto/template` 目录下（注：代表当前用户家目录；如果它们不存在，您需要手动复制它们到此位置）。

模板是个纯文本，使用到 golang [text/template](https://pkg.go.dev/text/template) ， `ParsedResult` 类型的结构体会被传入。你可以根据自己的能力来修改它们。

### 已知所谓的缺陷

- 在单行 DDL SQL 中不能正常工作；
- 在 DDL SQL 中的使用混杂（命名）风格（如蛇形 `snake_style`、驼峰 `camelStyle`、帕斯卡式 `PascalStyle` 和其他情况）的，不能很好地工作；
- 特殊的表或字段名，如 `365Days_table` 、`1st_field` 和 `biz.error.code.field` 等；
- 某些 MySQL 数据类型可能无法很好地映射成（合适的） Java 或 Golang 数据类型；
- Java 和 Golang 中会有一些未使用的引用（包名），您可以自行清理它们或使用诸如 `gofmt` 之类的工具；
- 缺少一些特殊测试用例。

### 类似项目及参考

- [liangyaopei/sqltogo](https://github.com/liangyaopei/sqltogo)
- [xwb1989/sqlparser](https://github.com/xwb1989/sqlparser)
- [nao1215/ddl-maker](https://github.com/nao1215/ddl-maker)
- [zeromicro/ddl-parser](https://github.com/zeromicro/ddl-parser)
- [blastrain/vitess-sqlparser](https://github.com/blastrain/vitess-sqlparser)
- [json-to-go](https://mholt.github.io/json-to-go/)
- [curl-to-go](https://mholt.github.io/curl-to-go/)
- [sql-to-go-struct-java-class-json-format](https://plugins.jetbrains.com/plugin/17336-sql-to-go-struct-java-class-json-format)
