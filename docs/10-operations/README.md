# Detection and Response

## Contents to produce
- `logging-standard.md`: events to log, fields, redaction rules, retention
- `alerting.md`: detections (failed login spikes, webhook signature failures, privilege changes, unusual refunds)
- `incident-response-plan.md`: roles, severity levels, communication, evidence handling
- `runbooks/`: one per scenario (account takeover wave, leaked secret, payment fraud, data exposure)
- `backup-restore.md`: RTO/RPO, restore test results

## Incident severity (draft)
| Sev | Definition | Response target |
|---|---|---|
| 1 | Confirmed data breach or active fraud | Immediate |
| 2 | Suspected compromise, major control failure | 4 hours |
| 3 | Contained anomaly | 1 business day |
