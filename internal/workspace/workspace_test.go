package workspace

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lucasdeprit/spec-forge/assets"
)

func TestInitCreatesExpectedStructure(t *testing.T) {
	root := t.TempDir()

	if err := Init(root, InitOptions{}); err != nil {
		t.Fatalf("Init returned error: %v", err)
	}

	for _, path := range []string{
		".specforge/workspace.html",
		".specforge/proposals",
		".specforge/tasks",
		".specforge/repos",
		".specforge/assets/specforge.css",
		".specforge/templates",
	} {
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(path))); err != nil {
			t.Fatalf("expected %s to exist: %v", path, err)
		}
	}

	workspaceHTML := readFile(t, filepath.Join(root, ".specforge", "workspace.html"))
	defaultWorkspaceHTML := readEmbeddedFile(t, "workspace.html")
	if workspaceHTML != defaultWorkspaceHTML {
		t.Fatalf("workspace.html does not match embedded source")
	}

	for _, expected := range []string{"<!doctype html>", "<sf-workspace>", "assets/specforge.css"} {
		if !strings.Contains(workspaceHTML, expected) {
			t.Fatalf("workspace.html does not contain %q", expected)
		}
	}

	css := readFile(t, filepath.Join(root, ".specforge", "assets", "specforge.css"))
	defaultCSS := readEmbeddedFile(t, "specforge.css")
	if css != defaultCSS {
		t.Fatalf("specforge.css does not match embedded source")
	}
}

func TestInitTwiceDoesNotOverwriteExistingFiles(t *testing.T) {
	root := t.TempDir()

	if err := Init(root, InitOptions{}); err != nil {
		t.Fatalf("Init returned error: %v", err)
	}

	workspacePath := filepath.Join(root, ".specforge", "workspace.html")
	if err := os.WriteFile(workspacePath, []byte("custom workspace"), 0o644); err != nil {
		t.Fatalf("write custom workspace: %v", err)
	}

	if err := Init(root, InitOptions{}); err != nil {
		t.Fatalf("second Init returned error: %v", err)
	}

	if got, want := readFile(t, workspacePath), "custom workspace"; got != want {
		t.Fatalf("workspace.html = %q, want %q", got, want)
	}
}

func TestInitForceOverwritesGeneratedFiles(t *testing.T) {
	root := t.TempDir()

	if err := Init(root, InitOptions{}); err != nil {
		t.Fatalf("Init returned error: %v", err)
	}

	workspacePath := filepath.Join(root, ".specforge", "workspace.html")
	if err := os.WriteFile(workspacePath, []byte("custom workspace"), 0o644); err != nil {
		t.Fatalf("write custom workspace: %v", err)
	}

	if err := Init(root, InitOptions{Force: true}); err != nil {
		t.Fatalf("force Init returned error: %v", err)
	}

	if got := readFile(t, workspacePath); !strings.Contains(got, "<sf-workspace>") {
		t.Fatalf("workspace.html = %q, want generated workspace", got)
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}

	return string(content)
}

func readEmbeddedFile(t *testing.T, path string) string {
	t.Helper()

	content, err := assets.Files.ReadFile(path)
	if err != nil {
		t.Fatalf("read embedded %s: %v", path, err)
	}

	return string(content)
}
