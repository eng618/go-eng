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

publish: release changelog

release:
	@echo "Starting release..."
	goreleaser release --clean --skip-publish

changelog:
	@echo "Updating changelog"
	git-chglog -o CHANGELOG.md
	git add --update
	git commit -m "chore: update changelog [skip ci]"
