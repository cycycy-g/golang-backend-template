package config

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration values
type Config struct {
	// Database settings
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`

	// Server settings
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	Environment   string `mapstructure:"ENVIRONMENT"`

	// Authentication
	JWTSecret   string        `mapstructure:"JWT_SECRET"`
	JWTDuration time.Duration `mapstructure:"JWT_DURATION"`

	// CORS
	AllowOrigins []string `mapstructure:"ALLOW_ORIGINS"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// Enable reading from environment variables
	// Environment variables take precedence over config file
	viper.AutomaticEnv()

	// Attempt to read config file
	if err = viper.ReadInConfig(); err != nil {
		// It's okay if config file doesn't exist
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return
		}
	}

	// Set defaults
	setDefaults()

	// Unmarshal config into struct
	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	// Convert JWT duration from hours to time.Duration
	config.JWTDuration = time.Duration(viper.GetDuration("JWT_DURATION")) * time.Hour
	if config.JWTDuration == 0 {
		return config, fmt.Errorf("invalid or missing JWT_DURATION in config file")
	}

	return
}

// setDefaults sets default values for configuration
func setDefaults() {
	viper.SetDefault("ENVIRONMENT", "development")
	viper.SetDefault("SERVER_ADDRESS", ":8080")
	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("JWT_DURATION", "1h") // 24 hours
}
