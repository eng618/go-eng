build:
	@echo "Building..."
	go build -v ./...

lint:
	@echo	"Linting..."
	golangci-lint run --fix

test:
	@echo "Testing..."
	CGO_ENABLED=1 go test ./... -race -count=1

validate: lint test build

publish: release changelog

release:
	@echo "Starting release..."
	goreleaser release --clean --skip=publish

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

mod-contexts:
	@echo "updating examples/contexts"
	cd examples/contexts && go get -u -v ./... && go mod tidy

mod-generics:
	@echo "updating examples/generics"
	cd examples/generics && go get -u -v ./... && go mod tidy

mod-web-service-gin:
	@echo "updating examples/web-service-gin"
	cd examples/web-service-gin && go get -u -v ./... && go mod tidy

deps-upgrade: mod-contexts mod-generics mod-web-service-gin
	@echo "updating base project"
	go get -u -v ./...
	go mod tidy

deps-cleancache:
	go clean -modcache

list:
	go list -mod=mod all
