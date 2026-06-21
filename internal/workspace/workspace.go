package workspace

import (
	"errors"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"strings"

	"github.com/lucasdeprit/spec-forge/assets"
)

type InitOptions struct {
	Force bool
}

type RepositoryOptions struct {
	ID   string
	Path string
	Type string
}

func Init(root string, options InitOptions) error {
	workspaceRoot := filepath.Join(root, ".specforge")

	for _, dir := range []string{
		workspaceRoot,
		filepath.Join(workspaceRoot, "proposals"),
		filepath.Join(workspaceRoot, "tasks"),
		filepath.Join(workspaceRoot, "repos"),
		filepath.Join(workspaceRoot, "assets"),
		filepath.Join(workspaceRoot, "templates"),
	} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create %s: %w", dir, err)
		}
	}

	files := map[string]string{
		filepath.Join(workspaceRoot, "workspace.html"):          "workspace.html",
		filepath.Join(workspaceRoot, "assets", "specforge.css"): "specforge.css",
	}

	for path, source := range files {
		content, err := assets.Files.ReadFile(source)
		if err != nil {
			return fmt.Errorf("read embedded %s: %w", source, err)
		}

		if err := writeFile(path, content, options.Force); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(path string, content []byte, force bool) error {
	flag := os.O_WRONLY | os.O_CREATE | os.O_EXCL
	if force {
		flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}

	file, err := os.OpenFile(path, flag, 0o644)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		}
		return fmt.Errorf("write %s: %w", path, err)
	}
	defer file.Close()

	if _, err := file.Write(content); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	return nil
}

func AddRepository(root string, options RepositoryOptions) error {
	if options.ID == "" {
		return fmt.Errorf("repo id is required")
	}

	if options.Path == "" {
		return fmt.Errorf("repo path is required")
	}

	if filepath.IsAbs(options.Path) {
		return fmt.Errorf("repo path must be relative to the SpecForge workspace")
	}

	if !isSupportedRepoType(options.Type) {
		return fmt.Errorf("unsupported repo type: %s", options.Type)
	}

	repoPath := filepath.Clean(options.Path)
	if _, err := os.Stat(filepath.Join(root, repoPath)); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("repo path does not exist: %s", options.Path)
		}
		return fmt.Errorf("check repo path %s: %w", options.Path, err)
	}

	workspacePath := filepath.Join(root, ".specforge", "workspace.html")
	content, err := os.ReadFile(workspacePath)
	if err != nil {
		return fmt.Errorf("read %s: %w", workspacePath, err)
	}

	workspaceHTML := string(content)
	escapedID := html.EscapeString(options.ID)
	if strings.Contains(workspaceHTML, `id="`+escapedID+`"`) {
		return fmt.Errorf("duplicate repo id: %s", options.ID)
	}

	openRepositoriesTag := "<sf-repositories>"
	closingTag := "</sf-repositories>"
	if !strings.Contains(workspaceHTML, openRepositoriesTag) || !strings.Contains(workspaceHTML, closingTag) {
		return fmt.Errorf("workspace.html missing sf-repositories")
	}

	repository := fmt.Sprintf(
		`      <sf-repository id="%s" path="%s" type="%s"></sf-repository>`+"\n",
		escapedID,
		html.EscapeString(filepath.ToSlash(repoPath)),
		html.EscapeString(options.Type),
	)

	if strings.Contains(workspaceHTML, openRepositoriesTag+closingTag) {
		workspaceHTML = strings.Replace(workspaceHTML, openRepositoriesTag+closingTag, openRepositoriesTag+"\n"+repository+"      "+closingTag, 1)
	} else {
		workspaceHTML = strings.Replace(workspaceHTML, closingTag, repository+"      "+closingTag, 1)
	}
	if err := os.WriteFile(workspacePath, []byte(workspaceHTML), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", workspacePath, err)
	}

	return nil
}

func isSupportedRepoType(repoType string) bool {
	switch repoType {
	case "backend", "frontend", "mobile", "infra", "library", "shared":
		return true
	default:
		return false
	}
}
