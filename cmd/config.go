package main

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/hcl"
	"github.com/knadh/koanf/providers/file"
)

// initConfig parses the config file and loads in `Config` object.
func initConfig(ko *koanf.Koanf, cfgPath string) (Config, error) {
	var (
		cfg Config
	)

	// Load the config files from the path provided.
	err := ko.Load(file.Provider(cfgPath), hcl.Parser(false))
	if err != nil {
		return cfg, err
	}

	err = ko.Unmarshal("", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
