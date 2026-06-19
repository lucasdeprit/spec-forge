package cli

import (
	"bytes"
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
	for _, expected := range []string{"Usage:", "Commands:", "version"} {
		if !strings.Contains(output, expected) {
			t.Fatalf("stdout %q does not contain %q", output, expected)
		}
	}

	if stderr.Len() != 0 {
		t.Fatalf("stderr = %q, want empty", stderr.String())
	}
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
