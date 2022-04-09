package main

import (
	"os"
	"os/exec"
	"path"
)

// Returns the default config path which is `~/.nomctx/config.hcl`.
func getDefaultCfgPath() string {
	var (
		cfgDefaultPath string
	)

	dir, _ := os.UserHomeDir()
	if dir != "" {
		cfgDefaultPath = path.Join(dir, ".nomctx/config.hcl")
	}

	return cfgDefaultPath
}

// isFZFInstalled determines if `fzf` exists in `PATH`.
func isFZFInstalled() bool {
	ok, _ := exec.LookPath("fzf")
	return ok != ""
}
