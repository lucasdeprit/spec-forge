# Specification

## Artifact Model

SpecForge has the following core artifacts:

- Workspace
- Domain
- Repository
- Codebase Map
- Proposal
- Task
- Implementation

## Responsibilities

### Workspace

The root of a SpecForge system.

It defines:

- workspace identity
- domains
- repositories
- global configuration

### Domain

A business area.

Examples:

- Authentication
- Profile
- Cards
- Payments
- Notifications
- Jobs

For v0.1, domains exist inside workspace.html. They are not separate documents.

### Repository

A technical execution unit.

Examples:

- ios-app
- backend-api
- web-admin
- infra

Repositories are not stored inside the SpecForge workspace. They are referenced by path.

### Codebase Map

A living repository-specific document.

Purpose:

- show what exists
- show where it is
- show how repository parts relate

Contains:

- stack
- domains
- paths
- file names
- dependencies between domains

Does not contain:

- agent instructions
- naming conventions
- work rules

Those belong in AGENTS.md or equivalent agent instructions.

### Proposal

A global document describing what should exist.

A Proposal is implementation agnostic.

Fields:

- title
- description
- goal
- requirements
- constraints
- out-of-scope
- acceptance-criteria
- status

States:

- draft
- ready
- completed
- archived

### Task

A global document belonging to one Proposal.

A Task divides the problem logically.

A Task is:

- technology agnostic
- repository agnostic
- implementation agnostic

States:

- todo
- doing
- blocked
- done

### Implementation

A repository-specific document belonging to one Task and one Repository.

An Implementation describes how to apply a Task in a concrete repository.

It can reference:

- files
- directories
- classes
- functions
- modules
- tests
- commands

## Relationships

- A Workspace has many Domains.
- A Workspace has many Repositories.
- A Repository has one living Codebase Map.
- A Proposal may reference one or more Domains.
- A Task belongs to one Proposal.
- An Implementation belongs to one Task and one Repository.

## IDs

Proposal:

```txt
PROP-YYYYMMDD-XXXX-slug
```

Task:

```txt
TASK-YYYYMMDD-XXXX-slug
```

Implementation:

```txt
IMPL-YYYYMMDD-XXXX-slug
```

Repository IDs and Domain IDs should be stable slugs.

Examples:

```txt
ios-app
backend-api
authentication
profile
payments
```

## Source of Truth

The source of truth is semantic HTML.

CSS is external.

Rules:

- no inline CSS
- no inline JavaScript
- no persistent duplicated JSON
- no mandatory MDX
- no mandatory React


## Artifact Documents

The v0.1 semantic HTML documents are:

- workspace.html
- codebase-map.html
- proposal.html
- task.html
- implementation.html
