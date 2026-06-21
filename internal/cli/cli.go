package cli

import (
	"fmt"
	"io"
	"strings"

	"github.com/lucasdeprit/spec-forge/internal/workspace"
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
	case "init":
		force, err := parseInitArgs(args[1:])
		if err != nil {
			fmt.Fprintln(stderr, err)
			return err
		}

		if err := workspace.Init(".", workspace.InitOptions{Force: force}); err != nil {
			fmt.Fprintln(stderr, err)
			return err
		}

		fmt.Fprintln(stdout, "initialized .specforge workspace")
		return nil
	case "repo":
		options, err := parseRepoArgs(args[1:])
		if err != nil {
			fmt.Fprintln(stderr, err)
			return err
		}

		if err := workspace.AddRepository(".", options); err != nil {
			fmt.Fprintln(stderr, err)
			return err
		}

		fmt.Fprintf(stdout, "registered repo %s\n", options.ID)
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

func parseInitArgs(args []string) (bool, error) {
	force := false
	for _, arg := range args {
		switch arg {
		case "--force":
			force = true
		default:
			return false, fmt.Errorf("unknown init option: %s", arg)
		}
	}

	return force, nil
}

func parseRepoArgs(args []string) (workspace.RepositoryOptions, error) {
	if len(args) != 5 || args[0] != "add" || args[3] != "--type" {
		return workspace.RepositoryOptions{}, fmt.Errorf("usage: specforge repo add <repo-id> <path> --type <type>")
	}

	return workspace.RepositoryOptions{
		ID:   args[1],
		Path: args[2],
		Type: args[4],
	}, nil
}

func printHelp(w io.Writer) {
	fmt.Fprint(w, strings.TrimPrefix(`
SpecForge structures specification-driven development workspaces.

Usage:
  specforge <command>

Commands:
  init       Create a SpecForge workspace
  repo add   Register a repository in the workspace
  help       Show this help message
  version    Show the SpecForge version
`, "\n"))
}
