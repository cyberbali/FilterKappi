# Discovery Questionnaire

Instructions: answer in character as the client. If you do not know, write a reasonable
**ASSUMPTION** and flag it for validation. Add follow-up questions as they occur to you.

Status key: `[ ]` open, `[x]` answered, `[A]` assumption

## A. Business and context
| ID | Question | Answer | Status |
|---|---|---|---|
| B1 | Single shop, chain, or franchise? Plans to expand? | | [ ] |
| B2 | Online ordering only, in-store only, or both? | | [ ] |
| B3 | Estimated customers, orders per day, peak hour load? | | [ ] |
| B4 | Cost of one hour of downtime? Of a full day? | | [ ] |
| B5 | Worst-case scenario the owner fears most? | | [ ] |
| B6 | Budget and timeline? Who maintains the system long term? | | [ ] |
| B7 | Risk appetite, and who signs off on accepted risk? | | [ ] |
| B8 | Brand and reputation sensitivity? | | [ ] |

## B. Users and roles
| ID | Question | Answer | Status |
|---|---|---|---|
| U1 | Which roles exist (customer, barista, manager, admin, supplier, owner)? | | [ ] |
| U2 | How many of each? Staff turnover rate? | | [ ] |
| U3 | What may each role do? Which data may each see? | | [ ] |
| U4 | Do baristas share devices or logins at the counter? | | [ ] |
| U5 | Are guest (no account) orders allowed? | | [ ] |
| U6 | Any minors as users? | | [ ] |
| U7 | Is MFA acceptable for customers? Mandatory for staff/admin? | | [ ] |
| U8 | How is account recovery done, and who can help a locked-out user? | | [ ] |

## C. Data
| ID | Question | Answer | Status |
|---|---|---|---|
| D1 | What personal data is collected (name, email, phone, address, location)? | | [ ] |
| D2 | Is card data stored, processed, or only tokenized by a provider? | | [ ] |
| D3 | What about order history, loyalty balances, gift card balances? | | [ ] |
| D4 | Any staff data (schedules, payroll, ID documents)? | | [ ] |
| D5 | Retention period per data type? Deletion on request? | | [ ] |
| D6 | Is any data shared with third parties (analytics, marketing, suppliers)? | | [ ] |
| D7 | User-generated content: reviews, photos, notes on orders? | | [ ] |

## D. Compliance and legal
| ID | Question | Answer | Status |
|---|---|---|---|
| L1 | Which regimes apply: PCI DSS scope, India DPDP Act, GDPR (any EU customers), tax/invoice rules? | | [ ] |
| L2 | Breach notification obligations and timelines? | | [ ] |
| L3 | Contractual duties to payment, delivery, or loyalty partners? | | [ ] |
| L4 | Privacy notice, consent, and cookie requirements? | | [ ] |
| L5 | Audit or evidence requirements (who must be able to prove what)? | | [ ] |

## E. Threat landscape
| ID | Question | Answer | Status |
|---|---|---|---|
| T1 | Who might attack: fraudsters, bots, competitors, disgruntled staff, ransomware crews, curious customers? | | [ ] |
| T2 | What do they want: free coffee, loyalty points, customer data, disruption, payments diverted? | | [ ] |
| T3 | Past incidents here or at comparable shops? | | [ ] |
| T4 | Which abuse is most likely on day one (promo abuse, fake accounts, review spam)? | | [ ] |
| T5 | What would be the most embarrassing public failure? | | [ ] |

## F. Functionality (drives attack surface)
| ID | Question | Answer | Status |
|---|---|---|---|
| F1 | Which features are launch-critical vs later? | | [ ] |
| F2 | Is there a loyalty program? Gift cards? Referral or promo codes? | | [ ] |
| F3 | Can customers upload images? Can staff? | | [ ] |
| F4 | Are there exports (receipts, CSV reports)? Who can run them? | | [ ] |
| F5 | Is there a public API or mobile app planned? | | [ ] |
| F6 | Which third-party integrations are needed (payments, email/SMS, suppliers)? | | [ ] |

## G. Technical and operational
| ID | Question | Answer | Status |
|---|---|---|---|
| O1 | Where does it run now (local) and later (cloud, on-prem, VPS)? | | [ ] |
| O2 | Database choice and who administers it? | | [ ] |
| O3 | How are secrets and keys stored and rotated? | | [ ] |
| O4 | Who has production access? Is there separation of duties? | | [ ] |
| O5 | Deployment process, environments, rollback? | | [ ] |
| O6 | What logging and monitoring exists? Who responds at 2 a.m.? | | [ ] |
| O7 | Backup frequency, RTO, RPO? Restore tested? | | [ ] |
| O8 | Staff devices: managed or personal? Shop Wi-Fi shared with customers? | | [ ] |
| O9 | Dependency and patching approach? | | [ ] |

## H. Constraints and non-functional
| ID | Question | Answer | Status |
|---|---|---|---|
| N1 | Performance targets (latency, concurrency)? | | [ ] |
| N2 | Accessibility and language requirements? | | [ ] |
| N3 | Checkout friction tolerance (the "don't annoy customers" limit)? | | [ ] |
| N4 | Offline operation at the counter if internet drops? | | [ ] |

## Follow-ups raised by answers
| ID | Follow-up question | Raised by | Answer |
|---|---|---|---|
| | | | |

## Assumptions to validate
| ID | Assumption | Impact if wrong | Validation plan |
|---|---|---|---|
| | | | |
