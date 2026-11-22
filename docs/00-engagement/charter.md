# Engagement Charter

## Purpose
Train security architecture skills by running a realistic engagement for a fictional
client, FilterKappi, from discovery through verification and operations.

## Roles
| Role | Played by | Responsibility |
|---|---|---|
| Client (owner of FilterKappi) | Me, in character | Business goals, budget, risk acceptance sign-off |
| Security architect | Me, in character | Questions, analysis, requirements, architecture, reviews |
| Dev team | Coding agent | Implements what the architect specifies, nothing more |

## Client persona (edit freely)
- Name: _______
- Business: single-location specialty coffee shop, planning online ordering and a loyalty program
- Budget: small. Prefers open source and low operational overhead
- Attitude to security: "Important, but don't make checkout annoying"
- Risk appetite: moderate. Will not accept card-data storage or customer-data breach risk
- Sign-off authority for accepted risk: the client persona

## Goals
1. Produce a traceable chain: **threat -> risk -> requirement -> control -> test**
2. Build a deliberately diverse (but small) feature set with a wide attack surface
3. Document every phase in this repo
4. Attack my own system and record findings honestly

## Rules of the exercise
- Client answers must be written down, with unknowns flagged as **ASSUMPTION**
- Architect decisions are recorded as ADRs with rejected alternatives
- Dev team does not make security decisions; it asks
- No real customer data, no real payment credentials, ever
- Deployment is local only for now. Any deployment change needs a new ADR and threat model update
- Accepted risks are signed off explicitly in the risk register

## Out of scope
Real payment processing, real customers, production hosting, physical security of the shop.

## Success criteria
- Every in-scope requirement has an owner, a control, and a test
- Residual risks are listed and explicitly accepted
- A before/after story exists: baseline findings vs hardened findings

## Daily routine
1. Read yesterday's journal "Next" section
2. Work one phase task or one dev slice
3. Commit via PR, update traceability matrix
4. Write `docs/journal/YYYY-MM-DD.md`
