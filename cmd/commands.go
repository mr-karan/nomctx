package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

// handleListClusters lists all the configured clusters.
func handleListClusters(c *cli.Context) error {
	cfg := c.App.Metadata["cfg"].(Config)
	clusters := listClusters(cfg)
	for _, cluster := range clusters {
		fmt.Fprintln(c.App.Writer, cluster)
	}
	return nil
}

// handleListNamespaces lists all the namespaces.
func handleListNamespaces(c *cli.Context) error {
	namespaces, err := listNamespaces()
	if err != nil {
		return err
	}
	for _, ns := range namespaces {
		fmt.Fprintln(c.App.Writer, ns)
	}
	return nil
}

// handleSetCluster sets the current cluster context.
func handleSetCluster(c *cli.Context) error {
	cfg := c.App.Metadata["cfg"].(Config)
	if c.NArg() == 0 {
		return errors.New("cluster name is required")
	}
	cName := c.Args().First()

	// Check if cluster is valid and exists in cfg.
	cluster, err := lookupCluster(cName, cfg.Clusters)
	if err != nil {
		return fmt.Errorf("error finding cluster: %w", err)
	}

	return setCluster(cluster, c.Bool("persist"))
}

// handleSetNamespace sets the current namespace.
func handleSetNamespace(c *cli.Context) error {
	if c.NArg() == 0 {
		return errors.New("namespace name is required")
	}
	ns := c.Args().First()
	if err := setNamespace(ns); err != nil {
		return err
	}
	return nil
}

// handleSwitchCluster switches the cluster context to a specified cluster, or prompts
// the user to select a cluster if no cluster is specified.
func handleSwitchCluster(c *cli.Context) error {
	cfg := c.App.Metadata["cfg"].(Config)
	cName := c.Args().First()

	// If cluster name is not provided, prompt the user to select a cluster.
	if cName == "" {
		clusters := listClusters(cfg)
		if isFZFInstalled() {
			var err error
			cName, err = selectInteractive(clusters)
			if err != nil {
				return fmt.Errorf("failed to select cluster interactively: %w", err)
			}
		} else {
			return fmt.Errorf("please provide a cluster name")
		}
	}

	// Check if cluster is valid and exists in cfg.
	cluster, err := lookupCluster(cName, cfg.Clusters)
	if err != nil {
		return fmt.Errorf("error finding cluster: %w", err)
	}

	// Switch the context by setting the env variables using handleSetCluster.
	if err := setCluster(cluster, c.Bool("persist")); err != nil {
		return fmt.Errorf("failed to switch cluster: %w", err)
	}

	return nil
}

// handleSwitchNamespace switches the namespace provided in the current context interactively.
func handleSwitchNamespace(c *cli.Context) error {
	ns := c.Args().First()

	// If the namespace is provided, set it directly.
	if ns != "" {
		if err := setNamespace(ns); err != nil {
			return fmt.Errorf("failed to set namespace: %w", err)
		}
		return nil
	}

	// If namespace is not provided, fetch the list of namespaces.
	namespaces, err := listNamespaces()
	if err != nil {
		return err
	}

	// Let the user select a namespace interactively.
	selectedNamespace, err := selectInteractive(namespaces)
	if err != nil {
		return fmt.Errorf("failed to select namespace interactively: %w", err)
	}

	// Set the selected namespace.
	if err := setNamespace(selectedNamespace); err != nil {
		return fmt.Errorf("failed to set namespace: %w", err)
	}
	return nil
}

// handleLogin logs into a cluster and prints the SecretID and ExpirationTTL.
func handleLogin(c *cli.Context) error {
	clusterName := c.String("cluster")

	if clusterName == "" {
		context, err := loadContext()
		if err != nil {
			return fmt.Errorf("failed to load context: %w", err)
		}
		clusterName = context.Cluster
	}

	// Lookup the cluster from the list of configured clusters.
	cfg := c.App.Metadata["cfg"].(Config)
	cluster, err := lookupCluster(clusterName, cfg.Clusters)
	if err != nil {
		return fmt.Errorf("failed to find cluster: %w", err)
	}

	// Check if the cluster has proper auth configured.
	if cluster.Auth == nil {
		return fmt.Errorf("cluster %s does not have auth configured", cluster.Name)
	}
	if cluster.Auth.Provider != "nomad" {
		return fmt.Errorf("unsupported provider: %s", cluster.Auth.Provider)
	}

	cmd := exec.Command("nomad", "login", "-method="+cluster.Auth.Method, "-address="+cluster.Address, "-json")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to login, output: %s, error: %w", output, err)
	}

	var loginResult struct {
		SecretID      string `json:"SecretID"`
		ExpirationTTL string `json:"ExpirationTTL"`
	}

	if err := json.Unmarshal(output, &loginResult); err != nil {
		return fmt.Errorf("failed to parse login result: %w", err)
	}

	// Set the token for the given cluster.
	cluster.Token = loginResult.SecretID

	return setCluster(cluster, c.Bool("persist"))
}

func handleCurrentCtx(c *cli.Context) error {
	context, err := loadContext()
	if err != nil {
		return fmt.Errorf("failed to load context: %w", err)
	}
	fmt.Fprintf(c.App.Writer, "Cluster: %s\nNamespace: %s\n", context.Cluster, context.Namespace)
	return nil
}
