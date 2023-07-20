# gRPC in Golang Starter Code

## Compile `.proto` file
```
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld.proto
```

Running this command generates two files
1. `helloworld.pb.go`
2. `helloworld_grpc.pb.go`

## Run Server
```
/server $ go server.go
```

## Run Client
```
/client $ go client.go
```

# Reference
https://github.com/grpc/grpc-go/tree/master/examples/helloworld
