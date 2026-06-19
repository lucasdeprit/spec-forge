# implementation.html

`implementation.html` is a repository-specific SpecForge artifact.

An Implementation belongs to one Task and one Repository.

It describes how a Task should be applied to a concrete repository.

This is the first artifact that can mention:

- repository paths
- files
- classes
- functions
- tests
- commands
- framework-specific details

## Required Sections

- sf-implementation
- sf-title
- sf-task-ref
- sf-repository-ref
- sf-summary
- sf-change-plan
- sf-validation

## Optional Sections

- sf-files
- sf-commands
- sf-risks
- sf-notes

## Rules

- An Implementation belongs to exactly one Task.
- An Implementation belongs to exactly one Repository.
- An Implementation should reference the repository Codebase Map.
- An Implementation can be regenerated or updated as the repository changes.
- Implementation documents are execution-oriented, not product-intent documents.
