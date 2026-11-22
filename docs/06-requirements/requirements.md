# Security Requirements

Standard: OWASP ASVS (target **Level 2**, adjust with justification).
ID format: `<DOMAIN>-<NNN>`. Domains: AUTH, SESS, ACCESS, INPUT, DATA, CRYPTO, API, FILE, LOG, CONF, PAY, LOGIC, INFRA, SDLC.

| ID | Requirement (shall...) | ASVS ref | Source threat/risk | Priority | Verification | Status |
|---|---|---|---|---|---|---|
| AUTH-001 | Store passwords using Argon2id or bcrypt with a per-user salt | V2.4 | T-001 / R-001 | Must | Unit test, code review | Draft |
| AUTH-002 | Rate limit and throttle repeated failed logins per account and per IP | V2.2 | T-001 / R-001 | Must | Test | Draft |
| ACCESS-001 | Enforce object-level authorization on every order lookup | V4.2 | T-003 | Must | Abuse-case test | Draft |
| PAY-001 | Verify webhook signature and reject replays | V13 | T-004 / R-002 | Must | Test | Draft |

## ASVS chapter applicability
| ASVS chapter | In scope? | Justification |
|---|---|---|
| V1 Architecture | | |
| V2 Authentication | | |
| V3 Session management | | |
| V4 Access control | | |
| V5 Validation, sanitization, encoding | | |
| V6 Stored cryptography | | |
| V7 Error handling and logging | | |
| V8 Data protection | | |
| V9 Communication | | |
| V10 Malicious code | | |
| V11 Business logic | | |
| V12 Files and resources | | |
| V13 API and web service | | |
| V14 Configuration | | |
