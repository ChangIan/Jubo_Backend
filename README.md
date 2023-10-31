# Jubo - Backend

## Environment
- Docker 24.0.2
- Go 1.20

## Initialize
```bash
# Run Go Init
$ go mod init changian.com/jubo
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/lib/pq
$ go get -u github.com/golang-jwt/jwt
$ go get -u gopkg.in/yaml.v2
# Build Local Image on Docker
$ docker build -t "jubobackend:0.0.1" -t "jubobackend:latest" .
# Run App on Docker
$ docker network create -d bridge common-network
$ docker run -d --restart=always --network=common-network --name jubobackend -p 8899:8899 jubobackend:latest