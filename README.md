# ddl-to-object

>   ## Notice: STATUS  WIP


>   ddl-to-object: a tool help to generate object files in different languages from sql ddl file.

### database supports

- only tested for MySQL/MariaDB DDL SQL

### language supports

>   You can support another language by adding a new template.

- java: generate entity class with snake_style to camelStyle, comments, and package directory support
- golang: generate to struct with tag and comments
- php: generate to simple class with namespace and comments support
- python: generate to simple object with comments support

### best practice

- A good-designed pattern in MySQL DDL, such as using singular nouns as table and column name, naming in `snake_case` style, with more comments, no table prefix, having a primary key etc.
- The rest, just using this tool to help you generate target language object files

#### command helps

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

#### usage example

```bash
$ ddl-to-object php -f ./output/samples/example_3.ddl.txt -n Modules\\Blog\\Models -t ./output/php/
$ ddl-to-object java -f ./output/samples/example_2.ddl.txt -p com.douyasi.sample.domain.entity -t ./output/java/
$ ddl-to-object go -f ./output/sampls/example_3.ddl.txt -p models -t ./output/go/
```

### known so-called bugs

- not work well in one-line DDL SQL
- not work well with mixed style (such as snake_style, camelStyle, PascalStyle and other cases mixed) in DDL SQL
- special table or field name, such as `365Days_table`, `1st_field` and `biz.error.code.field` etc 
- some MySQL data type may not mapper well to Java or Golang data type
- lack of some special testing cases

### similar projects and references

- [liangyaopei/sqltogo](https://github.com/liangyaopei/sqltogo)
- [xwb1989/sqlparser](https://github.com/xwb1989/sqlparser)
- [nao1215/ddl-maker](https://github.com/nao1215/ddl-maker)
- [zeromicro/ddl-parser](https://github.com/zeromicro/ddl-parser)
- [blastrain/vitess-sqlparser](https://github.com/blastrain/vitess-sqlparser)
- [json-to-go](https://mholt.github.io/json-to-go/)
- [curl-to-go](https://mholt.github.io/curl-to-go/)
- [sql-to-go-struct-java-class-json-format](https://plugins.jetbrains.com/plugin/17336-sql-to-go-struct-java-class-json-format)
