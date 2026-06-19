# workspace.html

`workspace.html` is the root document of a SpecForge workspace.

It defines:

- workspace identity
- domains
- repositories
- repository paths
- high-level repository roles

It does not describe code internals. Code internals belong to each repository's `codebase-map.html`.

## Required Sections

- sf-workspace
- sf-name
- sf-domains
- sf-repositories

## Optional Sections

- sf-description
- sf-owner
- sf-tags

## Rules

- Domains are business concepts.
- Repositories are technical execution units.
- Repositories are referenced by path.
- Repositories do not live inside the SpecForge workspace.
- For v0.1, domains live only inside workspace.html.
