# SpecForge

SpecForge is a CLI-first Specification Driven Development (SDD) system built for humans and AI agents.

The core idea is simple:

1. Build a living Codebase Map for each repository.
2. Define a Proposal describing the desired outcome.
3. Break the Proposal into Tasks.
4. Create repository-specific Implementations.

SpecForge does not generate intelligence.

SpecForge structures knowledge.

## Core Artifacts

- Codebase Map -> what exists
- Proposal -> what should exist
- Tasks -> how the problem is divided
- Implementation -> how the code changes

## Principles

- The agent should never have to rediscover the project.
- Documentation should be living, not forgotten.
- Product intent should be separated from implementation details.
- Repository knowledge should remain local to each repository.
- Proposals should describe outcomes, not technical solutions.

## Initial Workflow

specforge init
-> Creates or updates the Codebase Map

specforge proposal
-> Creates a Proposal

specforge tasks
-> Creates Tasks from a Proposal

specforge implementation
-> Creates repository-specific implementation plans

specforge sync
-> Updates repository knowledge after changes
