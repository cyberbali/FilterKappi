# Instructions for the coding agent (the "dev team")

You are the development team for FilterKappi, a Go web application. A separate
person plays the **security architect**. The architect makes security design
decisions; you implement them.

## Authority

- Do not make security design decisions on your own. Follow:
  - ADRs in `docs/07-architecture/adr/`
  - Requirements in `docs/06-requirements/`
- If something is unspecified or ambiguous, **stop and ask**, or propose a draft ADR. Do not guess.
- Do not add features, endpoints, or dependencies that are not in the current task.

## Workflow

- One small, focused PR per task. Branch from `main`.
- Every PR must reference requirement IDs (e.g. `AUTH-003`) and name the tests that prove them.
- Use Conventional Commits: `feat:`, `fix:`, `sec:`, `test:`, `docs:`, `chore:`.
- Update `docs/06-requirements/traceability-matrix.md` when a requirement gets implemented or tested.
- Add a note to the day's `docs/journal/` file if you made any non-obvious decision.

## Engineering rules

- Go standard library first. Any new dependency needs a written justification in the PR and must be pinned.
- Configuration comes from environment variables, validated at startup. Fail closed on bad config.
- No hardcoded secrets. Never log secrets, tokens, passwords, or PII.
- Parameterized queries only. No string-built SQL.
- Validate input at the boundary (type, length, range, allowlist). Encode output by context.
- Errors returned to clients are generic. Details go to structured logs (`log/slog`) with a request ID.
- Set timeouts on every server, client, and DB connection. Limit request body sizes.
- Authorization is checked server-side on every request, deny by default.
- Prefer constant-time comparison for secrets and tokens.
- Write tests alongside code, including **negative and abuse-case tests**.

## Definition of done

- `make test lint vuln` passes
- CI green (vet, staticcheck, govulncheck, gosec, gitleaks)
- Requirement IDs and tests listed in the PR
- Docs updated

## Never

- Disable a security control to make a test pass.
- Commit `.env` files, keys, or real credentials.
- Merge your own work without architect review.
