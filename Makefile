# Copyright (c) 2018 Yuriy Lisovskiy
# Distributed under the MIT software license, see the accompanying 
# file LICENSE or https://opensource.org/licenses/MIT

.PHONY: all clean install uninstall build coverage test

PACKAGES = ./src/lfp
COVER = coverage.out
COVER_REPORT = coverage.html
PREFIX = /usr/local/bin/lfp
PREFIX_UPD = /usr/local/bin/lfp-updater
TARGET = ./bin/lfp
TARGET_UPD = ./bin/lfp-updater

all: clean build

clean:
	@echo cleaning up redundant files...
	@rm -rf ./bin/ $(COVER) $(COVER_REPORT)

install:
	@echo Installing lfp...
	@sudo cp $(TARGET) $(PREFIX)
	@sudo cp $(TARGET_UPD) $(PREFIX_UPD)
	@echo Done.

uninstall:
	@echo Uninstalling lfp...
	@sudo rm -rf $(PREFIX)
	@sudo rm -rf $(PREFIX_UPD)
	@echo Done.

build:
	@bash ./scripts/build.sh

coverage: test
	@echo Generating coverage report...
	@go tool cover -html $(COVER) -o $(COVER_REPORT)
	@echo Done.

test:
	@echo Running tests...
	@go test -v -timeout 1h -covermode=count -coverprofile=$(COVER) $(PACKAGES)
	@echo Done.
