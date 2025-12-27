# Authentication

Use a GitHub Personal Access Token (PAT) to avoid low unauthenticated rate limits.

## Set and Get Token

```powershell
# Save your token
gh-activity set token <your_token>

# Show current token
gh-activity get token
```

## Storage

- Token is saved to a JSON config file resolved by `internal/config/config-path.go`.
- File includes:
  - `token`: string
  - `created_at`: timestamp
- Permissions are restricted and a 90-day age warning is displayed.

## Invalid Token Handling

- If GitHub returns `Bad credentials`:
  - CLI prints a helpful message
  - Deletes the saved token
  - Retries the request without authentication

## Creating a PAT

1. GitHub → Settings → Developer settings → Personal access tokens → Tokens (classic)
2. Generate new token
3. Recommended scope: `public_repo`
4. Copy the token and set it via `gh-activity set token <your_token>`

## Security Notes

- Do not commit your token.
- Rotate tokens periodically; the CLI warns if older than 90 days.