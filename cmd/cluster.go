package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func setCluster(cluster ClusterCfg, persist bool) error {
	// Persist the context.
	context := ContextCfg{
		Cluster:   cluster.Name,
		Namespace: "default",
	}
	if err := persistContext(context); err != nil {
		return fmt.Errorf("error persisting context: %w", err)
	}

	// Print the env variables to stdout or persist them to a file.
	if persist {
		basePath, err := getFilePathInHomeDir("")
		if err != nil {
			return err
		}

		// Persist the environment variables.
		return persistClusterVarsFile(cluster, basePath)
	}

	// If --persist is not passed, just print the env variables to stdout.
	exportClusterVars(cluster, os.Stdout)
	return nil
}

func listClusters(cfg Config) []string {
	clusters := make([]string, 0, len(cfg.Clusters))
	for _, c := range cfg.Clusters {
		clusters = append(clusters, c.Name)
	}
	return clusters
}

// Returns the metadata for a particular cluster.
func lookupCluster(name string, clusters []ClusterCfg) (ClusterCfg, error) {
	for _, c := range clusters {
		if c.Name == name {
			return c, nil
		}
	}
	return ClusterCfg{}, fmt.Errorf("no cluster with name %s found", name)
}

// Emits `export` commands on shell.
func exportClusterVars(c ClusterCfg, out io.Writer) {
	printIfNotEmpty := func(varName, value string) {
		if value != "" {
			fmt.Fprintf(out, "export %s=%s\n", varName, value)
		}
	}

	printIfNotEmpty("NOMAD_ADDR", c.Address)
	printIfNotEmpty("NOMAD_TOKEN", c.Token)
	printIfNotEmpty("NOMAD_HTTP_AUTH", c.HTTPAuth)
	printIfNotEmpty("NOMAD_REGION", c.Region)
	printIfNotEmpty("NOMAD_CACERT", c.CACert)
	printIfNotEmpty("NOMAD_CLIENT_CERT", c.ClientCert)
	printIfNotEmpty("NOMAD_CLIENT_KEY", c.ClientKey)
}

// persistClusterVars writes the cluster variables to a file
func persistClusterVars(c ClusterCfg, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "NOMAD_ADDR=%s\n", c.Address)
	if c.Token != "" {
		fmt.Fprintf(file, "NOMAD_TOKEN=%s\n", c.Token)
	}
	if c.HTTPAuth != "" {
		fmt.Fprintf(file, "NOMAD_HTTP_AUTH=%s\n", c.HTTPAuth)
	}
	if c.Region != "" {
		fmt.Fprintf(file, "NOMAD_REGION=%s\n", c.Region)
	}
	if c.Namespace != "" {
		fmt.Fprintf(file, "NOMAD_NAMESPACE=%s\n", c.Namespace)
	}
	if c.CACert != "" {
		fmt.Fprintf(file, "NOMAD_CACERT=%s\n", c.CACert)
	}
	if c.ClientCert != "" {
		fmt.Fprintf(file, "NOMAD_CLIENT_CERT=%s\n", c.ClientCert)
	}
	if c.ClientKey != "" {
		fmt.Fprintf(file, "NOMAD_CLIENT_KEY=%s\n", c.ClientKey)
	}

	return nil
}

// persistClusterVarsFile persists the cluster variables to a file.
func persistClusterVarsFile(cluster ClusterCfg, basePath string) error {
	sanitizedClusterName := sanitizeFilename(cluster.Name)
	filePath := filepath.Join(basePath, sanitizedClusterName+".env")

	// Ensure that the directory exists.
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	// Persist the environment variables.
	return persistClusterVars(cluster, filePath)
}
