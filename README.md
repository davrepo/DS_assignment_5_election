# Auction House

This is a distributed system for an auction house, implemented in Go using gRPC.

## Features

-Streaming RPCs for real-time updates.

- Fault-tolerance using the ring algorithm.

## Usage

1. Install the dependencies:

````bash
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go
GitHub Copilot: Sure, here's a basic README for your project:

```markdown
# Auction House

This is a distributed system for an auction house, implemented in Go using gRPC.

## Features

- Bidirectional streaming RPCs for real-time updates.
- Fault-tolerance using the ring algorithm.
- Client ID flag with random default value.

## Usage

1. Install the dependencies:

```bash
go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go
````

2. Generate Go code from the proto file:

```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/auctionhouse.proto
```

3. Run the server:

You can use the replicamanger to start a server

```bash
go run replicamanager/replicamanager.go 3001
```

To create more replicas run the command again with a new port which should be the next port after 3001 as we are using the ring algo

```bash
go run replicamanager/replicamanager.go 3002
```

4. Run the client:

Then you can run the client on the same port as the primary server

```bash
go run client/client.go 3001
```
