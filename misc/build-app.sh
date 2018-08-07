#!/bin/bash
echo ""
echo "run: go get -d -v"
echo "run: go build -v"
docker run -it --rm -v "$PWD":/usr/src/go-misc -w /usr/src/go-misc golang /bin/bash
