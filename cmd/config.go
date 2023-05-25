package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type ClusterCfg struct {
	Name      string   `hcl:",label"`
	Address   string   `hcl:"address"`
	Namespace string   `hcl:"namespace,optional"`
	Region    string   `hcl:"region,optional"`
	HTTPAuth  string   `hcl:"http_auth,optional"`
	Token     string   `hcl:"token,optional"`
	Auth      *AuthCfg `hcl:"auth,block"`
}

type AuthCfg struct {
	Method   string `hcl:"method"`
	Provider string `hcl:"provider"`
}

// ContextCfg is the data structure for storing the currently active context.
type ContextCfg struct {
	Cluster   string `hcl:"cluster"`
	Namespace string `hcl:"namespace"`
}

type Config struct {
	Clusters []ClusterCfg `hcl:"cluster,block"`
}

// initConfig parses the config file and loads in `Config` object.
func initConfig(cfgPath string) (Config, error) {
	var (
		cfg Config
	)

	if err := hclsimple.DecodeFile(cfgPath, nil, &cfg); err != nil {
		return cfg, fmt.Errorf("error loading config file %s: %w", cfgPath, err)
	}

	return cfg, nil
}
