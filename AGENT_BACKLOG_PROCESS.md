# Agent Backlog Completion Process

This document defines how an implementation agent should work from the repository backlog, prove completion, and mark items complete with commit and test evidence.

The goal is to make backlog completion deterministic, reviewable, and auditable. An agent must not claim an item is complete unless the repository contains evidence showing what changed, which commit contains the change, and which validation commands passed.

## Required Files

Recommended repo-local files:

- `docs/BACKLOG.md` - human-readable backlog grouped by milestone or priority.
- `docs/AGENT_BACKLOG_PROCESS.md` - this process.
- `docs/decisions/` - optional architecture decision records for meaningful product or technical decisions.

If `docs/BACKLOG.md` does not exist yet, create it before marking any roadmap item complete.

## Backlog Item Format

Every backlog item should use this structure:

```md
## BLG-001: Short item title

- Status: todo | in-progress | blocked | review | complete
- Priority: P0 | P1 | P2 | P3
- Owner: agent | human | unassigned
- Area: seo | consent | blog | notifications | payments | subscriptions | communities | forms | platform
- Source: ROADMAP.md#section-name
- Created: YYYY-MM-DD
- Completed: YYYY-MM-DD or blank
- Completed Commit: commit SHA or blank
- Backlog Update Commit: commit SHA or blank
- Pull Request: PR URL or blank

### Problem

What user or platform problem this item solves.

### Scope

What is included.

### Out of Scope

What is intentionally not included.

### Acceptance Criteria

- [ ] Observable outcome 1.
- [ ] Observable outcome 2.
- [ ] Observable outcome 3.

### Required Tests

- [ ] `cd Frontend && npm run build`
- [ ] `cd Frontend && npm run test:run`
- [ ] `cd Backend && go test ./...`
- [ ] Add/update feature-specific tests where applicable.

### Evidence

- Changed files:
  - `path/to/file`
- Test evidence:
  - Command: `...`
  - Result: pass/fail
  - Notes: short notes
- Review notes:
  - Any risks, follow-ups, or known limitations.
```

## Agent Workflow

### 1. Select a Backlog Item

The agent must select exactly one backlog item unless the human explicitly asks for a batch.

Selection rules:

1. Prefer `Status: todo` items with the highest priority.
2. Do not start blocked items.
3. Do not start items without acceptance criteria.
4. Do not start items where the scope is ambiguous unless the ambiguity can be resolved from existing repo docs.
5. Mark the selected item as `Status: in-progress` before implementation if the agent is allowed to commit process updates.

### 2. Build a Context Pack

Before editing code, the agent should identify:

- Relevant roadmap section.
- Relevant existing backend routes, stores, middleware, and tests.
- Relevant frontend pages, API clients, auth hooks, and tests.
- Existing data model or collection conventions.
- Required environment variables.
- Expected validation commands.

The agent should keep this context small and directly related to the backlog item.

### 3. Implement the Change

Implementation rules:

1. Make the smallest coherent change that satisfies the acceptance criteria.
2. Keep auth, role checks, and entitlement checks explicit.
3. Prefer reusable backend and frontend patterns already used in the repo.
4. Add migrations or data-shape notes when persistent data changes.
5. Add tests in the same commit where practical.
6. Do not mix unrelated backlog items into the same implementation commit.

### 4. Run Required Tests

The standard validation commands are:

```bash
cd Frontend && npm run build
cd Frontend && npm run test:run
cd Backend && go test ./...
```

For changes involving live infrastructure, also run when the services are available:

```bash
cd Backend && go test -v -tags=liveintegration ./integration
```

Test rules:

1. If frontend code changed, run frontend build and frontend tests.
2. If backend code changed, run backend tests.
3. If Docker, Keycloak, MongoDB, MinIO, auth, or integration behavior changed, run or explicitly defer live integration tests with a reason.
4. If only documentation changed, record `Documentation-only change; tests not required` in evidence.
5. Failed tests block completion.
6. Skipped tests must include a reason and the item must remain `review` or `blocked`, not `complete`, unless the human approves the skip.

### 5. Commit the Change

Commit message format:

```text
Complete BLG-001 short item title
```

For partial work:

```text
Progress BLG-001 short item title
```

For documentation/process work:

```text
Document BLG-001 short item title
```

The final completion commit SHA is the source of truth for the completed backlog item.

### 6. Update the Backlog Item

After the implementation commit exists and tests pass, update the backlog item:

```md
- Status: complete
- Completed: YYYY-MM-DD
- Completed Commit: abc123fullcommithash
- Backlog Update Commit: def456fullcommithash
```

Update acceptance criteria checkboxes:

```md
### Acceptance Criteria

- [x] Observable outcome 1.
- [x] Observable outcome 2.
- [x] Observable outcome 3.
```

Update test evidence:

