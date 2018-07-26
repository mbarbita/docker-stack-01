 #!/bin/bash
echo ""
echo "go get -v"
echo "go build -v"
docker run -it --rm -v "$PWD":/usr/src/go-mysql-grafana -w /usr/src/go-mysql-grafana golang /bin/bash
