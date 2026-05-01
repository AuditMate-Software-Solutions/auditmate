package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OutputDir string        `yaml:"output_dir"`
	LogLevel  string        `yaml:"log_level"`
	Timeout   time.Duration `yaml:"timeout"`
	Reset     bool          `yaml:"reset"`
}

func Default() Config {
	return Config{
		OutputDir: "auditmate-output",
		LogLevel:  "info",
		Timeout:   10 * time.Second,
		Reset:     false,
	}
}

func Load(path string) Config {
	cfg := Default()

	data, err := os.ReadFile(path)
	if err == nil {
		_ = yaml.Unmarshal(data, &cfg)
	}

	// ENV overrides (production standard)
	if v := os.Getenv("AUDITMATE_OUTPUT_DIR"); v != "" {
		cfg.OutputDir = v
	}

	if v := os.Getenv("AUDITMATE_LOG_LEVEL"); v != "" {
		cfg.LogLevel = v
	}

	return cfg
}
