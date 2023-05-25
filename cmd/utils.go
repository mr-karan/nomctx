package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// isFZFInstalled determines if `fzf` exists in `PATH`.
func isFZFInstalled() bool {
	ok, _ := exec.LookPath("fzf")
	return ok != ""
}

// selectInteractive presents a list of choices to the user and returns the
// selected choice.
func selectInteractive(choices []string) (string, error) {
	var (
		cmd = exec.Command("fzf", "--ansi", "--no-preview")
		out bytes.Buffer
	)

	cmd.Stdin = strings.NewReader(strings.Join(choices, "\n"))
	cmd.Stderr = os.Stderr
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return strings.TrimSpace(out.String()), nil
}

// getFilePathInHomeDir returns the full path of a file in the home directory.
func getFilePathInHomeDir(fileName string) (string, error) {
	// Get the home directory.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	// Construct the full path of the file.
	return filepath.Join(homeDir, ".nomctx", fileName), nil
}

// sanitizeFilename cleans up string to be used as a safe filename.
func sanitizeFilename(name string) string {
	// Create a regular expression that matches invalid filename characters.
	re := regexp.MustCompile(`[^a-zA-Z0-9\._-]`)

	// Replace invalid characters with a hyphen.
	safeName := re.ReplaceAllString(name, "-")

	return safeName
}
