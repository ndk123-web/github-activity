# GitHub Activity CLI Tool

A command-line tool to fetch and display GitHub user activity in a simple and organized way.

## What This Tool Does

This CLI tool helps you see what a GitHub user has been doing - their pushes, pull requests, and issues. You can filter and limit the results according to your needs.

## Features Implemented

### 1. **User Scope** (`user`)
Get information about a specific GitHub user's activity.

#### a) Push Events (`pushes`)
- Shows all push events by the user
- Displays pushes organized by repository
- Optional: Limit the number of results

**Example:**
```bash
gh-activity user username pushes
gh-activity user username pushes --limit 5
```

#### b) Pull Requests (`pulls`)
- Shows all pull request activity
- Filters by state (open, closed, merged)
- Displays PRs organized by repository
- Optional: Limit the number of results

**Example:**
```bash
gh-activity user username pulls --state open
gh-activity user username pulls --state merged --limit 5
```

#### c) Issues (`issues`)
- Shows all issue activity
- Filters by state (open, closed)
- Displays issues organized by repository
- Optional: Limit the number of results

**Example:**
```bash
gh-activity user username issues --state open
gh-activity user username issues --state closed --limit 10
```

### 2. **Set GitHub Token** (`set token`)
Save your GitHub personal access token securely in your home directory.

**Example:**
```bash
gh-activity set token your_github_token_here
```

**What happens:**
- Token is saved in `~/.gh-activity/config.json`
- File permissions are set to secure (0600)
- Timestamp is recorded when token was set
- You'll get a warning if token is older than 90 days

### 3. **Get GitHub Token** (`get token`)
View your currently saved GitHub token.

**Example:**
```bash
gh-activity get token
```

### 4. **Version Info**
Check the version of the tool.

**Example:**
```bash
gh-activity version
gh-activity --version
gh-activity -v
```

## Available Commands

| Scope | Command | Required Flags | Optional Flags | Description |
|-------|---------|----------------|----------------|-------------|
| `user` | `pushes` | - | `--limit` | Get push events |
| `user` | `pulls` | `--state` | `--limit` | Get pull requests |
| `user` | `issues` | `--state` | `--limit` | Get issues |
| `set` | `token` | - | - | Set GitHub token |
| `get` | `token` | - | - | Get saved token |

## Flags Explained

### `--limit`
- Controls how many results to show
- **Default:** 2
- **Maximum for issues:** 50
- **Example:** `--limit 10`

### `--state`
- Filters results by state
- **For pulls:** `open`, `closed`, `merged`
- **For issues:** `open`, `closed`
- **Required** for pulls and issues commands

## How to Use

### Basic Structure
```bash
gh-activity <scope> <command> [flags]
```

### Examples

1. **See recent pushes of a user:**
```bash
gh-activity user octocat pushes
```

2. **See open pull requests with limit:**
```bash
gh-activity user octocat pulls --state open --limit 5
```

3. **See closed issues:**
```bash
gh-activity user octocat issues --state closed --limit 10
```

4. **Save your GitHub token:**
```bash
gh-activity set token ghp_yourTokenHere
```

5. **Check saved token:**
```bash
gh-activity get token
```

## Project Structure

```
github-activity/
├── cmd/
│   └── main.go                 # Main entry point
├── internal/
│   ├── config/
│   │   ├── config-path.go     # Config file path management
│   │   └── gh-token.go        # Token storage and retrieval
│   ├── custom-error/
│   │   └── global_error.go    # Custom error handling
│   ├── github/
│   │   ├── events.go          # GitHub API interaction
│   │   └── valid_scope.go     # Scope and command validation
│   ├── handlers/
│   │   ├── issue_event.go     # Issue event handler
│   │   ├── pull_event.go      # Pull request handler
│   │   └── push_event.go      # Push event handler
│   ├── models/
│   │   ├── event_model.go     # Event data models
│   │   └── rules.go           # Command rules and scopes
│   └── services/
│       ├── issue_event.go     # Issue processing service
│       ├── pull_event.go      # Pull request processing service
│       └── push_event.go      # Push processing service
├── go.mod
└── README.md
```

## Error Handling

The tool provides clear error messages for:
- Missing arguments
- Invalid scope or commands
- Invalid flags
- Missing required flags
- Token issues
- API errors

## Security Features

- GitHub token is stored securely with restricted file permissions (0600)
- Token age warning (90 days)
- Timestamp tracking for token creation

## Current Version

**v1.0.0**

## Technical Details

- **Language:** Go 1.25.3
- **API:** GitHub REST API v3
- **Default API Limit:** 60 events per user
- **Config Location:** `~/.gh-activity/config.json`

## Notes

- Default limit for all commands is 2 results
- Maximum limit for issues is 50
- Token storage location: `~/.gh-activity/config.json`
- The tool fetches up to 60 recent events from GitHub API

## Future Scope

- Repository scope (`repo`) with info command
- More filters and sorting options
- Support for organizations
- Enhanced formatting and colors

---

**Developer:** github.com/ndk123-web  
**Project:** github-activity
