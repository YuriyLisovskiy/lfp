# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying 
# file LICENSE or https://opensource.org/licenses/MIT

all: test target install

target:
	@bash ./build.sh

install:
	@bash ./install.sh

PACKAGES = ./src

coverage: test
	@echo Generating coverage report...
	@go tool cover -html coverage.out -o coverage.html
	@echo Done.

test:
	@echo Running tests...
	@go test -v -timeout 1h -covermode=count -coverprofile=coverage.out ${PACKAGES}
	@echo Done.

clean:
	@rm -rf bin/ coverage.out coverage.html
