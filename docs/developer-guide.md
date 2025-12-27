# Developer Guide

This document explains how `gh-activity` is structured and how data flows through the CLI, so you can extend or maintain it confidently.

---

## Overview

- Language: Go
- Entry point: `cmd/main.go`
- Internal modules under `internal/`
  - `config/`: Token storage and config-file helpers
  - `github/`: HTTP client for GitHub Events API + validation helpers
  - `handlers/`: CLI-facing formatting and orchestration per command
  - `services/`: Business logic and aggregation per event type
  - `models/`: Typed DTOs for JSON decoding and shared types
  - `custom-error/`: Lightweight error wrapper for consistent messages

The CLI fetches recent GitHub user events, transforms them into summaries per command, and prints readable, script-friendly output.

---

## Request Flow

1. `cmd/main.go` parses args: `<scope> <username> <command> [flags]`.
2. Validates scope/command via `internal/github/valid_scope.go` and rules from `internal/models/rules.go`.
3. Builds user events URL: `https://api.github.com/users/<username>/events?per_page=100`.
4. Calls `internal/github/events.go::FetchGitHubApiData()` to retrieve and decode events:
   - Adds `Authorization: Bearer <token>` header if a token is available.
   - Handles errors and specific GitHub responses (e.g., bad credentials) gracefully.
5. Dispatches to a `handler` based on command:
   - `pushes` → `internal/handlers/push_event.go`
   - `pulls` → `internal/handlers/pull_event.go`
   - `issues` → `internal/handlers/issue_event.go`
   - `watches` → `internal/handlers/watch_event.go`
   - `summary` → `internal/handlers/summary_all.go`
6. Handlers delegate business logic to `services/*` and print formatted output.

---

## Modules

### `internal/config`
- `gh-token.go`: Stores and loads a personal access token in a user config file (JSON). Implements age warnings (older than 90 days).
- `config-path.go`: Resolves platform-specific config file path and RW helpers.

### `internal/github`
- `events.go`: HTTP GET to GitHub Events API, auth header injection, error parsing and retry logic (removes invalid token and retries once unauthenticated).
- `valid_scope.go`: Simple validators using `slices.Contains`.

### `internal/models`
- `event_model.go`: Structs for decoding GitHub Events API JSON (`GitResponseObject`, `ActorModel`, `RepoModel`, etc.).
- `rules.go`: Contains CLI rules mapping commands to valid flags.

### `internal/services`
- Encapsulate event filtering and aggregation logic.
- `push_event.go`, `pull_event.go`, `issue_event.go`, `watch_event.go`: Provide methods like `GetPushEventsRepoWise(limit)` etc.
- `summary_all.go`: Builds a map of event-type → repo → count. Note on `--limit` semantics below.

### `internal/handlers`
- Responsible for user-facing output formatting.
- Each handler prints a small overview and then a table-like output aligned with spaces (`strings.Repeat` and formatting verbs). The summary handler uses Go’s `tabwriter` for clean column alignment.

---

## Summary Command Details

- Entry: `internal/handlers/summary_all.go::GetAllSummary(limit, jsonData)`
- Service: `internal/services/summary_all.go::GetAllSummary(limit)` → returns `map[string]map[string]int64` keyed by event type (`pushEvents`, `pullEvents`, `issueEvents`, `watchEvents`).
- Handler computes totals by summing counts per repo for each event type, prints:
  - A metrics table (Metric/Count) using `tabwriter`
  - An aggregated "Top Repositories (by total activity)" list combining all event types, sorted by total desc

### `--limit` Semantics (current)
- `summary_all.go` service uses `IsGreaterThanLimit(len(typeMap), limit)` to cap the number of distinct repositories per event type.
- This means `--limit N` restricts how many unique repos (not total events) are included for each event type.
- If you want `--limit` to mean “process only the last N events overall”: change the service loop to stop on a global counter (e.g., `cnt >= limit`) instead of checking `len(map)` per type.

---

## Adding a New Command

1. Define the command in `internal/models/rules.go` (valid flags and scope mapping).
2. Add handler in `internal/handlers/<your_command>.go` with an interface and constructor.
3. Implement service functions in `internal/services/<your_command>.go` that take `[]models.GitResponseObject` and return structured results.
4. Wire it in `cmd/main.go`:
   - Validate flags
   - Instantiate handler and call methods
5. Print output aligned consistently (reuse the style from other handlers; prefer `tabwriter` for multi-column tables).

---

## Error Handling & Tokens

- Use `customerror.Wrap(context, err)` to wrap messages consistently.
- Token management:
  - Set: `gh-activity set token <token>` → writes JSON with timestamp.
  - Get: `gh-activity get token` → prints token and warns if missing/old.
- `events.go` retries unauthenticated if token is bad (and deletes stored token to avoid repeated failures).

---

## Build & Run

- Local run (requires Go):

```powershell
# Windows
go run .\cmd\main.go user <username> summary --limit 10
```

- Typical commands:

```powershell
# Pushes
go run .\cmd\main.go user <username> pushes --limit 10

# Pull requests (state required)
go run .\cmd\main.go user <username> pulls --state open --limit 5

# Issues (state required)
go run .\cmd\main.go user <username> issues --state closed --limit 5

# Watches
go run .\cmd\main.go user <username> watches --limit 5
```

---

## Style & Conventions

- Keep handlers focused on formatting and user messages.
- Keep services pure and testable: input events → output aggregates.
- Use explicit types (`int64` for counters from services) to avoid mismatches.
- Prefer sorted output for deterministic results.
- Avoid changing public interfaces unless necessary; keep edits minimal and focused.

---

## Testing Ideas

- Unit tests for services (aggregation correctness, limit semantics).
- Handlers: table formatting snapshots (or simple alignment checks).
- HTTP: mock `FetchGitHubApiData` for predictable inputs.

---

## Packaging (Future)

- Ship platform-specific binaries in a `bin/` folder within release zips.
- Add a Makefile or GoReleaser config for reproducible builds.

---

## Future Enhancements

- Global `--limit` semantics option (last N events overall).
- Additional commands (e.g., comments, reviews, stars by repo).
- Configurable output styles (JSON/plain/table).
- Caching recent events to reduce API calls.

