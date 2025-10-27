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

test-coverage: ## run tests with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

test-report: ## run tests with JUnit report generation
	@command -v go-junit-report >/dev/null 2>&1 || go install github.com/jstemmer/go-junit-report/v2@latest
	mkdir -p test-results
	go test -v -coverprofile=coverage.out -covermode=atomic ./... 2>&1 | tee test-output.txt
	go-junit-report -in test-output.txt -out test-results/junit.xml
	go tool cover -html=coverage.out -o test-results/coverage.html
	@echo "Unit test report generated in test-results/"

octocov-local: ## run octocov locally (requires octocov to be installed)
	@command -v octocov >/dev/null 2>&1 || (echo "octocov is not installed. Install with: go install github.com/k1LoW/octocov@latest" && exit 1)
	octocov

clean: ## clean build artifacts and test files
	rm -f gh-${EXTENSION_NAME}
	rm -f coverage.out coverage.html
	rm -f test-output.txt
	rm -rf test-results/
	rm -f go.work

go-work:
	# (cd .. && gh repo clone srz-zumix/go-gh-extension)
	ln -snf ../go-gh-extension go-gh-extension
	# go work init
	go work use .
	go work use ./go-gh-extension
	go work sync
