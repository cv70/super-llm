package config

import (
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	LLMs []*LLMConfig `yaml:"llms"`
	LeaderName string `yaml:"leader_name"`
}

type LLMConfig struct {
	BaseURL          string        `yaml:"base_url"`
	Model            string        `yaml:"model"`
	APIKey           string        `yaml:"api_key"`
	MaxTokens        *int          `yaml:"max_tokens,omitempty"`
	Temperature      *float32      `yaml:"temperature,omitempty"`
	TopP             *float32      `yaml:"top_p,omitempty"`
	PresencePenalty  *float32      `yaml:"presence_penalty,omitempty"`
	FrequencyPenalty *float32      `yaml:"frequency_penalty,omitempty"`
}

func LoadConfig() (*Config, error) {
	var once sync.Once
	var cfg *Config
	var err error

	once.Do(func() {
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			configPath = "config/config.yaml"
		}

		data, readErr := os.ReadFile(configPath)
		if readErr != nil {
			err = fmt.Errorf("failed to read config file: %w", readErr)
			return
		}

		unmarshalErr := yaml.Unmarshal(data, &cfg)
		if unmarshalErr != nil {
			err = fmt.Errorf("failed to unmarshal config: %w", unmarshalErr)
			return
		}
	})

	return cfg, err
}

var (
	globalConfig *Config
	configOnce   sync.Once
)

// GetConfig returns the global config instance
func GetConfig() *Config {
	configOnce.Do(func() {
		cfg, err := LoadConfig()
		if err != nil {
			panic(fmt.Sprintf("failed to load config: %v", err))
		}
		globalConfig = cfg
	})
	return globalConfig
}
