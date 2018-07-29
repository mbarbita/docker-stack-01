 #!/bin/bash
echo ""
echo "run: go get -v file.go"
echo "run: go build -v file.go"
docker run -it --rm -v "$PWD":/usr/src/go-mysql-grafana -w /usr/src/go-mysql-grafana golang /bin/bash
