package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/alexbegoon/go-dcc/internal/module"

	"github.com/alexbegoon/go-dcc/internal/locomotive"
)

// Config allows to store configuration settings to initialize go-dcc.
type Config struct {
	Locomotives    []*locomotive.Locomotive `json:"locomotives"`
	RailwayModules []*module.Railway        `json:"railway_modules"`
}

// LoadConfig parses a configuration file and returns a Config object.
func LoadConfig(path string) (*Config, error) {
	conf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = json.Unmarshal(conf, &cfg)
	if err != nil {
		return nil, err
	}
	log.Printf("Loaded configuration for %d locomotive(s)", len(cfg.Locomotives))
	log.Printf("Loaded configuration for %d railway module(s)", len(cfg.RailwayModules))

	return &cfg, nil
}

func (c *Config) Save(path string) error {
	pretty, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, pretty, 0o600) // nolint:gomnd

	return err
}
