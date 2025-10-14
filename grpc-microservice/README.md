

# gRPC Microservice

A simple gRPC-based microservice written in Go, demonstrating communication between a client and server using protocol buffers.

## 🧩 Features
- Implements two RPC methods: `SayHello` and `GetTime`.
- Uses `.proto` definition with generated Go code via `protoc`.
- Includes a gRPC client and server example.
- Demonstrates use of context and deadlines.

## 📁 Project Structure
```
grpc-microservice/
├── api/
│   └── gen/               # Generated Go files from .proto
├── cmd/
│   ├── client/            # gRPC client
│   └── server/            # gRPC server
├── api/hello.proto        # Protocol buffer definition
└── README.md
```

## ⚙️ Setup
1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Generate gRPC code:
   ```bash
   protoc --go_out=api/gen --go_opt=paths=source_relative \
          --go-grpc_out=api/gen --go-grpc_opt=paths=source_relative \
          api/hello.proto
   ```

## 🚀 Run the Server
```bash
go run ./cmd/server
```

## 💬 Run the Client
```bash
go run ./cmd/client
```

## ✅ Example Output
```
SayHello: Hello, Turman!
SayTime: 2025-10-14T12:34:56Z
```

## 🧠 Learning Focus
- Practice building gRPC microservices in Go.
- Understand .proto-based code generation.
- Learn bidirectional client-server communication.
- Explore reflection and debugging with grpcurl.