# GitHub Events API

`gh-activity` uses the public GitHub Events API to fetch recent activity for a user.

## Endpoint

```
GET https://api.github.com/users/<username>/events?per_page=100
```

- Returns the **most recent** events (not full history).
- `per_page=100` fetches up to 100 events in one call.
- Pagination beyond 100 is currently not implemented.

## Event Types Used

- `PushEvent` — pushes to repositories
- `PullRequestEvent` — pull request opened/closed/merged
- `IssuesEvent` — issue opened/closed
- `WatchEvent` — repository starred

Other event types exist but are ignored for current commands.

## Data Mapping

- JSON is decoded into `internal/models/event_model.go` types:
  - `GitResponseObject` (root)
  - `ActorModel`, `RepoModel`, `PayloadModel`
- Key fields used:
  - `Type`, `Repo.Name`, `Payload.Action`, `CreatedAt`

## Rate Limiting & Auth

- Unauthenticated requests have stricter rate limits.
- If you set a Personal Access Token (PAT), requests include `Authorization: Bearer <token>`.
- Invalid or expired tokens trigger a graceful recovery: the token is removed and the request is retried unauthenticated (see `internal/github/events.go`).

## Known Behaviors

- One `PushEvent` may include multiple commits — counts reflect events, not commits.
- Events are time-ordered; the tool does not yet filter by date ranges, only by recent window and `--limit` semantics.