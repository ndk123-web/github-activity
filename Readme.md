# gh-activity

A lightweight, cross-platform CLI tool to inspect **recent GitHub user activity** using GitHub’s public Events API.

`gh-activity` is designed for developers who want **quick, readable insights** into pushes, pull requests, and issues — directly from the terminal.

---

## Table of Contents

- [Get Started](#get-started)
- [Installation](#installation)

  - [Windows](#windows)
  - [Linux](#linux)
  - [macOS](#macos)

- [Authentication](#authentication)
- [Usage](#usage)

  - [Scopes & Commands](#scopes--commands)
  - [Pushes](#pushes)
  - [Pull Requests](#pull-requests)
  - [Issues](#issues)
  - [Watches](#watches)

- [Flags](#flags)
- [Notes](#notes)
- [Roadmap](#roadmap)
- [License](#license)

---

## Get Started

`gh-activity` is shipped as a **single self-contained binary**.

You download the release for your operating system, add it to your `PATH`, and start using it immediately — **no Go installation required**.

---

## Installation

Releases are distributed as **ZIP archives** that already contain the correctly named binary inside a `bin/` directory.

### General Structure (inside the archive)

```
gh-activity/
└── bin/
    └── gh-activity        (or gh-activity.exe on Windows)
```

You only need to add the `bin/` directory to your system `PATH`.

---

### Windows

1. Download the latest release:

```
gh-activity-windows-amd64.zip
```

2. Extract the archive
3. Add the extracted `bin` folder to **Environment Variables → PATH**
4. Open a new terminal and verify:

```bash
gh-activity --help
```

---

### Linux

1. Download:

```
gh-activity-linux-amd64.zip
```

2. Extract:

```bash
unzip gh-activity-linux-amd64.zip
```

3. Add `bin/` to PATH:

```bash
export PATH=$PATH:/path/to/gh-activity/bin
```

4. (Optional) If the binary is not executable:

```bash
chmod +x gh-activity
```

5. Verify:

```bash
gh-activity --help
```

---

### macOS

1. Download the correct archive:

- Intel: `gh-activity-darwin-amd64.zip`
- Apple Silicon: `gh-activity-darwin-arm64.zip`

2. Extract:

```bash
unzip gh-activity-darwin-amd64.zip
```

3. Add `bin/` to PATH

4. (Optional) If the binary is not executable:

```bash
chmod +x gh-activity
```

5. Verify:

```bash
gh-activity --help
```

---

## Authentication

### 🔐 Creating a GitHub Personal Access Token

To avoid API rate limits, you can create a **GitHub Personal Access Token**.

1. Go to **GitHub → Settings**
2. Open **Developer settings** (left sidebar)
3. Click **Personal access tokens**
4. Select **Tokens (classic)**
5. Click **Generate new token**
6. Give it a name (e.g. `gh-activity`)
7. Select scopes:

   * ✅ `public_repo` *(recommended)*
8. Generate the token and **copy it immediately**

Then set it in the CLI:
```bash
gh-activity set token <your-token>
```

### What happens internally

- Token is stored locally at:

```
~/.gh-activity/config.json
```

- File permissions are restricted (`0600`)
- A timestamp is recorded
- You’ll receive a warning if the token is older than 90 days

To view the saved token:

```bash
gh-activity get token
```

---

## Usage

### Scopes & Commands

The CLI follows a **strict positional structure**:

```bash
gh-activity <scope> <entity> <command> [flags]
```

Currently supported scope:

- `user`

---

### Pushes

View recent push activity by a user, grouped by repository.

```bash
gh-activity user <username> pushes [--limit N]
```

**Flags:**

- `--limit` _(optional)_ — Number of recent push events (default: 2)

---

### Pull Requests

View pull request activity filtered by state.

```bash
gh-activity user <username> pulls --state <open|closed|merged> [--limit N]
```

**Flags:**

- `--state` _(required)_ — `open`, `closed`, or `merged`
- `--limit` _(optional)_ — Number of results (default: 2)

---

### Issues

View issue activity (pull requests are automatically excluded).

```bash
gh-activity user <username> issues --state <open|closed> [--limit N]
```

**Flags:**

- `--state` _(required)_ — `open` or `closed`
- `--limit` _(optional)_ — Number of results (default: 2)

---

### Watches

View watch/star events grouped by repository.

```bash
gh-activity user <username> watches [--limit N]
```

**What it shows:**
- Count of recent `WatchEvent` per repository for the user’s recent activity window
- An overview summary and a neat table like other commands

**Flags:**

- `--limit` _(optional)_ — Number of distinct repositories to include (default: 2, max: 50)

**Example:**

```bash
gh-activity user octocat watches --limit 5
```

---

## Flags

### `--limit`

- Controls how many events are displayed
- Default: `2`

### `--state`

- Required for `pulls` and `issues`
- Pull requests: `open`, `closed`, `merged`
- Issues: `open`, `closed`

---

## Notes

- Data is fetched from GitHub’s **Events API**
- Events are **recent activity only** (not full history)
- One push event ≠ one commit
- Output is optimized for terminal readability and scripting

---

## License

MIT

---
