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
```

## Output

```bash
.
|   .gitignore
|   app.go
|   Dockerfile
|   go.mod
|   go.sum
|   jubo.exe
|   README.md
|   tree.txt
|   
+---.vscode
|       launch.json
|       
+---config
|       config.docker.yaml
|       config.yaml
|       
+---controller
|   \---v1
|           ordersController.go
|           patientController.go
|           
+---helper
|       postgresql.go
|       
+---middleware
|       cors.go
|       jwtToken.go
|       
+---model
|   |   orders.go
|   |   patient.go
|   |   transactionLog.go
|   |   
|   +---base
|   |       response.go
|   |       
|   \---common
|           config.go
|           operatorClaims.go
|           
+---repository
|   \---v1
|           orders.go
|           patient.go
|           transactionLog.go
|           
+---router
|       router.go
|       
+---service
|   \---v1
|           serviceOrders.go
|           servicePatient.go
|           
+---startup
|       global.go
|       
\---utility
        config.go
```
