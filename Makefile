.PHONY: all
all: lint test build

.PHONY: build
build:
	CGO_ENABLED=0 go build -x -trimpath -ldflags "-s -w -X 'github.com/takei0107/acrun/internal/lang.CJson=$$(cat internal/lang/filetypes/c.json)'" -o acrun

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: lint
lint:
	go vet .

.PHONY: test
test:
	go test ./...

.PHONY: docker-image
docker-image:
	docker build -t acrun:latest .
