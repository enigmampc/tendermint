# Unreleased Changes

## v0.34.25

### BREAKING CHANGES

- CLI/RPC/Config

- Apps

- P2P Protocol

- Go API

- Blockchain Protocol

### FEATURES

### IMPROVEMENTS
- [cli] \#9171 add `--hard` flag to rollback command (and a boolean to the `RollbackState` method). This will rollback
  state and remove the last block. This command can be triggered multiple times. The application must also rollback
  state to the same height. (@tsutsu, @cmwaters)

- [consensus] \#9760 Save peer LastCommit correctly to achieve 50% reduction in gossiped precommits. (@williambanfield)

### BUG FIXES

