build:
	@echo "Building..."
	go build -v ./...

lint:
	@echo	"Linting..."
	golangci-lint run --fix

test:
	@echo "Testing..."
	CGO_ENABLED=1 go test ./... -race

validate: lint test build

release:
	@echo "Starting release..."
	goreleaser release --skip-publish

changelog:
	@echo "Updating changelog"
	git-chglog -o CHANGELOG.md
