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
  padding: 4rem 1.5rem;
  background:
    radial-gradient(circle at top left, rgba(255, 255, 255, 0.92), transparent 34rem),
    linear-gradient(135deg, #f5f5f7 0%, #eef1f5 46%, #f8f9fb 100%);
  color: #1d1d1f;
  font-family: -apple-system, BlinkMacSystemFont, "SF Pro Display", "Segoe UI", sans-serif;
  font-size: 16px;
  line-height: 1.55;
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
  max-width: 980px;
  margin: 0 auto;
}

sf-title {
  color: #111113;
  font-size: clamp(2.25rem, 6vw, 4.75rem);
  font-weight: 700;
  letter-spacing: -0.055em;
  line-height: 0.95;
  margin-bottom: 1.5rem;
}

sf-domains,
sf-repositories,
sf-proposals,
sf-tasks {
  margin-bottom: 1rem;
  padding: 1.15rem 1.25rem;
  background: rgba(255, 255, 255, 0.82);
  border: 1px solid rgba(0, 0, 0, 0.08);
  border-radius: 20px;
  box-shadow: 0 18px 50px rgba(0, 0, 0, 0.06);
  backdrop-filter: blur(18px);
}

sf-domains::before,
sf-repositories::before,
sf-proposals::before,
sf-tasks::before {
  display: block;
  color: #1d1d1f;
  font-size: 0.78rem;
  font-weight: 700;
  letter-spacing: 0.08em;
  margin-bottom: 0.5rem;
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
    padding: 2rem 1rem;
  }

  sf-domains,
  sf-repositories,
  sf-proposals,
  sf-tasks {
    border-radius: 16px;
    padding: 1rem;
  }
}
`
