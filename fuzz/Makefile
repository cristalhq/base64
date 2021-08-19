.PHONY: install
install:
	@go install \
		github.com/dvyukov/go-fuzz/go-fuzz@latest \
		github.com/dvyukov/go-fuzz/go-fuzz-build@latest

.PHONY: build
build:
	go-fuzz-build

.PHONY: fuzz
fuzz:
	go-fuzz -bin=./base64-fuzz.zip -workdir=workdir
