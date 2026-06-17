# Agent-MD Backlog

## BLG-001: Initialize Agent-MD Repository Structure

- Status: todo
- Priority: P0
- Owner: agent
- Area: platform
- Created: 2026-06-17

### Problem

The harness requires a consistent directory structure and configuration model before workflow automation can occur.

### Acceptance Criteria

- [ ] agent-md init creates .agent-md/
- [ ] agent-md init creates .agent-md/runs/
- [ ] agent-md init creates .agent-md/config.yaml
- [ ] running twice does not overwrite config
- [ ] go test ./... passes

### Required Tests

- [ ] go test ./...
