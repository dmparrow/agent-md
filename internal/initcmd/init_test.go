package initcmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/dmparrow/agent-md/internal/config"
)

func TestInitCreatesStructureAndPreservesConfig(t *testing.T) {
	root := t.TempDir()

	if err := Init(root); err != nil {
		t.Fatalf("Init() error = %v", err)
	}

	agentDir := filepath.Join(root, config.DirName)
	runsDir := filepath.Join(agentDir, config.RunsDirName)
	configPath := filepath.Join(agentDir, config.ConfigName)

	assertDirExists(t, agentDir)
	assertDirExists(t, runsDir)

	content, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if got, want := string(content), config.DefaultYAML; got != want {
		t.Fatalf("config content = %q, want %q", got, want)
	}

	customContent := []byte("version: 99\n")
	if err := os.WriteFile(configPath, customContent, 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	if err := Init(root); err != nil {
		t.Fatalf("second Init() error = %v", err)
	}

	content, err = os.ReadFile(configPath)
	if err != nil {
		t.Fatalf("ReadFile() after second init error = %v", err)
	}

	if got, want := string(content), string(customContent); got != want {
		t.Fatalf("config content after second init = %q, want %q", got, want)
	}
}

func assertDirExists(t *testing.T, path string) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat(%q) error = %v", path, err)
	}

	if !info.IsDir() {
		t.Fatalf("%q is not a directory", path)
	}
}
