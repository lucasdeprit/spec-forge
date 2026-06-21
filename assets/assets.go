package assets

import "embed"

// Files contains the default SpecForge workspace files copied by `specforge init`.
//
//go:embed workspace.html specforge.css
var Files embed.FS
