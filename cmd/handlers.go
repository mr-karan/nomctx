package main

import (
	"fmt"
	"log"
	"os"
)

func handleListNamespaces(cfg Config) {
	ns, err := listNamespaces(cfg)
	if err != nil {
		log.Fatalf("error listing namespace: %s", err.Error())
	}
	fmt.Fprintf(os.Stdout, "%s", ns)
	os.Exit(0)
}

func handleListClusters(cfg Config) {
	listClusters(cfg, os.Stdout)
	os.Exit(0)
}

func handleSetNamespace() {
	exportNamespace(ko.String("set-namespace"), os.Stdout)
	os.Exit(0)
}

func handleSetCluster(cfg Config) {
	// Fetch cluster metadata.
	cluster, err := lookupCluster(ko.String("set-cluster"), cfg)
	if err != nil {
		log.Fatalf("error looking cluster: %v", err)
	}

	// Emit `export` commands to `stdout`.
	exportClusterVars(cluster, os.Stdout)
	os.Exit(0)
}

func handleSwitchNamespace(cfg Config) {
	if !isFZFInstalled() {
		// Fallback to just listing of namespaces.
		handleListNamespaces(cfg)
	} else {
		ns, err := switchNamespace()
		if err != nil {
			log.Fatalf("error fetching namespace")
		}
		exportNamespace(ns, os.Stdout)
		os.Exit(0)
	}

}

func handleSwitchCluster(cfg Config) {
	if !isFZFInstalled() {
		// Fallback to just listing of clusters.
		handleListClusters(cfg)
	} else {
		// If `fzf` exists, we can show a prompt.
		c, err := switchCluster()
		if err != nil {
			log.Fatalf("error fetching cluster")
		}

		// Fetch cluster metadata.
		cluster, err := lookupCluster(c, cfg)
		if err != nil {
			log.Fatalf("error looking cluster: %v", err)
		}

		// Emit `export` commands to `stdout`.
		exportClusterVars(cluster, os.Stdout)
		os.Exit(0)
	}
}
