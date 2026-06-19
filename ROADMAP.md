# Roadmap

## v0.1

CLI-first multi-repository workspace.

Goals:

- initialize a SpecForge workspace
- register repositories
- generate living Codebase Maps
- create global Proposals
- create global Tasks from Proposals
- create repository-specific Implementations from Tasks
- validate semantic HTML artifacts

Commands:

```bash
specforge init
specforge repo add <id> <path>
specforge map <repo-id>
specforge proposal
specforge tasks <proposal-id>
specforge implementation <task-id> --repo <repo-id>
specforge sync <repo-id>
specforge validate
```

Deliverables:

- workspace.html schema
- codebase-map.html schema
- proposal.html schema
- task.html schema
- implementation.html schema
- global CSS
- example artifacts

Non-goals:

- web UI
- hosted service
- Jira integration
- GitHub integration
- React runtime
- MDX runtime

## v0.2

- better domain detection
- richer validation
- improved visual themes
- optional context extraction for agents
- workspace index generation

## v1.0

- stable semantic HTML specification
- plugin system
- enterprise workflow support
