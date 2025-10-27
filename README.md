# gh-merge-base-next

[![Build and Test](https://github.com/srz-zumix/gh-merge-base-next/actions/workflows/build.yml/badge.svg)](https://github.com/srz-zumix/gh-merge-base-next/actions/workflows/build.yml)

gh-merge-base-next is a tool to find the next commit in a merge base.

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
