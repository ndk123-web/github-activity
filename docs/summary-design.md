# Summary Command Design

This doc explains how the `summary` command aggregates and prints data.

## Goal

Provide a consolidated view of recent activity (pushes, pulls, issues, watches) with:
- A metrics table (Metric/Count)
- An aggregated list of top repositories by total activity

## Data Flow

1. `cmd/main.go` routes `summary` to `internal/handlers/summary_all.go`.
2. Handler calls `internal/services/summary_all.go::GetAllSummary(limit)`.
3. Service returns `map[string]map[string]int64` keyed by event type:
   - `pushEvents`, `pullEvents`, `issueEvents`, `watchEvents`
   - Inner map: `repo -> count`
4. Handler:
   - Sums repo counts per event type to compute accurate totals
   - Prints a tabular metrics section using Go `tabwriter`
   - Aggregates totals across all event types per repo and prints a sorted list

## Output Example

```
📊 Activity Summary (last 10 events)

Metric                     Count
Total Events               54
Total Push Events          45
Total Pull Request Events  2
Total Issues Events        2
Total Watch Events         5

Top Repositories (by total activity):
------------------------------------------
ndk123-web/DSA                 21 events
ndk123-web/github-activity     20 events
ndk123-web/observability-learning 4 events
ndk123-web/study-sync-ai       4 events
spf13/cobra                    1 events
```

## `--limit` Semantics (current)

- The service caps the number of **distinct repositories per event type** using `IsGreaterThanLimit(len(typeMap), limit)`.
- This is **not** a global cap on last N events.

### Alternative (global event cap)

- Replace per-type `len(map)` checks with a global counter `cnt` in the service.
- Stop processing when `cnt >= limit`.
- Pros: aligns with typical expectations of `--limit`.
- Cons: may skew distribution across event types depending on chronology.

## Formatting Choices

- Metrics section uses `tabwriter` to keep columns aligned in various terminals.
- Aggregated repos printed with fixed-width left column for readability.
- Output is deterministic: entries sorted by count desc, then name asc.