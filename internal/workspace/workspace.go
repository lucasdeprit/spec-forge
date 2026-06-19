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
    radial-gradient(circle at 18% 0%, rgba(255, 255, 255, 0.72), transparent 30rem),
    linear-gradient(135deg, #f8f1e6 0%, #f4ecdf 48%, #eee2d1 100%);
  color: #191714;
  font-family: Inter, -apple-system, BlinkMacSystemFont, "Helvetica Neue", "Segoe UI", sans-serif;
  font-size: 17px;
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
  width: min(960px, 100%);
  margin: 0 auto;
  padding: clamp(1.5rem, 4vw, 3rem);
  background: rgba(255, 250, 241, 0.88);
  border: 1px solid rgba(82, 63, 40, 0.14);
  border-radius: 32px;
  box-shadow: 0 28px 80px rgba(83, 61, 34, 0.12);
}

sf-title {
  color: #191714;
  font-size: clamp(2.5rem, 7vw, 5rem);
  font-weight: 850;
  letter-spacing: -0.075em;
  line-height: 0.94;
  margin-bottom: 2rem;
}

sf-domains,
sf-repositories,
sf-proposals,
sf-tasks {
  margin-bottom: 1rem;
  padding: 1.15rem;
  background: rgba(255, 255, 255, 0.48);
  border: 1px solid #ded2c0;
  border-radius: 22px;
}

sf-domains::before,
sf-repositories::before,
sf-proposals::before,
sf-tasks::before {
  display: block;
  margin-bottom: 0.85rem;
  padding-bottom: 0.65rem;
  border-bottom: 1px solid #ded2c0;
  color: #191714;
  font-family: "SFMono-Regular", "Roboto Mono", Menlo, monospace;
  font-size: 0.82rem;
  font-weight: 760;
  letter-spacing: 0.1em;
  line-height: 1.15;
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
    padding: 1rem;
    font-size: 16px;
  }

  sf-workspace {
    border-radius: 24px;
  }
}
`
