build:
	@go build -o ./bin/gdfs ./cmd/gDFS/main.go

run: build
	@./bin/gdfs

test:
	@go test ./... -v
