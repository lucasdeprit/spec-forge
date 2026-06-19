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
  min-height: 100vh;
  padding: 4.5rem 2.75rem;
  background: #f3eadb;
  color: #050505;
  font-family: -apple-system, BlinkMacSystemFont, "Helvetica Neue", "Segoe UI", sans-serif;
  font-size: 18px;
  line-height: 1.62;
}

body * {
  box-sizing: border-box;
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
  max-width: 1180px;
  margin: 0 auto;
}

sf-title {
  color: #050505;
  font-family: "SFMono-Regular", "Roboto Mono", "Menlo", monospace;
  font-size: clamp(3rem, 9vw, 7.5rem);
  font-weight: 900;
  letter-spacing: -0.09em;
  line-height: 0.88;
  margin-bottom: 2.75rem;
  text-transform: uppercase;
}

sf-domains,
sf-repositories,
sf-proposals,
sf-tasks {
  margin-bottom: 2.1rem;
  padding-top: 1.25rem;
  border-top: 2px solid #050505;
}

sf-domains::before,
sf-repositories::before,
sf-proposals::before,
sf-tasks::before {
  display: block;
  color: #050505;
  font-family: "SFMono-Regular", "Roboto Mono", "Menlo", monospace;
  font-size: clamp(1.55rem, 3vw, 2.35rem);
  font-weight: 900;
  letter-spacing: -0.06em;
  line-height: 1;
  margin-bottom: 1rem;
  text-transform: uppercase;
}

sf-domains::before {
  content: "Domains";
}

sf-repositories::before {
  content: "Repositories";
}

sf-proposals::before {
  content: "Proposals";
}

sf-tasks::before {
  content: "Tasks";
}

@media (max-width: 720px) {
  body {
    padding: 2.5rem 1.2rem;
    font-size: 16px;
  }

  sf-domains,
  sf-repositories,
  sf-proposals,
  sf-tasks {
    margin-bottom: 1.7rem;
    padding-top: 1rem;
  }
}
`
