run: build
	./kreator

build: fmt
	go build -o kreator cmd/kreator/main.go

test: fmt
	go test -v

fmt:
	go fmt ./cmd/kreator/... ./internal/.../... ./test/...