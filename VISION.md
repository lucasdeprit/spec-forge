# Vision

Modern AI-assisted development suffers from a recurring problem:

Every agent repeatedly rediscovers the same codebase.

Developers spend time explaining architecture, module boundaries, repository roles and product context instead of focusing on intent.

SpecForge exists to solve this problem.

SpecForge is a multi-repository Specification Driven Development workspace based on semantic HTML documents that are readable by humans and AI agents.

The core principle is:

The agent should never have to rediscover the project.

## Product Philosophy

SpecForge does not generate intelligence.

SpecForge structures knowledge.

It separates software change into clear artifacts:

- Domain: the business area
- Repository: the technical execution unit
- Codebase Map: what exists
- Proposal: what should exist
- Task: how the problem is divided
- Implementation: how the code changes in a specific repository

## Why HTML

SpecForge uses semantic HTML as a single source of truth.

HTML gives the project:

- explicit structure
- browser-native readability
- CSS-based presentation
- agent-readable semantics
- no need for React, MDX or duplicated JSON

The same document can be read by a human in a browser and by an agent as structured semantic content.

## Enterprise Direction

SpecForge is designed for multi-repository systems.

Large products are usually organized around business domains and implemented across multiple repositories.

SpecForge models this directly:

- Domains represent business.
- Repositories represent execution.
- Proposals and Tasks are global.
- Codebase Maps and Implementations are repository-specific.
