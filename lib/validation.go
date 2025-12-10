package lib

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// ValidateInputFile validates the input DDL file
func ValidateInputFile(filePath string) error {
	if filePath == "" {
		return errors.New("input file path cannot be empty")
	}

	if !strings.HasSuffix(strings.ToLower(filePath), ".sql") &&
		!strings.HasSuffix(strings.ToLower(filePath), ".ddl") &&
		!strings.HasSuffix(strings.ToLower(filePath), ".txt") {
		return errors.New("input file should have .sql, .ddl, or .txt extension")
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("input file does not exist: " + filePath)
	}

	return nil
}

// ValidateOutputDir validates and creates the output directory
func ValidateOutputDir(outputDir string) error {
	if outputDir == "" {
		return nil // Use current directory
	}

	// Clean the path
	outputDir = filepath.Clean(outputDir)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return errors.New("failed to create output directory: " + err.Error())
	}

	return nil
}

// ValidatePackageName validates package/namespace names
func ValidatePackageName(packageName, language string) error {
	if packageName == "" {
		return nil // Use default
	}

	switch language {
	case "go":
		if strings.Contains(packageName, " ") {
			return errors.New("Go package name cannot contain spaces")
		}
	case "java":
		if !strings.Contains(packageName, ".") {
			return errors.New("Java package name should follow dot notation (e.g., com.example.domain)")
		}
	case "php":
		if !strings.Contains(packageName, "\\") {
			return errors.New("PHP namespace should use backslash notation (e.g., App\\Models)")
		}
	}

	return nil
}
