# Available architecture combinations for Go: https://golang.org/doc/install/source#environment

VERSION ?= dev

linux/amd64:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -buildmode=c-shared -o out_fluent-bit-test-linux-amd64-${VERSION}.so .

linux/arm64:
	CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc CXX=aarch64-linux-gnu-g++ go build -buildmode=c-shared -o out_fluent-bit-test-linux-arm64-${VERSION}.so .

clean:
	rm -rf *.so *.h *~