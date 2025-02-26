MAIN=cmd/balance/main.go
OUTPUT=balance

AIR=github.com/air-verse/air@latest
FORMATTER=mvdan.cc/gofumpt@latest
LINTER=github.com/golangci/golangci-lint/cmd/golangci-lint@latest
MOCKGEN=go.uber.org/mock/mockgen@latest
WIRE=github.com/google/wire/cmd/wire@latest

.PHONY: default
default: vet fix fmt lint test

.PHONY: vet
vet:
	@echo "go vet"
	@go vet ./...

.PHONY: fix
fix:
	@echo "go fix"
	@go fix ./...

.PHONY: fmt
fmt:
	@echo "go fmt"
	@go run $(FORMATTER) -l -w .

.PHONY: lint
lint:
	@echo "go lint"
	@go run $(LINTER) run

.PHONY: test
test:
	@echo "go test"
	@go test ./... -count=1

.PHONY: update
update:
	@go test ./... -update

.PHONY: air
air:
	@echo "go air"
	@go run $(AIR)

.PHONY: run
run:
	@echo "go run (without live reloading)"
	@go run github.com/tksasha/balance/cmd/balance

.PHONY: build
build:
	@echo "go build"
	@go build -o $(OUTPUT) $(MAIN)

.PHONY: clear
clear:
	@echo "go clear"
	@rm $(OUTPUT)

.PHONY: clean
clean: clear

.PHONY: gen
gen: wire mockgen

.PHONY: mockgen
mockgen:
	@go run $(MOCKGEN) \
		-source internal/app/cash/interfaces.go \
		-package mocks \
		-destination internal/app/cash/test/mocks/interfaces.mock.go
	@go run $(MOCKGEN) \
		-source internal/app/category/interfaces.go \
		-package mocks \
		-destination internal/app/category/test/mocks/interfaces.mock.go
	@go run $(MOCKGEN) \
		-source internal/app/item/interfaces.go \
		-package mocks \
		-destination internal/app/item/test/mocks/interfaces.mock.go
	@go run $(MOCKGEN) \
		-source internal/app/index/interfaces.go \
		-package mocks \
		-destination internal/app/index/test/mocks/interfaces.mock.go
	@go run $(MOCKGEN) \
		-source internal/backoffice/category/interfaces.go \
		-package mocks \
		-destination internal/backoffice/category/test/mocks/interfaces.mock.go
	@go run $(MOCKGEN) \
		-source internal/backoffice/cash/interfaces.go \
		-package mocks \
		-destination internal/backoffice/cash/test/mocks/interfaces.mock.go

.PHONY: wire
wire:
	@echo "wire gen"
	@go run $(WIRE) internal/wire/wire.go

.PHONY: migration # to create new migration
migration:
	@if [ -z "$(name)" ]; then echo "name is required"; exit 1; fi
	touch "internal/db/migrations/$(shell date "+%Y%m%d%H%M%S")_$(name).sql"
