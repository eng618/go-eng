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

# -----------------------------------------------------------------------------
# Modules support

deps-reset:
	git checkout -- go.mod
	go mod tidy

tidy:
	go mod tidy

deps-list:
	go list -m -u -mod=readonly all

deps-upgrade:
	go get -u -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all
