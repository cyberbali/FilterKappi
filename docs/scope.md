# Scope

## Business context
FilterKappi is a to-go coffee shop (cuppa, americano, espresso, and similar) with 2 locations.
Customers order and pay online, then collect at the shop using a pickup code.
The system is intended to grow over time to a wide feature set (see Features by wave).

- Volume: about 400 orders per day across both locations, about 60% in the morning rush
- Peak load is low (roughly 1 to 2 orders per minute), so abuse will stand out in logs
- Downtime cost: about 20 customers lost per hour at 9 a.m. on a Monday
- Owner constraint: security wins when it conflicts with convenience, but the customer
  interface should stay sleek

## Users and roles
| Role | Can | Cannot |
|---|---|---|
| Customer | Register (name and phone), log in, browse menu, order, pay online, view own order history, cancel own order before preparation starts | See anyone else's data |
| Barista (scoped to one location) | See that location's orders (customer name and items only), verify pickup code, mark orders ready or collected, mark a product out of stock, cancel an order (reason required, logged) | Edit orders, see customer phone numbers, change prices, see sales totals, see the other location |
| Admin | Everything: edit orders, manage menu, prices, offers, and users, view the log of critical business decisions, access both locations | N/A |

Notes:
- No guest checkout. A name and phone number are mandatory.
- Barista accounts are individual, even on the shared counter tablet.
- Admin accounts are individual and named. A minimum of two admins. No shared credentials.

## Features by wave
| Wave | Features | Main attack classes |
|---|---|---|
| 1 | Registration, login, MFA, account recovery; menu and search; cart and ordering with pickup code | Credential and OTP abuse, enumeration, session attacks, account takeover, SQL injection, XSS, IDOR, price tampering, logic abuse |
| 2 | Payment (mock provider and webhook); admin-created offers (replaces loyalty and gift cards for now) | Webhook forgery, replay, SSRF, offer reuse and stacking, race conditions, fraud |
| 3 | Barista portal; admin panel (menu, prices, offers, users, order edit, critical-decision log) | RBAC, privilege escalation, location boundary bypass, CSRF, audit gaps |
| 4 | Image upload; reviews; receipt and CSV export | Malicious files, path traversal, stored XSS, spam, CSV injection, data leakage |
| 5 | Public REST API; supplier and inventory integration | Mass assignment, rate limiting, API key handling, third-party trust, supply chain |

## Out of scope for now
- QR code at the shop entrance (deferred: a swappable physical entry point adds complexity)
- Gift cards and points-based loyalty
- Payment methods other than online prepayment
- Guest checkout
- Real payment processing and real messaging providers (both mocked locally)

## Data held
| Data | Classification | Notes |
|---|---|---|
| Customer name and phone | Confidential | Personal data under India's DPDP Act |
| Order history | Confidential | Retention rule below |
| Pickup codes | Restricted | Short-lived |
| Password and session material, MFA secrets | Restricted | |
| Staff and admin accounts | Restricted | |
| Offers and prices | Internal | Integrity matters |
| Audit log of critical decisions | Confidential | Integrity matters |

No card data is stored. Payment is handled by a provider and only references are kept.

## Key decisions
1. **QR code removed.** Customers reach the app by URL only.
2. **Cancellation is allowed but audited.** Customer: only before preparation starts.
   Barista: reason required. Admin: any time. Every cancellation is written to a log
   shown in a "Critical decisions" section of the admin panel. Paid cancellations refund
   to the original payment method only.
3. **Order edits are admin-only.** This prevents theft after payment.
4. **Three roles only:** customer, barista, admin.
5. **Location boundary.** A barista sees only their own location's orders and stock.
6. **Customer login.** Phone verification by one-time code (mocked provider behind an
   interface), then a rotating long-lived session. TOTP is optional for customers and
   mandatory for baristas and admins.
7. **Offboarding.** Admin disables accounts immediately. Any account inactive for 7 days is
   auto-disabled. Accounts are disabled, not deleted, so audit records stay attributable.
8. **Admin continuity.** At least two named admin accounts. A break-glass procedure is
   documented later.
9. **Network access.** Staff and admin portals are reachable only from the shop network,
   with admin remote access through a VPN. Customers use a separate guest endpoint.
10. **Pickup code** confirms the right person collects the right order.

## Compliance and retention
- India's DPDP Act applies to the name and phone number.
- Order history is kept for 1 year (owner rule, for tax). [A] Verify with an accountant:
  tax record-keeping periods in India are commonly longer.
- On a deletion request, personal fields are anonymized and transaction records are kept
  for the retention period. [A] Confirm.

## Attackers and crown jewels
- Most feared: fraudsters. The worst day is a fraud incident.
- Also: ex-baristas (hence disable on offboarding and after 7 days inactive), bots.
- Crown-jewel goals: refund and cancellation abuse, offer abuse, customer or admin account
  takeover, payment bypass (for example a forged webhook), customer data theft.

## Operations and recovery
- One operator (the owner) is on call. No separation of duties. This is an accepted risk.
- Budget: whatever the security requirements justify (a simplification for this exercise).
- Staff use a shared counter tablet on a secure Wi-Fi network. Customers use guest access.
- [A] Recovery time target: 4 hours.
- [A] Recovery point target: no more than 15 minutes of paid-order data lost.

## Assumptions to validate
| ID | Assumption | Status | Impact if wrong |
|---|---|---|---|
| A1 | Average order value is about 250 INR | Owner accepted | Changes fraud impact ratings |
| A2 | 1-year order retention satisfies tax rules | Open: verify with accountant | Compliance gap |
| A3 | Deletion means anonymizing personal fields | Open: confirm legally | Legal and design impact |
| A4 | RTO 4 hours, RPO 15 minutes | Owner accepted | Backup and recovery design |
| A5 | Customers accept optional TOTP plus long sessions | Owner accepted | Account takeover risk |
| A6 | Customer can cancel only before preparation starts | Owner accepted | Refund fraud exposure |

## Open items
- Refund policy details (who triggers, limits)
- Offer rules (stacking, single use, expiry)
- Cancellation alert threshold per barista per shift
