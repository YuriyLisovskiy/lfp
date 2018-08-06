#!/usr/bin/env bash

if hash go 2>/dev/null
then
    echo Building the binary for target platform...
	go run ./build.go
    echo Done.
else
    echo build.sh: can\'t build a binary$'\n'$'\t'reason: golang is not installed, please install it and then try again
fi
