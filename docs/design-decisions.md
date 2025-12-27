# Design Decisions

Key choices made in `gh-activity` and why.

## Separation of Concerns

- **Handlers**: own user-facing formatting and orchestrate calls.
- **Services**: own business logic and aggregation; pure functions over event data.
- **Models**: encapsulate JSON decoding types.

This separation keeps logic testable and output consistent.

## Output Formatting

- Use simple alignment with formatting verbs and `strings.Repeat`.
- For multi-column summaries, prefer Go `tabwriter` for consistent spacing across terminals.

## Limits Semantics

- Current `--limit`: cap distinct repositories per event type.
- Rationale: avoids overly long lists and keeps output focused.
- Trade-off: not a strict “last N events”. A future global cap can be implemented easily.

## Error Handling

- `customerror.Wrap(context, err)` standardizes messages.
- Clear messages for invalid flags/scopes/commands.

## Token Management

- Store token with timestamp; warn if older than 90 days.
- On `Bad credentials`, remove token and retry unauthenticated to unblock users.

## Deterministic Ordering

- Sort outputs by count desc, then repo name asc for stable, readable results.

## Simplicity Over Features

- Use the REST Events API with a single page (up to `per_page=100`).
- Avoid pagination until needed; keep the CLI fast and predictable.

