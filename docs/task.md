# task.html

`task.html` is a global SpecForge artifact.

A Task belongs to one Proposal.

A Task describes how the problem is divided logically.

A Task is not implementation-specific.

It should not mention:

- repositories
- files
- classes
- concrete framework details

Those details belong to Implementation documents.

## Required Sections

- sf-task
- sf-title
- sf-description
- sf-purpose
- sf-status
- sf-proposal-ref

## Optional Sections

- sf-domains
- sf-dependencies
- sf-expected-output
- sf-validation-notes

## Status

Allowed statuses:

- todo
- doing
- blocked
- done

## Rules

- A Task belongs to exactly one Proposal.
- A Task may reference one or more Domains.
- A Task may depend on other Tasks.
- A Task is a reasoning unit, not a ticket and not an implementation plan.
