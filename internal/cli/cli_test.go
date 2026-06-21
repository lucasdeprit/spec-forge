package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunVersion(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	err := Run([]string{"version"}, &stdout, &stderr)
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	if got, want := stdout.String(), "specforge "+Version+"\n"; got != want {
		t.Fatalf("stdout = %q, want %q", got, want)
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunHelp(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	err := Run([]string{"help"}, &stdout, &stderr)
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	output := stdout.String()
	for _, expected := range []string{"Usage:", "Commands:", "init", "version"} {
		if !strings.Contains(output, expected) {
			t.Fatalf("stdout %q does not contain %q", output, expected)
		}
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
}

func TestRunInit(t *testing.T) {
	chdir(t, t.TempDir())

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	err := Run([]string{"init"}, &stdout, &stderr)
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}

	for _, path := range []string{
		".specforge/workspace.html",
		".specforge/proposals",
		".specforge/tasks",
		".specforge/repos",
		".specforge/assets/specforge.css",
		".specforge/templates",
	} {
		if _, err := os.Stat(filepath.FromSlash(path)); err != nil {
			t.Fatalf("expected %s to exist: %v", path, err)
		}
	}

	if !strings.Contains(stdout.String(), "initialized .specforge workspace") {
		t.Fatalf("stdout = %q, want init message", stdout.String())
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
}

func chdir(t *testing.T, dir string) {
	t.Helper()

	previous, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working directory: %v", err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatalf("change working directory: %v", err)
	}

	t.Cleanup(func() {
		if err := os.Chdir(previous); err != nil {
			t.Fatalf("restore working directory: %v", err)
		}
	})
}

func TestRunUnknownCommand(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	err := Run([]string{"nope"}, &stdout, &stderr)
	if err == nil {
		t.Fatal("Run returned nil error, want error")
	}

	if stdout.Len() != 0 {
		t.Fatalf("stdout = %q, want empty", stdout.String())
	}

	if !strings.Contains(stderr.String(), "unknown command: nope") {
		t.Fatalf("stderr = %q, want unknown command message", stderr.String())
	}
}
