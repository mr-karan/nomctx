package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	DEFAULT_CONFIG_FILE = "config.hcl"
)

var (
	// Version of the build. This is injected at build-time.
	buildString           = "unknown"
	defaultConfigFilePath string
)

func init() {
	path, err := getFilePathInHomeDir(DEFAULT_CONFIG_FILE)
	if err != nil {
		log.Fatalf("Failed to get default config file path: %v", err)
	}

	defaultConfigFilePath = path
}

func main() {
	app := &cli.App{
		Name:  "nomctx",
		Usage: "Faster way to switch across multiple Nomad clusters and namespaces",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: defaultConfigFilePath,
				Usage: "Path to a config file to load.",
			},
		},
		Before: func(c *cli.Context) error {
			cfg, err := initConfig(c.String("config"))
			if err != nil {
				return fmt.Errorf("error initialising config: %w", err)
			}
			// Set the config in the app metadata.
			c.App.Metadata["cfg"] = cfg
			return nil
		},
		// Default.
		Action: func(c *cli.Context) error {
			if c.Args().Len() == 0 {
				// No command provided.
				return handleSwitchCluster(c)
			}
			return cli.ShowAppHelp(c)
		},
		Commands: []*cli.Command{
			{
				Name:   "list-clusters",
				Usage:  "List all clusters",
				Action: handleListClusters,
			},
			{
				Name:   "list-namespaces",
				Usage:  "List all namespaces",
				Action: handleListNamespaces,
			},
			{
				Name:      "set-cluster",
				Usage:     "Set the current cluster context",
				ArgsUsage: "CLUSTER_NAME",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "persist",
						Usage: "Persist the environment variables to a .env file",
					},
				},
				Action: handleSetCluster,
			},
			{
				Name:      "set-namespace",
				Usage:     "Set namespace",
				ArgsUsage: "NAMESPACE",
				Action:    handleSetNamespace,
			},
			{
				Name:  "switch-cluster",
				Usage: "Switch cluster",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "persist",
						Usage: "Persist the environment variables to a .env file",
					},
				},
				Action: handleSwitchCluster,
			},
			{
				Name:   "switch-namespace",
				Usage:  "Switch namespace",
				Action: handleSwitchNamespace,
			},
			{
				Name:   "current-context",
				Usage:  "Display the current context",
				Action: handleCurrentCtx,
			},
			{
				Name:      "login",
				Usage:     "Login to a cluster",
				ArgsUsage: "CLUSTER",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "cluster",
						Usage: "Name of the cluster to login to",
					},
					&cli.BoolFlag{
						Name:  "persist",
						Usage: "Persist the environment variables to a .env file",
					},
				},
				Action: handleLogin,
			},
		},
		Version: buildString,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
