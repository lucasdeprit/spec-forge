package cli

import (
	"fmt"
	"io"
)

const Version = "0.1.0"

func Run(args []string, stdout io.Writer, stderr io.Writer) error {
	if len(args) == 0 {
		printHelp(stdout)
		return nil
	}

	switch args[0] {
	case "help", "--help", "-h":
		printHelp(stdout)
		return nil
	case "version", "--version", "-v":
		fmt.Fprintf(stdout, "specforge %s\n", Version)
		return nil
	default:
		fmt.Fprintf(stderr, "unknown command: %s\n\n", args[0])
		printHelp(stderr)
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func printHelp(w io.Writer) {
	fmt.Fprint(w, `SpecForge structures specification-driven development workspaces.

Usage:
  specforge <command>

Commands:
  help      Show this help message
  version   Show the SpecForge version
`)
}
