# Concepts

## Workspace

The root SpecForge artifact.

A workspace represents a product, platform or system composed of one or more repositories.

## Domain

A domain represents business.

Examples:

- Authentication
- Profile
- Cards
- Payments
- Jobs

For v0.1, domains are declared inside workspace.html.

## Repository

A repository represents execution.

Repositories are referenced by path and are not stored inside the SpecForge workspace.

## Codebase Map

A living map of a repository.

It answers:

- what exists
- where it is
- how domains relate inside this repository

## Proposal

A global description of the desired outcome.

It answers:

- what should exist
- why
- what requirements must be met
- what is out of scope
- how success is validated

## Task

A global logical division of a Proposal.

It does not know technology or repositories.

## Implementation

A repository-specific application of a Task.

It is the first artifact that knows concrete files, classes, commands and tests.
