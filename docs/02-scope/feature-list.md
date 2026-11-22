# Feature Scope and Attack Surface

Pick 7-8 features. Mark selected ones. Keep each small.

| # | Feature | In scope? | Actors | Key attack classes | Phase to build |
|---|---|---|---|---|---|
| 1 | Registration, login, password reset, MFA | [ ] | Visitor, customer | Credential stuffing, enumeration, session attacks, ATO | |
| 2 | Menu and search | [ ] | All | SQLi, XSS | |
| 3 | Cart and ordering | [ ] | Customer | IDOR, price tampering, logic abuse | |
| 4 | Payment (mock provider + webhook) | [ ] | Customer, provider | Webhook forgery, replay, SSRF | |
| 5 | Loyalty points / gift cards | [ ] | Customer | Race conditions, double spend, fraud | |
| 6 | Barista/staff portal | [ ] | Barista | RBAC, privilege escalation | |
| 7 | Admin panel (menu, pricing, users) | [ ] | Admin | CSRF, authz flaws, audit gaps | |
| 8 | Image upload | [ ] | Customer, admin | Malicious files, path traversal, sniffing | |
| 9 | Reviews / comments | [ ] | Customer | Stored XSS, spam | |
| 10 | Receipt / CSV export | [ ] | Customer, manager | CSV injection, data leakage | |
| 11 | Public REST API | [ ] | Third party | Mass assignment, rate limit, key handling | |
| 12 | Supplier / inventory integration | [ ] | Supplier system | Third-party trust, supply chain | |

## Selected set rationale
_Why these features give diverse coverage._

## Explicitly out of scope
-

## Entry points inventory (fill after selection)
| Entry point | Method | Auth required | Input types | Notes |
|---|---|---|---|---|
