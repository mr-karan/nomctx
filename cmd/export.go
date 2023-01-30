package main

import (
	"fmt"
	"io"
)

// Returns the metadata for a particular cluster.
func lookupCluster(name string, cfg Config) (Cluster, error) {
	for _, c := range cfg.Clusters {
		if v, found := c[name]; found {
			return v[0], nil
		}
	}
	return Cluster{}, fmt.Errorf("no cluster with name %s found", name)
}

// Emits `export` commands on shell.
func exportClusterVars(c Cluster, out io.Writer) {
	fmt.Fprintf(out, "export NOMAD_ADDR=%s\n", c.Address)
	if c.Token != "" {
		fmt.Fprintf(out, "export NOMAD_TOKEN=%s\n", c.Token)
	}
	if c.HTTPAuth != "" {
		fmt.Fprintf(out, "export NOMAD_HTTP_AUTH=%s\n", c.HTTPAuth)
	}
	if c.Region != "" {
		fmt.Fprintf(out, "export NOMAD_REGION=%s\n", c.Region)
	}
	if c.Namespace != "" {
		fmt.Fprintf(out, "export NOMAD_NAMESPACE=%s\n", c.Namespace)
	}
}

func exportNamespace(n string, out io.Writer) {
	if n != "" {
		fmt.Fprintf(out, "export NOMAD_NAMESPACE=%s\n", n)
	}
}
