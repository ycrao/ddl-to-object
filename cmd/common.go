package cmd

import (
	"ddl-to-object/lib"
	"fmt"
	"os"
	"strings"
)

// CommandConfig holds configuration for language-specific commands
type CommandConfig struct {
	Language       string
	DefaultPackage string
	FileExtension  string
	PackageKey     string
}

// ExecuteLanguageCommand is a common handler for all language commands
func ExecuteLanguageCommand(config CommandConfig) error {
	// Load configuration
	appConfig, err := lib.LoadConfig(configFile)
	if err != nil {
		lib.Warnf("Failed to load config: %v, using defaults", err)
		appConfig = lib.DefaultConfig()
	}

	// Set log level based on verbose flag or config
	if verbose {
		lib.SetLogLevel(lib.DEBUG)
	} else {
		switch appConfig.LogLevel {
		case "debug":
			lib.SetLogLevel(lib.DEBUG)
		case "info":
			lib.SetLogLevel(lib.INFO)
		case "warn":
			lib.SetLogLevel(lib.WARN)
		case "error":
			lib.SetLogLevel(lib.ERROR)
		}
	}

	lib.Debugf("Starting %s code generation", config.Language)

	// Validate inputs
	if err := lib.ValidateInputFile(from); err != nil {
		return fmt.Errorf("input validation failed: %w", err)
	}

	var packageName string
	switch config.Language {
	case "php":
		packageName = ns
		if packageName == "" {
			packageName = appConfig.DefaultPackages["php"]
		}
	default:
		packageName = pk
		if packageName == "" {
			packageName = appConfig.DefaultPackages[config.Language]
		}
	}

	if err := lib.ValidatePackageName(packageName, config.Language); err != nil {
		return fmt.Errorf("package validation failed: %w", err)
	}

	if !stdout && !dryRun {
		if err := lib.ValidateOutputDir(to); err != nil {
			return fmt.Errorf("output directory validation failed: %w", err)
		}
	}

	// Prepare output directory
	toDir := "./"
	if to != "" {
		toDir = strings.TrimRight(to, "/") + "/"
	}

	lib.Infof("Reading DDL file: %s", from)

	// Read and parse DDL file
	content, err := lib.ReadFile(from)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	lib.Debugf("Parsing DDL content (%d bytes)", len(content))

	result, err := lib.Parse(string(content))
	if err != nil {
		return fmt.Errorf("failed to parse DDL: %w", err)
	}

	lib.Infof("Parsed table: %s with %d columns", result.TableName, len(result.Columns))

	// Set language-specific package name
	switch config.Language {
	case "go":
		if packageName != "" {
			packageArr := strings.Split(packageName, ".")
			if len := len(packageArr); len > 0 {
				result.GoPackageName = packageArr[len-1]
			}
		}
	case "java":
		result.JavaPackageName = packageName
	case "php":
		result.PhpNamespaceName = packageName
	case "python":
		// Python doesn't need package name modification
	}

	// Generate target filename
	var targetFile string
	switch config.Language {
	case "go":
		targetFile = toDir + result.SnakeObjectName + "_types" + config.FileExtension
	case "java":
		targetFile = toDir + result.ObjectName + config.FileExtension
	default:
		targetFile = toDir + result.SnakeObjectName + config.FileExtension
	}

	lib.Debugf("Target file: %s", targetFile)

	// Load and execute template
	tpl, err := lib.ReadTemplate(config.Language)
	if err != nil {
		return fmt.Errorf("failed to load template: %w", err)
	}

	if dryRun {
		fmt.Printf("=== DRY RUN MODE ===\n")
		fmt.Printf("Would generate %s file: %s\n", config.Language, targetFile)
		fmt.Printf("Package/Namespace: %s\n", packageName)
		fmt.Printf("Table: %s -> Object: %s\n", result.TableName, result.ObjectName)
		fmt.Printf("Columns (%d):\n", len(result.Columns))
		for _, col := range result.Columns {
			fmt.Printf("  - %s (%s) -> %s\n", col.Name, col.DataType, getColumnType(col, config.Language))
		}
		fmt.Printf("=== END DRY RUN ===\n")
		return nil
	}

	if stdout {
		lib.Debug("Outputting to stdout")
		if err := tpl.Execute(os.Stdout, result); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
	} else {
		// Check if file exists and backup if configured
		if appConfig.OutputSettings.BackupExisting {
			if _, err := os.Stat(targetFile); err == nil {
				backupFile := targetFile + ".bak"
				lib.Infof("Backing up existing file to: %s", backupFile)
				if err := lib.CopyFile(targetFile, backupFile); err != nil {
					lib.Warnf("Failed to backup file: %v", err)
				}
			}
		}

		file, err := os.Create(targetFile)
		if err != nil {
			// Fallback to stdout
			lib.Warnf("Failed to create output file, using stdout: %v", err)
			if err := tpl.Execute(os.Stdout, result); err != nil {
				return fmt.Errorf("failed to execute template: %w", err)
			}
			return nil
		}
		defer file.Close()

		if err := tpl.Execute(file, result); err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
		lib.Infof("Generated %s file: %s", config.Language, targetFile)
	}

	return nil
}

// getColumnType returns the appropriate type for a column in the given language
func getColumnType(col lib.Column, language string) string {
	switch language {
	case "go":
		return col.GoType
	case "java":
		return col.JavaType
	case "php":
		return col.PhpType
	default:
		return col.DataType
	}
}
