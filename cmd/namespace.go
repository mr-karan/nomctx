package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// setNamespace sets the namespace provided in the current context.
func setNamespace(ns string) error {
	ok, err := lookupNamespace(ns)
	if err != nil {
		return fmt.Errorf("failed to lookup namespace: %w", err)
	}
	if !ok {
		return fmt.Errorf("namespace %s does not exist", ns)
	}

	// Load current context.
	context, err := loadContext()
	if err != nil {
		return fmt.Errorf("failed to load context: %w", err)
	}

	// Update the context.
	context.Namespace = ns
	if err := persistContext(context); err != nil {
		return fmt.Errorf("failed to persist context: %w", err)
	}

	// Output the export command.
	fmt.Fprintf(os.Stdout, "export %s=%s\n", "NOMAD_NAMESPACE", ns)

	return nil
}

// listNamespaces returns a list of namespaces.
func listNamespaces() ([]string, error) {
	output, err := exec.Command("nomad", "namespace", "list", "-t", "{{range .}}{{printf \"%s\\n\" .Name}}{{end}}").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces, error: %v, output: %s", err, output)
	}
	namespaces := strings.Split(strings.TrimSpace(string(output)), "\n")
	return namespaces, nil
}

// Checks the status of namespace, whether it exists or not.
func lookupNamespace(ns string) (bool, error) {
	output, err := exec.Command("nomad", "namespace", "status", ns).CombinedOutput()
	if err != nil {
		if strings.Contains(string(output), fmt.Sprintf("Namespace \"%s\" matched no namespaces", ns)) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check namespace status, error: %v, output: %s", err, output)
	}
	return true, nil
}
