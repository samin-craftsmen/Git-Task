
## Dependencies

- Merge PR #2

## What does this PR do?

- Consolidates 6 commits into 3 commits in branch 1.

## Type of Change

- [x] Refactor

## What was changed

- 6 commits were squashed into 3

## Changelog

Refactor: Consolidate 6 commits into 3


## How to Test

1. Manual Testing

## How QA Should Test

- Look at the commits

## Rollback Plan

| Scenario | Command(s) | Notes |
|----------|------------|-------|
| Undo rebase locally (keep changes) | `git reflog`<br>`git reset --hard <commit-before-rebase>` | Use `git reflog` to find the commit before rebase. All local changes after that commit will be lost. |
| Undo rebase on remote (force push) | `git push origin <branch> --force` | Overwrites GitHub history. Only do if branch is not shared or team agrees. |
| Undo without rewriting history (safe) | `git revert <commit-hash>` | Creates new commits that undo the rebased commits. Safe for shared branches. |

## Checklist

- [x] Requirements from Task - 1,2,3 are met

## Screenshots (if applicable)

- N/A

## Note for Reviewer

- N/A