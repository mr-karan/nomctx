package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func listNamespaces(cfg Config) (string, error) {
	output, err := exec.Command("nomad", "namespace", "list", "-t", "'{{range .}}{{printf \"%s\\n\" .Name}}{{end}}'").CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

// Wraps around `nomad namespace list` around `fzf` to show a prompt for list of namespaces to switch.
// Returns the namespace selected by user.
func switchNamespace() (string, error) {
	var (
		cmd = exec.Command("fzf", "--ansi", "--no-preview")
		out bytes.Buffer
	)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = &out

	// os.Args[0] is the current program. It basically is doing the equivalent of `nomctx --list | fzf`.
	cmd.Env = append(os.Environ(), "FZF_DEFAULT_COMMAND=nomad namespace list -t  '{{range .}}{{printf \"%s\\n\" .Name}}{{end}}'")
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}
