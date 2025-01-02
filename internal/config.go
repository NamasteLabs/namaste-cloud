package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config structure to store global configuration like active cloud provider.
type Config struct {
	ActiveCloud string `json:"active_cloud"`
}

// GetConfigFilePath returns the path to the configuration file.
func GetConfigFilePath() (string, error) {
	configDir, err := GetUserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

// GetUserConfigDir returns the user-specific configuration directory.
func GetUserConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}
	return filepath.Join(homeDir, ".namaste-cloud"), nil
}

// EnsureConfigDir ensures the configuration directory exists.
func EnsureConfigDir() error {
	configDir, err := GetUserConfigDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(configDir, 0700) // Ensure the directory is private
}

// SaveActiveCloudProvider saves the active cloud provider to the configuration file.
func SaveActiveCloudProvider(activeCloud string) error {
	// Ensure the configuration directory exists
	if err := EnsureConfigDir(); err != nil {
		return err
	}

	// Load existing configuration or initialize a new one
	cfg, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load existing config: %w", err)
	}

	// Update the active cloud provider
	cfg.ActiveCloud = activeCloud

	// Save the updated configuration
	return SaveConfig(cfg)
}

// LoadActiveCloudProvider loads the active cloud provider from the configuration file.
func LoadActiveCloudProvider() (string, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return "", err
	}

	if cfg.ActiveCloud == "" {
		return "", fmt.Errorf("No active cloud provider set. Please use `namaste-cloud use-cloud` to select one.")
	}

	return cfg.ActiveCloud, nil
}

// SaveConfig saves the entire configuration to a file.
func SaveConfig(cfg Config) error {
	if err := EnsureConfigDir(); err != nil {
		return err
	}

	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("failed to write config to file: %w", err)
	}

	return nil
}

// LoadConfig loads the configuration from a file.
func LoadConfig() (Config, error) {
	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Return default configuration if the file doesn't exist
			return Config{}, nil
		}
		return Config{}, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to read config from file: %w", err)
	}

	return cfg, nil
}
