# SpecForge

SpecForge is a multi-repository Specification Driven Development workspace for humans and AI agents.

SpecForge does not generate intelligence.

SpecForge structures knowledge.

## Core Idea

The agent should never have to rediscover the project.

SpecForge separates business intent from technical execution through a workspace that manages domains, repositories, proposals, tasks and implementations.

## Conceptual Model

- Workspace -> the ecosystem
- Domain -> business area
- Repository -> technical execution unit
- Codebase Map -> what exists in a repository
- Proposal -> what should exist
- Task -> how the problem is divided
- Implementation -> how a task is applied to a repository

## Main Flow

Codebase Map
-> Proposal
-> Tasks
-> Implementation

## Global vs Repository-Specific

Global workspace artifacts:

- Domains
- Proposals
- Tasks

Repository-specific artifacts:

- Codebase Maps
- Implementations

## Source of Truth

SpecForge uses semantic HTML as its source of truth.

Presentation is handled through external CSS.

Rules:

- No inline CSS
- No inline JavaScript
- No mandatory MDX
- No mandatory React
- No persistent JSON duplicate
- No multiple sources of truth

## Workspace Layout

SpecForge lives outside the repositories it documents.

Example:

```txt
company/
├── specforge/
│   └── .specforge/
├── ios-app/
├── backend-api/
└── web-admin/
```

## Initial CLI Goals

```bash
specforge init
specforge repo add <id> <path>
specforge map <repo-id>
specforge proposal
specforge tasks <proposal-id>
specforge implementation <task-id> --repo <repo-id>
specforge sync <repo-id>
```
