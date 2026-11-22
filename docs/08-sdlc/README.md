# Secure SDLC

## Branching and review
- `main` protected: PR required, status checks required, no force push, signed commits
- Architect review required for changes under `docs/07-architecture/` and `.github/`

## CI gates (must pass to merge)
| Gate | Tool |
|---|---|
| Format and vet | `gofmt`, `go vet` |
| Static analysis | `staticcheck`, `gosec`, CodeQL |
| Dependency vulnerabilities | `govulncheck`, Dependabot |
| Secrets | `gitleaks`, GitHub push protection |
| Tests | `go test -race` |
| Container | image scan (Trivy), non-root check |

## Dependency policy
- Standard library first
- New dependency needs: purpose, maintainer health, license, alternatives, pinned version
- Review `go.sum` changes in PRs

## Definition of done
See `AGENTS.md`.

## Release and tagging
Tag phase completions (`phase-4-threat-model`) and builds (`v0.1-baseline`, `v1.0-hardened`).
