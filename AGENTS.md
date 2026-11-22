# Instructions for the coding agent

You are the dev team. A separate person plays the security architect.
Implement what is asked, nothing more. Do not make security design decisions.
If something is unspecified or ambiguous, stop and ask.

Follow `docs/requirements.md`. Do not edit anything under `docs/` unless asked (except `docs/journal/` for that date, which must be created if not present for that day and updated every time a significant change takes place or a learning is done, and `docs/version.md`, which must be bumped every time a change takes place based on Semantic Versioning: MAJOR, MINOR, or PATCH).

## Version control
Do not run git commands, create commits, branches, or PRs. The architect handles
all of that. Leave changes in the working tree.

## Rules
- Go standard library first. Adding a dependency needs a written justification
  and a pinned version. Ask before adding one.
- Config comes from environment variables, validated at startup. Fail closed on
  bad config. No hardcoded secrets, and no real secrets in `.env.example`.
- Never log secrets or personal data.
- Parameterized queries only. Validate input at the boundary. Encode output by context.
- Generic errors to clients, details to structured logs (`log/slog`).
- Set timeouts on servers, clients, and DB connections. Limit request body size.
- Authorization is checked server-side on every request, deny by default.
- Write tests, including negative and abuse-case tests.
- Do not add features, endpoints, or files that are not in the current task.

## When you finish a task
- Must update `docs/version.md` based on Semantic Versioning (MAJOR for breaking changes, MINOR for new features/controls, PATCH for fixes/docs/refactoring).
- Must update the version in the title of `README.md` (e.g. `# FilterKappi (v 1.0.0) ☕`) whenever the version is updated, and update other sections if needed.
Reply with a short summary:
1. What you built or changed (files)
2. Which requirement IDs it covers (e.g. AUTH-001)
3. Which tests prove it
4. Anything you were unsure about or deliberately did not do