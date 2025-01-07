# Notification Service gRPC# github.com/harlitad/notitication-service

## Generate Proto
```
protoc --go_out=./service --go-grpc_out=. ./proto/service.proto
```

### Run Server
```
cp .env-example .env
go mod tidy
go run main.go
```