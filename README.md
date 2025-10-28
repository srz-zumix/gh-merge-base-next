# gh-merge-base-next

[![Build and Test](https://github.com/srz-zumix/gh-merge-base-next/actions/workflows/build.yml/badge.svg)](https://github.com/srz-zumix/gh-merge-base-next/actions/workflows/build.yml)

gh-merge-base-next is a tool to find the next commit on first-parent path from merge-base.

## Installation

```bash
gh extension install srz-zumix/gh-merge-base-next
```

## Features

### Core Functionality

The tool identifies the next commit from a merge-base following the first-parent path toward either the base or head branch. This enables precise, commit-by-commit merge strategies.

### How It Works

Given a typical branching scenario:

```text
* d761e77 (feature) C: Feature development 1  
| * 0eb5947 (main)  B: Main branch development
|/
* cdddb51            A: Initial commit (merge-base)
```

#### Walk to Head

```bash
gh merge-base-next main feature
# Returns: d761e77 (C: Feature development 1)
```

The tool:

1. Finds the merge-base between `main` and `feature` → `cdddb51` (A)
2. Walks along the first-parent path from merge-base toward `feature`
3. Returns the next commit → `d761e77` (C)

### Advanced Scenarios

#### Complex Merge Structures

```text
*   f123456 (feature) Merge commit
|\
| * e789012           Feature work 2
* | d456789           Feature work 1
|/
* c123456             Common ancestor
```

The tool handles complex merge structures by following first-parent relationships, ensuring consistent behavior even with merge commits.

#### Multiple Commits on Branch

```text
* e789012 (feature) Latest feature work
* d456789           Previous feature work
* c123456           First feature work
| * b234567 (main)  Main development
|/
* a123456           Merge-base
```

For step-by-step merging:

- 1st iteration: `gh merge-base-next main feature` → `c123456`
- 2nd iteration: `gh merge-base-next c123456 feature` → `d456789`
- 3rd iteration: `gh merge-base-next d456789 feature` → `e789012`

## Usage

### Basic Usage

Find the next commit from merge-base to head:

```bash
gh merge-base-next <base> <head>
```

### Options

- `--repo, -R string`: Target repository in the format 'owner/repo' (optional)
- `--walk-to, -T string`: Specifies whether the next commit should walk to 'base' or 'head' (default: "head")
- `--format string`: Output format: {json}
- `--template, -t string`: Format JSON output using a Go template
- `--jq, -q expression`: Filter JSON output using a jq expression

### Examples

#### Find next commit from branch to another branch

```bash
gh merge-base-next main feature
```

This command finds the next commit from the merge-base of 'main' and 'feature' branches toward the 'feature' branch.

#### Find next commit toward base branch

```bash
gh merge-base-next main feature --walk-to base
```

This command finds the next commit from the merge-base toward the 'main' branch.

#### Specify target repository

```bash
gh merge-base-next main feature --repo owner/repo
```

#### Get JSON output

```bash
gh merge-base-next main feature --format json
```

#### Use with specific commit SHA

```bash
gh merge-base-next abc123 def456
```

## Motivation

In multi-branch development environments, I adopted a merge strategy that performs merges one commit at a time to minimize conflict resolution responsibilities. This tool is designed to support that workflow by identifying the specific commits to merge.

### The Challenge with Traditional Branch Merging

Typical branch merges often contain commits from multiple developers, requiring significant effort to resolve conflicts. When merging an entire feature branch at once, developers must resolve all conflicts simultaneously, which can be time-consuming and error-prone.

### One-Commit-at-a-Time Strategy

This tool enables a commit-by-commit merge strategy that:

- **Minimizes Conflict Scope**: By merging one commit at a time, conflicts are isolated and easier to resolve
- **Distributes Responsibility**: Each commit author is responsible for resolving conflicts in their specific changes
- **Enables Automation**: The merge process can be automated with targeted conflict resolution prompts
- **Maintains Safety**: Repository rulesets and GitHub Actions prevent disruptive operations, ensuring safe automated merging

### Technical Implementation

The tool returns the next commit on the first-parent path from the common ancestor to the target branch. It's implemented to work without checkout operations, enabling fast execution in GitHub Actions workflows.

This tool serves as the foundation for automated merge systems that maintain code quality while reducing manual merge overhead.

## Development

### Testing

```bash
# Run all tests
make test

# Run tests with coverage report
make test-coverage

# Run tests with JUnit report generation
make test-report
```

### Test Reports

The project generates comprehensive test reports in the `test-results/` directory:

- **JUnit XML Reports**: `junit.xml` for CI/CD integration
- **Coverage Reports**: HTML format coverage report (`coverage.html`)
- **Octocov Reports**: Coverage reporting via Octocov with badge generation  
- **Unit Tests**: All tests are now properly organized as unit tests within respective packages
- **Test Artifacts**: Uploaded to GitHub Actions for PR reviews

Test reports are automatically generated in CI/CD and can be viewed in:

- GitHub Actions summary page
- PR checks and status
- Artifacts download from workflow runs
