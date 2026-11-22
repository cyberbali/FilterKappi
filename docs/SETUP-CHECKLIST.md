# GitHub setup checklist

- [ ] Create repo (private at first), push this skeleton
- [ ] Replace `@your-github-username` in `.github/CODEOWNERS`
- [ ] Settings > Branches: protect `main` (require PR, require status checks, block force pushes, require linear history, require signed commits)
- [ ] Settings > Code security: enable Dependabot alerts and updates, secret scanning, push protection, CodeQL default setup
- [ ] Set up commit signing (SSH or GPG) and add key to GitHub
- [ ] Create milestones: one per phase (0-12)
- [ ] Create labels: `threat-model`, `requirement`, `adr`, `finding`, `task`, `risk`, `docs`
- [ ] Commit `AGENTS.md` before the first agent run
- [ ] Optional: tag `phase-0-setup`
