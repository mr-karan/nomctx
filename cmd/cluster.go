package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func listClusters(cfg Config, stdout io.Writer) {
	for _, cluster := range cfg.Clusters {
		for name := range cluster {
			fmt.Fprintf(stdout, "%s\n", name)
		}
	}
}

// Wraps around `nomctx` around `fzf` to show a prompt for list of clusters to switch.
// Returns the cluster selected by user.
func switchCluster() (string, error) {
	var (
		cmd = exec.Command("fzf", "--ansi", "--no-preview")
		out bytes.Buffer
	)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = &out

	// os.Args[0] is the current program. It basically is doing the equivalent of `nomctx --list | fzf`.
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("FZF_DEFAULT_COMMAND=%s --config %s --list-clusters", os.Args[0], ko.String("config")))

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
