package main

import (
	"fmt"
	"log"
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/posflag"
	flag "github.com/spf13/pflag"
)

var (
	// Version of the build. This is injected at build-time.
	buildString = "unknown"
	ko          = koanf.New(".")
)

type Config struct {
	Clusters []map[string][]Cluster `koanf:"clusters"`
}

type Cluster struct {
	Address   string `koanf:"address"`
	Namespace string `koanf:"namespace"`
	Region    string `koanf:"region"`
	Token     string `koanf:"token"`
	HTTPAuth  string `koanf:"http_auth"`
}

func main() {
	f := flag.NewFlagSet("nomctx", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

	// Register flags.
	f.String("config", getDefaultCfgPath(), "Path to a config file to load.")
	f.BoolP("version", "v", false, "Show version of nomctx")

	// Cluster commands.
	f.Bool("switch-cluster", true, "Switch cluster") // Default action.
	f.Bool("list-clusters", false, "List all clusters")
	f.String("set-cluster", "", "Set cluster")

	// Namesapce commands.
	f.Bool("switch-namespace", false, "Switch namespace")
	f.Bool("list-namespaces", false, "List all namespaces")
	f.String("set-namespace", "", "Set namespace")

	// Parse and Load Flags.
	err := f.Parse(os.Args[1:])
	if err != nil {
		log.Fatalf("error parsing flags: %v", err)
	}
	if err = ko.Load(posflag.Provider(f, ".", ko), nil); err != nil {
		log.Fatalf("error loading flags: %v", err)
	}

	// Initialise config.
	cfg, err := initConfig(ko, ko.String("config"))
	if err != nil {
		log.Fatalf("error initialising config: %v", err)
	}

	// If version flag is set, output version and quit.
	if ko.Bool("version") {
		fmt.Printf("%s\n", buildString)
		os.Exit(0)
	}

	if ko.Bool("list-clusters") {
		handleListClusters(cfg)
	}

	if ko.Bool("list-namespaces") {
		handleListNamespaces(cfg)
	}

	if ko.String("set-namespace") != "" {
		handleSetNamespace()
	}

	if ko.String("set-cluster") != "" {
		handleSetCluster(cfg)
	}

	if ko.Bool("switch-namespace") {
		handleSwitchNamespace(cfg)
	}

	if ko.Bool("switch-cluster") {
		handleSwitchCluster(cfg)
	}
}
