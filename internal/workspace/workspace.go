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
  padding: 3rem 1.25rem;
  background:
    radial-gradient(circle at 20% -10%, rgba(255, 255, 255, 0.82), transparent 32rem),
    linear-gradient(135deg, #faf4ea 0%, #f5edde 50%, #eee3d3 100%);
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
  padding: clamp(1.4rem, 3.4vw, 2.6rem);
  background: rgba(255, 253, 248, 0.92);
  border: 1px solid rgba(77, 57, 32, 0.16);
  border-radius: 14px;
  box-shadow: 0 18px 55px rgba(84, 62, 35, 0.09);
}

sf-title {
  color: #191714;
  font-size: clamp(2.35rem, 5vw, 4.2rem);
  font-weight: 820;
  letter-spacing: -0.065em;
  line-height: 0.95;
  margin-bottom: 1.85rem;
}

sf-domains,
sf-repositories,
sf-proposals,
sf-tasks {
  margin-bottom: 1rem;
  padding: 1.15rem;
  background: rgba(255, 255, 255, 0.56);
  border: 1px solid #d8c8b4;
  border-radius: 10px;
  box-shadow: 0 8px 24px rgba(91, 66, 36, 0.035);
}

sf-domains::before,
sf-repositories::before,
sf-proposals::before,
sf-tasks::before {
  display: block;
  margin-bottom: 0.85rem;
  padding-bottom: 0.65rem;
  border-bottom: 1px solid #d8c8b4;
  color: #191714;
  font-family: "SFMono-Regular", "Roboto Mono", Menlo, monospace;
  font-size: 0.72rem;
  font-weight: 780;
  letter-spacing: 0.14em;
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
    padding: 0.9rem;
    font-size: 16px;
  }

  sf-workspace {
    border-radius: 12px;
    padding: 1.15rem;
  }
}
`
