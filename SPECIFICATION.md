# Specification

## Artifact Model

Codebase Map
-> What exists

Proposal
-> What should exist

Tasks
-> How the problem is divided

Implementation
-> How the code changes

## Flow

Codebase Map
    -> Proposal
    -> Tasks
    -> Implementation

## Codebase Map

Purpose:

- Describe the current repository.
- Act as the navigational truth of the repository.
- Remain a living document.

Rules:

- One active map per repository.
- No versioned map history.
- Updated through init, map, and sync.

## Proposal

Purpose:

- Describe the desired outcome.
- Remain implementation agnostic.

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

## Tasks

Purpose:

- Divide a Proposal into reasoning units.

Rules:

- Technology agnostic.
- Repository agnostic.
- No implementation details.

States:

- todo
- doing
- blocked
- done

## Implementation

Purpose:

- Apply Tasks to a specific repository.

Rules:

- Repository specific.
- Technology aware.
- References the Codebase Map.

## IDs

Proposal:

PROP-YYYYMMDD-XXXX-slug

Task:

TASK-YYYYMMDD-XXXX-slug

Implementation:

IMPL-YYYYMMDD-XXXX-slug
