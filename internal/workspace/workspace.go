package workspace

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
		filepath.Join(workspaceRoot, "workspace.html"):          defaultWorkspaceHTML,
		filepath.Join(workspaceRoot, "assets", "specforge.css"): defaultCSS,
	}

	for path, content := range files {
		if err := writeFile(path, content, options.Force); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(path string, content string, force bool) error {
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

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	return nil
}

const defaultWorkspaceHTML = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>SpecForge Workspace</title>
    <link rel="stylesheet" href="assets/specforge.css">
  </head>
  <body>
    <sf-workspace>
      <sf-title>SpecForge Workspace</sf-title>
      <sf-domains></sf-domains>
      <sf-repositories></sf-repositories>
      <sf-proposals></sf-proposals>
      <sf-tasks></sf-tasks>
    </sf-workspace>
  </body>
</html>
`

const defaultCSS = `body {
  margin: 0;
  padding: 2rem;
  background: #f8fafc;
  color: #111827;
  font-family: ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
}

sf-workspace,
sf-title,
sf-domains,
sf-repositories,
sf-proposals,
sf-tasks {
  display: block;
}

sf-workspace {
  max-width: 960px;
  margin: 0 auto;
}

sf-title {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 1rem;
}

sf-domains,
sf-repositories,
sf-proposals,
sf-tasks {
  background: white;
  border: 1px solid #e5e7eb;
  border-radius: 12px;
  padding: 1rem;
  margin-bottom: 1rem;
}

sf-domains::before {
  content: "Domains";
  display: block;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

sf-repositories::before {
  content: "Repositories";
  display: block;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

sf-proposals::before {
  content: "Proposals";
  display: block;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

sf-tasks::before {
  content: "Tasks";
  display: block;
  font-weight: 700;
  margin-bottom: 0.5rem;
}
`
