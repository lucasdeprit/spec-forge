# codebase-map.html

`codebase-map.html` is a living repository-specific map.

It answers:

- what exists in a repository
- where it is
- what stack is used
- which business domains appear in the repository
- how those domains relate to each other

It is not an agent instruction file.

Agent behavior belongs in `AGENTS.md` or equivalent instructions.

## Required Sections

- sf-codebase-map
- sf-repository-ref
- sf-stack
- sf-structure
- sf-domains

## Optional Sections

- sf-domain-dependencies
- sf-entrypoints
- sf-generated-metadata

## Rules

- One living Codebase Map per registered repository.
- The map references a repository declared in `workspace.html`.
- The map references domains declared in `workspace.html`.
- The map is generated and updated by SpecForge.
- Human comments may be preserved in explicit non-generated sections.
- The map should avoid reading or storing full source file contents.
- The map should store folder paths and relevant file names.
- The map should ignore noisy directories such as `.git`, `node_modules`, `dist`, `build`, `target`, `.next`, `DerivedData`, and similar generated folders.

## Stack

The stack should include:

- languages
- frameworks
- package managers

Tools may be added later, but v0.1 should stay lightweight.

## Structure

The structure should be filtered.

Default depth: 3.

Depth should be configurable.

## Domains

Domains represent business areas.

A domain in the map should reference a domain from `workspace.html`.

Example:

```html
<sf-domain ref="authentication" confidence="high">
```

A domain can have multiple paths inside the same repository.

## Dependencies

Domain dependencies should be directional.

Example:

```html
<sf-dependency from="profile" to="authentication" confidence="medium" />
```

Confidence should be used because dependencies may be inferred.
