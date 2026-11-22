# FilterKappi ☕

> Great filter coffee keeps out the grounds; FilterKappi keeps out everything untrusted.

A testament to secure application design - hypothetical, security-hardened coffee shop web application, built in Go by **Cyberbali**.

It exists to practise security architecture: analysing a system, modelling
threats, specifying controls, building them, and attacking the result.

## Approach

1. Interview the "client" and define the scope
2. Model threats
3. Write requirements for the threats that matter
4. Build a small slice, then attack it
5. Fix what broke and record the findings

## Docs

| File | Purpose |
|---|---|
| [docs/scope.md](docs/scope.md) | Features, users, data |
| [docs/threat-model.md](docs/threat-model.md) | Diagram, boundaries, threats |
| [docs/requirements.md](docs/requirements.md) | Security requirements and design decisions |
| [docs/findings.md](docs/findings.md) | Results of attacking my own code |
| [docs/journal/](docs/journal/) | Running notes |

## Stack

Go (standard library first), local deployment only.

## Status

Scoping. No application code yet.