```md
### Evidence

- Changed files:
  - `Backend/cmd/api/routes.go`
  - `Frontend/src/app/pages/ExamplePage.tsx`
- Test evidence:
  - Command: `cd Frontend && npm run build`
    - Result: pass
  - Command: `cd Frontend && npm run test:run`
    - Result: pass
  - Command: `cd Backend && go test ./...`
    - Result: pass
- Completion commit: `abc123fullcommithash`
- Review notes:
  - No known follow-ups.
```

Then commit the backlog update:

```text
Mark BLG-001 complete
```

This creates a clean audit trail:

1. Implementation commit proves the change.
2. Test evidence proves validation.
3. Backlog update commit records completion.

## Completion Gate

A backlog item can only move to `complete` when all of these are true:

- The implementation exists in a committed change.
- The backlog item has a `Completed Commit` SHA.
- The backlog item has a `Backlog Update Commit` SHA.
- Acceptance criteria are checked off.
- Required tests are listed with pass results, or a documentation-only exemption is recorded.
- Any skipped validation has a clear reason and human approval.
- Any new environment variables are documented.
- Any new public API behavior is documented.
- Any new user-facing behavior is represented in frontend/UI docs or screenshots where relevant.

If any of these are false, use `Status: review`, `Status: blocked`, or `Status: in-progress` instead of `complete`.

## Status Definitions

| Status | Meaning |
| --- | --- |
| `todo` | Ready to be selected. |
| `in-progress` | Agent or human is actively working on it. |
| `blocked` | Cannot continue without missing input, failing dependency, or unresolved decision. |
| `review` | Implementation exists but needs review, missing validation, or human sign-off. |
| `complete` | Merged/committed, tested, and recorded with completion evidence. |

## Commit Hash Rules

- Use the full commit SHA where possible.
- The `Completed Commit` should point to the implementation commit, not just the backlog metadata update commit.
- The `Backlog Update Commit` should point to the commit that marks the item complete in `docs/BACKLOG.md`.
- If multiple commits are required, list all commits in the evidence section and put the final implementation commit in `Completed Commit`.
- Never use an uncommitted diff as completion evidence.

## Test Evidence Rules

Good evidence:

```md
- Command: `cd Backend && go test ./...`
  - Result: pass
  - Notes: all backend packages passed locally
```

Weak evidence that should not be accepted for completion:

```md
- Tests: probably fine
- Tests: not run
- Tests: skipped
```

Allowed documentation-only evidence:

```md
- Command: not run
  - Result: not applicable
  - Notes: documentation-only change; no runtime code changed
```

## Agent Final Response Format

When an agent finishes an item, it should report:

```md
## Summary

- Completed BLG-001: item title.
- Main change: short summary.

## Files Changed

- `path/to/file`

## Validation

- `cd Frontend && npm run build` - pass
- `cd Frontend && npm run test:run` - pass
- `cd Backend && go test ./...` - pass

## Backlog Evidence

- Completed Commit: `abc123fullcommithash`
- Backlog update commit: `def456fullcommithash`
- Status: complete

## Follow-ups

- Any remaining work, or `None`.
```

## Example Completed Backlog Item

```md
## BLG-SEO-001: Add SEO metadata model to publishable pages

- Status: complete
- Priority: P0
- Owner: agent
- Area: seo
- Source: ROADMAP.md#1-seo-system
- Created: 2026-06-01
- Completed: 2026-06-02
- Completed Commit: 1111111111111111111111111111111111111111
- Pull Request:

### Problem

Public CMS pages need editable SEO metadata so published content can be indexed and shared correctly.

### Scope

Add a reusable SEO metadata shape to publishable pages and validate core fields.

### Out of Scope

Sitemap generation and social preview UI.

### Acceptance Criteria

- [x] Pages can store title, description, canonical URL, and robots settings.
- [x] Backend validates duplicate or invalid metadata where applicable.
- [x] API responses include SEO metadata for page resources.

### Required Tests

- [x] `cd Backend && go test ./...`

### Evidence

- Changed files:
  - `Backend/cmd/api/store.go`
  - `Backend/cmd/api/routes.go`
  - `Backend/cmd/api/routes_test.go`
- Test evidence:
  - Command: `cd Backend && go test ./...`
    - Result: pass
- Completion commit: `1111111111111111111111111111111111111111`
- Review notes:
  - Sitemap work remains in BLG-SEO-002.
```

## Handling Blocked Items

If the agent cannot complete an item, update the status to `blocked` and record:

- What blocked the work.
- What was attempted.
- What decision or dependency is required.
- Whether any partial commits exist.
- Which tests were run.

Do not mark blocked work as complete.

## Handling Partial Work

If useful partial work was committed but acceptance criteria are not complete:

- Keep status as `in-progress` or `review`.
- Add the partial commit hash to evidence.
- Leave `Completed Commit` blank.
- List remaining acceptance criteria unchecked.

## Human Review Checklist

Before accepting an agent-completed item, the reviewer should confirm:

- The completed commit exists in the repository.
- The diff matches the backlog item scope.
- Tests listed in evidence are appropriate for the changed files.
- Acceptance criteria are actually satisfied.
- No unrelated scope was bundled into the completion commit.
- Follow-up items were created for deferred work.
