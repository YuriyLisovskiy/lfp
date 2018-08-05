# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying 
# file LICENSE or https://opensource.org/licenses/MIT

BINARY = lofp
FLAGS = main.go

all: test target

target:
	@echo Building the binary for target platform...
	@go build -o bin/${BINARY} ${FLAGS}
	@echo Done.

PACKAGES = ./src/...

coverage:
	@echo Running tests...
	@go test -v -timeout 1h -covermode=count -coverprofile=coverage.out ${PACKAGES}
	@echo Generating coverage report...
	@go tool cover -html coverage.out -o coverage.html
	@echo Done

test:
	@echo Running tests...
	@go test -timeout 1h -covermode=count -coverprofile=coverage.out ${PACKAGES}
	@echo Done

clean:
	@-rm -rf bin/
