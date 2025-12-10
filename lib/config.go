package lib

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds application configuration
type Config struct {
	DefaultPackages map[string]string `json:"default_packages"`
	TemplateDir     string            `json:"template_dir"`
	LogLevel        string            `json:"log_level"`
	OutputSettings  OutputSettings    `json:"output_settings"`
}

// OutputSettings holds output-related configuration
type OutputSettings struct {
	CreateDirectories bool `json:"create_directories"`
	OverwriteFiles    bool `json:"overwrite_files"`
	BackupExisting    bool `json:"backup_existing"`
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		DefaultPackages: map[string]string{
			"go":     "main",
			"java":   "com.example.sample.domain.entity",
			"php":    "App\\Models",
			"python": "",
		},
		TemplateDir: "~/.dto/template",
		LogLevel:    "info",
		OutputSettings: OutputSettings{
			CreateDirectories: true,
			OverwriteFiles:    true,
			BackupExisting:    false,
		},
	}
}

// LoadConfig loads configuration from file or returns default
func LoadConfig(configPath string) (*Config, error) {
	config := DefaultConfig()

	if configPath == "" {
		// Try default locations
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return config, nil // Use default if can't get home dir
		}
		configPath = filepath.Join(homeDir, ".dto", "config.json")
	}

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return config, nil // Use default if file doesn't exist
	}

	// Read and parse config file
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(data, config); err != nil {
		return config, err
	}

	return config, nil
}

// SaveConfig saves configuration to file
func SaveConfig(config *Config, configPath string) error {
	if configPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configDir := filepath.Join(homeDir, ".dto")
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return err
		}
		configPath = filepath.Join(configDir, "config.json")
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}
