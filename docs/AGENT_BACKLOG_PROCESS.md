# Agent Backlog Completion Process

## Core Principle

Agent proposes diffs.
Harness promotes diffs.
Git records truth.
Tests prove completion.
Handover preserves context.

The harness owns workflow. Agents are replaceable.

## Workflow

Backlog Item
→ Context Pack
→ Plan
→ Branch
→ Implement
→ Test
→ Commit
→ Handover Artifact
→ Backlog Update
→ Review
→ Complete

## Required Artifacts

For every backlog item:

- Context Pack
- Test Evidence
- Handover Artifact
- Implementation Commit
- Backlog Update

## Completion Gate

An item may only be marked complete when:

- Acceptance criteria are satisfied.
- Required tests pass.
- Evidence is recorded.
- Implementation commit exists.
- Handover artifact exists.

If any are missing, the item remains in-progress, review, or blocked.

## Status Flow

todo
→ selected
→ context-ready
→ in-progress
→ review
→ complete

## Handover Template

### Summary

### Files Changed

### Tests Run

### Decisions Made

### Risks

### Follow-up Work
