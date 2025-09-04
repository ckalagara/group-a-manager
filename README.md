# group-a-manager

# dev Setup 

## install
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"

```

## protobuff generation
```
protoc --go_out=. --go-grpc_out=. proto/manager.proto
```

## build
```
docker build -t group-a-manager . 

```
## run
```
docker run -p 50051:50051 group-a-manager
2025/09/04 02:07:27 Starting gRPC server on port 50051: 2025-09-04 02:07:27.694568221 +0000 UTC m=+0.002126292
2025/09/04 02:07:27 Creating gRPC server: 2025-09-04 02:07:27.697590429 +0000 UTC m=+0.005148501
2025/09/04 02:07:27 Serving gRPC server: 2025-09-04 02:07:27.697645346 +0000 UTC m=+0.005203417

```

## grpcurl

### Reflections
```
grpcurl -plaintext localhost:50051 list          
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
manager.Manager

```

### Health

```
grpcurl -plaintext \
  -proto proto/manager.proto \
  -import-path ./proto \
  -d '{}' \
  localhost:50051 \
  manager.Manager/Health 
{
  "status": "Up"
}
````