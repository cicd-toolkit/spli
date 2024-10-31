package configmanager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
)

// Config struct to handle configuration operations
type Config struct {
	filePath string
	profile  string
	cfg      *ini.File
}

// DefaultConfigPath returns the default configuration file path
func DefaultConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("failed to get user home directory: %w", err))
	}
	return filepath.Join(homeDir, ".spli", "config")
}

// NewConfig function to initialize a new Config instance
func NewConfig() (*Config, error) {
	filePath := DefaultConfigPath()

	// Ensure the directory exists
	configDir := filepath.Dir(filePath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0700)
		if err != nil {
			return nil, fmt.Errorf("failed to create config directory: %w", err)
		}
	}

	cfg, err := ini.Load(filePath)
	if err != nil {
		// If the file does not exist, create a new empty ini file
		cfg = ini.Empty()
		err = cfg.SaveTo(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create new config file: %w", err)
		}
	}
	return &Config{filePath: filePath, cfg: cfg, profile: "default"}, nil
}

// GetString retrieves a string value given a section and key
func (c *Config) GetValue(key string) string {
	return c.cfg.Section(c.profile).Key(key).String()
}

func (c *Config) GetString(section, key string) string {
	env_key := fmt.Sprintf("SPLUNK_%s", strings.ToUpper(key))
	if value, exists := os.LookupEnv(env_key); exists {
		return value
	}
	return c.cfg.Section(section).Key(key).String()
}

// GetInt retrieves an integer value given a section and key
func (c *Config) GetInt(section, key string) (int, error) {
	return c.cfg.Section(section).Key(key).Int()
}

func (c *Config) SetValue(key, value string) error {
	c.cfg.Section(c.profile).Key(key).SetValue(value)
	return c.cfg.SaveTo(c.filePath)
}

// SetString sets a string value for a given section and key
func (c *Config) SetString(section, key, value string) error {
	c.cfg.Section(section).Key(key).SetValue(value)
	return c.cfg.SaveTo(c.filePath)
}

// SetInt sets an integer value for a given section and key
func (c *Config) SetInt(section, key string, value int) error {
	c.cfg.Section(section).Key(key).SetValue(fmt.Sprintf("%d", value))
	return c.cfg.SaveTo(c.filePath)
}

func (c *Config) Sections() []string {
	return c.cfg.SectionStrings()
}

func (c *Config) DeleteKey(section, key string) error {
	c.cfg.Section(section).DeleteKey(key)
	return c.cfg.SaveTo(c.filePath)
}

// DeleteSection deletes a section
func (c *Config) DeleteSection(section string) error {
	c.cfg.DeleteSection(section)
	return c.cfg.SaveTo(c.filePath)
}
