# OfflinePush

[Simplified Chinese](README.zh-CN.md)

OfflinePush is a gRPC service suite for GeTui REST API v2 and APNs token-based delivery. It includes authentication, push, device and tag management, statistics, and an application-facing proxy package.

## Requirements

- Go 1.16 or later
- GeTui REST API v2 credentials
- MongoDB for the statistics service
- An APNs token-authentication key only for APNs delivery
- `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc`, and `protoc-go-inject-tag` only when changing the protocol

## Configuration

Create `conf.json` in the repository root. It is intentionally ignored by Git and must contain MongoDB, GeTui, and, when applicable, APNs credentials. All services read this file from their current working directory.

The default local gRPC addresses are `localhost:8080` (Auth), `:8082` (Push), `:8084` (User), and `:8086` (Stats). Services use insecure gRPC and should remain on a private network.

## Run

Start Auth first, then start the remaining services in separate terminals:

```sh
go run ./server/auth
go run ./server/push
go run ./server/user
go run ./server/stats
```

The service contracts are in [`proto`](proto). APNs currently supports single-device push only; the remaining APIs use GeTui.

## Test

```sh
go vet ./...
go test ./...
```

Tests can call dependent services and providers. Use non-production credentials and ensure their dependencies are available.

## Generate Protocol Code

Install `protoc-go-inject-tag` before generating protocol code:

```sh
go install github.com/favadi/protoc-go-inject-tag@latest
```

```sh
./genPb.sh
```

## License

[MIT](LICENSE)
