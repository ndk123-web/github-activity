# CLI Grammar

This document defines the command-line structure for `gh-activity`.

## Syntax

```
gh-activity <scope> <entity> <command> [flags]
```

### Grammar (EBNF-like)

```
scope      ::= "user" | "set" | "get"
entity     ::= <username> | "token"
command    ::= command-user | command-set | command-get
command-user ::= "pushes" | "pulls" | "issues" | "watches" | "summary"
command-set  ::= (entity == "token")
command-get  ::= (entity == "token")
flags      ::= (limit-flag)? (state-flag)?
limit-flag ::= "--limit" <int>
state-flag ::= "--state" ("open" | "closed" | "merged")
```

## Examples

```powershell
# Show pushes grouped by repository
gh-activity user octocat pushes --limit 10

# Pull requests (state required)
gh-activity user octocat pulls --state open --limit 5

# Issues (state required)
gh-activity user octocat issues --state closed --limit 5

# Watches grouped by repository
gh-activity user octocat watches --limit 5

# Combined summary
gh-activity user octocat summary --limit 10

# Token management
gh-activity set token <your_token>
gh-activity get token
```

## Notes

- `--limit` currently caps distinct repositories per event type (pushes/pulls/issues/watches), not the global number of events.
- `--state` is mandatory for `pulls` and `issues`.
- Scope, command, and flags are validated against rules in `internal/models/rules.go` and helpers in `internal/github/valid_scope.go`.