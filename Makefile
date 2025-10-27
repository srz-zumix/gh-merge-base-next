EXTENSION_NAME=merge-base-next

help: ## Display this help screen
	@grep -E '^[a-zA-Z][a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed -e 's/^GNUmakefile://' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

install: ## install gh extention
	gh extension remove "srz-zumix/gh-${EXTENSION_NAME}" || :
	gh extension remove "${EXTENSION_NAME}" || :
	gh extension install .


install-released:
	gh extension remove "${EXTENSION_NAME}" || :
	gh extension install "srz-zumix/gh-${EXTENSION_NAME}"

build: ## build the binary
	go build -o gh-${EXTENSION_NAME}

test: ## run all tests
	go test -v ./...

test-integration: build ## run integration tests
	go test -v ./integration_test/...

test-coverage: ## run tests with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-report: ## run tests with JUnit report generation
	@command -v go-junit-report >/dev/null 2>&1 || go install github.com/jstemmer/go-junit-report/v2@latest
	mkdir -p test-results
	go test -v -coverprofile=coverage.out -covermode=atomic ./cmd/ ./merge-base-next/ ./version/ 2>&1 | tee test-output.txt
	go-junit-report -in test-output.txt -out test-results/junit.xml
	go tool cover -html=coverage.out -o test-results/coverage.html
	@echo "Unit test report generated in test-results/"

test-integration-report: build ## run integration tests with JUnit report
	@command -v go-junit-report >/dev/null 2>&1 || go install github.com/jstemmer/go-junit-report/v2@latest
	mkdir -p test-results
	go test -v ./integration_test/... 2>&1 | tee integration-test-output.txt || true
	go-junit-report -in integration-test-output.txt -out test-results/integration-junit.xml
	@echo "Integration test report generated in test-results/"

test-all-report: ## run all tests with comprehensive reporting
	@$(MAKE) test-report
	@$(MAKE) test-integration-report
	@echo "All test reports generated in test-results/"

octocov-local: ## run octocov locally (requires octocov to be installed)
	@command -v octocov >/dev/null 2>&1 || (echo "octocov is not installed. Install with: go install github.com/k1LoW/octocov@latest" && exit 1)
	octocov

clean: ## clean build artifacts and test files
	rm -f gh-${EXTENSION_NAME}
	rm -f coverage.out coverage.html
	rm -f test-output.txt integration-test-output.txt
	rm -rf test-results/
	rm -f go.work

go-work:
	# (cd .. && gh repo clone srz-zumix/go-gh-extension)
	ln -snf ../go-gh-extension go-gh-extension
	# go work init
	go work use .
	go work use ./go-gh-extension
	go work sync
