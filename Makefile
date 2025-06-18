all: 
	CGO_ENABLED=0 go build -ldflags "-X 'github.com/takei0107/acrun/internal/lang.CJson=$$(cat internal/lang/filetypes/c.json)'" -o /tmp/acrun main.go
