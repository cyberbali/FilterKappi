# Verification

## Plan
| Method | Scope | When | Output |
|---|---|---|---|
| Unit and abuse-case tests | Each requirement | Every PR | CI results |
| SAST | gosec, staticcheck, CodeQL | Every PR | CI results |
| SCA | govulncheck, Dependabot | Every PR, weekly | Alerts |
| DAST | OWASP ZAP baseline against local deployment | Per milestone | `zap-report-<date>.md` |
| Manual pen test | Each feature slice, using the threat model | Per milestone | `pentest-<date>.md` |
| Config review | Docker, compose, headers, TLS | Per milestone | Checklist |

## Baseline vs hardened
Record findings against `v0.1-baseline` first, then re-test the hardened version.

## Findings log
| ID | Date | Source | Severity | Description | Linked req / risk | Status |
|---|---|---|---|---|---|---|
