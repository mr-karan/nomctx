package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

// persistContext writes the current context to ~/.nomctx/context.hcl.
func persistContext(context ContextCfg) error {
	contextFilePath, err := getFilePathInHomeDir("context.hcl")
	if err != nil {
		return err
	}
	contextFileContent := fmt.Sprintf("namespace = %q\ncluster = %q\n", context.Namespace, context.Cluster)

	if err := os.MkdirAll(filepath.Dir(contextFilePath), 0755); err != nil {
		return err
	}

	if err := ioutil.WriteFile(contextFilePath, []byte(contextFileContent), 0644); err != nil {
		return err
	}

	return nil
}

// loadContext reads the current context from ~/.nomctx/context.hcl.
func loadContext() (ContextCfg, error) {
	contextFilePath, err := getFilePathInHomeDir("context.hcl")
	if err != nil {
		return ContextCfg{}, err
	}

	var context ContextCfg
	if err := hclsimple.DecodeFile(contextFilePath, nil, &context); err != nil {
		return ContextCfg{}, fmt.Errorf("error loading context file %s: %w", contextFilePath, err)
	}

	return context, nil
}
