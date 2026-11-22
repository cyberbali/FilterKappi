# Asset Inventory and Data Classification

## Classification scheme
| Level | Definition | Examples | Handling |
|---|---|---|---|
| Public | Safe to disclose | Menu, prices | None |
| Internal | Business use, low harm if leaked | Order counts, inventory | Auth required |
| Confidential | Harm to individuals or business | Customer PII, order history | Encrypt in transit and at rest, access logged |
| Restricted | Severe harm if disclosed or altered | Credentials, secrets, keys, payment tokens | Strongest controls, minimal access |

## Data assets
| ID | Asset | Classification | Location | Owner | Retention | Regulation |
|---|---|---|---|---|---|---|
| A-01 | Customer account (email, name, phone) | Confidential | DB | | | |
| A-02 | Password hashes | Restricted | DB | | | |
| A-03 | Session tokens | Restricted | Cookie, DB | | | |
| A-04 | Order history | Confidential | DB | | | |
| A-05 | Loyalty/gift card balances | Confidential (integrity-critical) | DB | | | |
| A-06 | Payment tokens / provider references | Restricted | DB | | | |
| A-07 | Staff accounts and roles | Restricted | DB | | | |
| A-08 | Audit logs | Confidential | Log store | | | |
| A-09 | Uploaded images | Internal | File store | | | |
| A-10 | Application secrets and keys | Restricted | Env / vault | | | |

## System assets
| ID | Asset | Notes |
|---|---|---|
| S-01 | Web application | |
| S-02 | Database | |
| S-03 | Reverse proxy / TLS termination | |
| S-04 | CI/CD pipeline and repo | |
| S-05 | Staff devices | |
| S-06 | Backups | |

## Security objectives per asset (CIA + non-repudiation)
| Asset | Confidentiality | Integrity | Availability | Notes |
|---|---|---|---|---|
