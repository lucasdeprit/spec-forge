package workspace

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lucasdeprit/spec-forge/assets"
)

type InitOptions struct {
	Force bool
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
