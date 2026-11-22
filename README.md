# FilterKappi ☕

> Great filter coffee keeps out the grounds; FilterKappi keeps out everything untrusted.

A testament to secure application design - hypothetical, security-hardened coffee shop web application, built in Go by **Cyberbali**.

The goal is not the coffee shop. The goal is to practise the full security
engineering lifecycle: discovery, threat modeling, risk assessment, requirements,
architecture, secure build, verification, and operations, all documented here.

## How this repo is organised

| Path | Purpose |
|---|---|
| `docs/00-engagement` | Charter, personas, rules of the exercise |
| `docs/01-discovery` | Architect questionnaire and client answers |
| `docs/02-scope` | Features and their attack surface |
| `docs/03-assets` | Asset inventory and data classification |
| `docs/04-threat-model` | DFDs, trust boundaries, STRIDE, abuse cases, attack trees |
| `docs/05-risk-register` | Risks, ratings, treatments |
| `docs/06-requirements` | Security requirements (ASVS-based) and traceability matrix |
| `docs/07-architecture/adr` | Architecture Decision Records |
| `docs/08-sdlc` | Secure SDLC, CI gates, dependency policy |
| `docs/09-verification` | Test plans, scan results, pen-test reports |
| `docs/10-operations` | Logging, alerting, runbooks, incident response |
| `docs/journal` | Daily log (`YYYY-MM-DD.md`) |
| `cmd/`, `internal/`, `migrations/`, `test/` | Application code |

## Status

See the latest entry in `docs/journal/` and the open milestones.

## Phases

0 Setup, 1 Discovery, 2 Scope, 3 Assets, 4 Threat model, 5 Risk, 6 Requirements,
7 Architecture, 8 SDLC, 9 Build, 10 Verification, 11 Detection and response,
12 Residual risk and retrospective.

> This is a training project. Nothing here has been security reviewed by anyone but its author.
