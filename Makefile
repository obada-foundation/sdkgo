lint:
	golangci-lint --config .golangci.yml run --print-issued-lines --out-format=github-actions ./...

test:
	go test ./... -v
