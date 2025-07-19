/*
Copyright © 2025 SHAKTHI WARNAKULASURIYA <shakthiraveen2002@gmail.com>
*/
package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	OpenAIKey       string `mapstructure:"openai_key"`
	GeminiKey       string `mapstructure:"gemini_key"`
	ClaudeKey       string `mapstructure:"claude_key"`
	DefaultProvider string `mapstructure:"default_provider"`
}

func SaveConfig(openai, gemini, claude, defaultProvider string) error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	appConfigPath := filepath.Join(configDir, "chatcli")
	err = os.MkdirAll(appConfigPath, 0700)
	if err != nil {
		log.Printf("Error creating config directory: %v\n", err)
		return err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(appConfigPath)

	viper.Set("openai_key", openai)
	viper.Set("gemini_key", gemini)
	viper.Set("claude_key", claude)
	viper.Set("default_provider", defaultProvider)

	err = viper.WriteConfigAs(filepath.Join(appConfigPath, "config.yaml"))
	if err != nil {
		log.Printf("Error writing config: %v\n", err)
	}
	log.Printf("Configuration file path: %s\n", filepath.Join(appConfigPath, "config.yaml"))
	return err
}

func LoadConfig() (*Config, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	appConfigPath := filepath.Join(configDir, "chatcli")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(appConfigPath)

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
