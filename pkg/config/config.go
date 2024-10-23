package config

import (
	"fmt"
	"log"
	"github.com/spf13/viper"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	GeminiKey	string `mapstructure:"GEMINI_KEY"`
	ChatGPTKey  string `mapstructure:"CHATGPT_KEY"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	// Explicitly bind environment variables to config keys
	keys := []string{
		"PORT",
		"DATABASE_URL",
		"GEMINI_KEY",
		"CHATGPT_KEY",
	}

	for _, key := range keys {
		if err := viper.BindEnv(key); err != nil {
			return nil, fmt.Errorf("failed to bind environment variable %s: %w", key, err)
		}
	}

	// Optionally read from a config file if specified by an environment variable
	configFile := viper.GetString("CONFIG_FILE")
	if configFile != "" {
		viper.SetConfigFile(configFile)
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
		log.Println("Using config file:", configFile)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Load from .env local
	if cfg == (Config{}) {
		viper.AutomaticEnv()
		viper.SetConfigFile(".env")

		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			return nil, err
		}
	}

	// Debug: Print the loaded configuration
	log.Printf("Loaded Config: %+v\n", cfg)

	// Validate required config values
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func validateConfig(cfg Config) error {
	missingEnvVars := []string{}

	if cfg.Port == "" {
		missingEnvVars = append(missingEnvVars, "KEYCLOAK_CLIENT_SECRET")
	}
	if cfg.DatabaseURL == "" {
		missingEnvVars = append(missingEnvVars, "DATABASE_URL")
	}
	// Add more validation checks as necessary

	if len(missingEnvVars) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingEnvVars)
	}

	return nil
}
