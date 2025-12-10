# ddl-to-object

[![GoVersion](https://img.shields.io/github/go-mod/go-version/ycrao/ddl-to-object)](https://github.com/ycrao/ddl-to-object/blob/master/go.mod)
[![Release](https://img.shields.io/github/v/release/ycrao/ddl-to-object)](https://github.com/ycrao/ddl-to-object/releases)
![Stars](https://img.shields.io/github/stars/ycrao/ddl-to-object)
[![MIT license](https://img.shields.io/github/license/ycrao/ddl-to-object)](https://opensource.org/licenses/MIT)
[![OpenIssue](https://img.shields.io/github/issues/ycrao/ddl-to-object)](https://github.com/ycrao/ddl-to-object/issues?q=is%3Aopen+is%3Aissue)

[SimplifiedChinese README/简体中文读我](./README_zh-CN.md)

ddl-to-object: a tool help to generate object files in different languages from sql ddl file.

## Database supports

- only tested for MySQL/MariaDB DDL SQL

## Language supports

PR is welcome! You can do some coding stuff to support another language.

- java: generate entity class with auto snake_style to camelStyle naming in properties, bring comments, using lombok plugin for getter/setter, with package directory support
- golang: generate to struct with tags and comments
- php: generate to simple class with namespace and comments support
- python: generate to simple object with comments support
- support any other program language?: pull request is welcome

## Best practice

- A good-designed pattern in MySQL DDL, such as using singular nouns as table and column name, naming in `snake_case` style, with more comments, no table prefix, having a primary key etc.
- The rest, just using this tool to help you generate target language object files

## Installation

Download targeted OS zip file, unzip it, and move the binary file (`ddl-to-object` or `ddl-to-object.exe`) to `/usr/local/bin/` or other auto-load environment path.

By default, you need copy this project template files into to `~/.dto/template` directory manually (note: `~` for current user home workdir).

Then you run it in terminal from anywhere. Get helps from its help print or below.

## Command helps

```bash
  ddl-to-object go         Generate golang target object file
  ddl-to-object java       Generate java target object file
  ddl-to-object php        Generate php target object file
  ddl-to-object python     Generate python target object file
  -c, --config string   config file path (default: ~/.dto/config.json)
  -f, --from path       from path which a single-table DDL file located
  -h, --help            help for ddl-to-object
  -n, --ns namespace    namespace name for php, only in php command (default: App\Models)
  -p, --pk package      package name, only in java or go command (default: com.example.sample.domain.entity)
  -s, --stdout          enable stdout or not, default set false to disable
  -t, --to path         output to target path or location, create directory automatically if it not existed
  -v, --verbose         enable verbose output
      --dry-run         show what would be generated without creating files
```

## Usage examples

```bash
ddl-to-object php -f ./output/samples/example_3.ddl.txt -n Modules\\Blog\\Models -t ./output/php/
ddl-to-object java -f ./output/samples/example_2.ddl.txt -p com.douyasi.sample.domain.entity -t ./output/java/
ddl-to-object go -f ./output/samples/example_3.ddl.txt -p models -t ./output/go/
```

## Output examples

See output directory.

- [java](./output/java/Article.java)
- [golang](./output/go/article_types.go)
- [php](./output/php/Article.php)
- [python](./output/python/article.py)

## How to modify templates

As installation intro, default template files located in `~/.dto/template` directory (note: `~` for current user home workdir; if they're not existed, you need copy them by yourself manually).

The template is a raw text by using golang [text/template](https://pkg.go.dev/text/template) with `ParsedResult` type struct passed in. You can modify them as you can.

## Known so-called bugs

- not work well in one-line DDL SQL
- not work well with mixed style (such as snake_style, camelStyle, PascalStyle and other cases mixed) in DDL SQL
- special table or field name, such as `365Days_table`, `1st_field` and `biz.error.code.field` etc 
- some MySQL data type may not mapper well to Java or Golang data type
- some unused imports in Java and Golang, you can clean them by yourself or using some tool like `gofmt`
- lack of some special testing cases

## Configuration

You can create a configuration file to customize default settings. Configuration file location: `~/.dto/config.json`

Example configuration:

```json
{
  "default_packages": {
    "go": "models",
    "java": "com.yourcompany.domain.entity",
    "php": "App\\Models",
    "python": ""
  },
  "template_dir": "~/.dto/template",
  "log_level": "info",
  "output_settings": {
    "create_directories": true,
    "overwrite_files": true,
    "backup_existing": false
  }
}
```

## New Features

- **Improved error handling**: Better error messages and validation
- **Configuration file support**: Customize default packages and settings
- **Verbose mode**: Use `-v` flag for detailed output
- **Dry-run mode**: Use `--dry-run` to preview what would be generated
- **Backup functionality**: Optional backup of existing files
- **Better logging**: Leveled log output
- **Version information**: Detailed build information

## Building

Use the provided Makefile:

```bash
# Build for single platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Format code
make fmt
```

## Similar projects and references

- [liangyaopei/sqltogo](https://github.com/liangyaopei/sqltogo)
- [xwb1989/sqlparser](https://github.com/xwb1989/sqlparser)
- [nao1215/ddl-maker](https://github.com/nao1215/ddl-maker)
- [zeromicro/ddl-parser](https://github.com/zeromicro/ddl-parser)
- [blastrain/vitess-sqlparser](https://github.com/blastrain/vitess-sqlparser)
- [json-to-go](https://mholt.github.io/json-to-go/)
- [curl-to-go](https://mholt.github.io/curl-to-go/)
- [sql-to-go-struct-java-class-json-format](https://plugins.jetbrains.com/plugin/17336-sql-to-go-struct-java-class-json-format)
