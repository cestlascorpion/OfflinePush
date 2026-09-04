# OfflinePush

[中文文档](README.zh-CN.md)

OfflinePush is a Go gRPC service suite for GeTui REST API v2. It provides push delivery, device and tag management, delivery statistics, and a small application-facing push proxy. APNs token authentication and single-device delivery are also available.

## Architecture

```text
application --> proxy package --> push service --> GeTui or APNs

user service  ---+
stats service ---+--> auth service --> GeTui or APNs token provider
push service  ---+
```

The services use fixed local addresses and communicate through insecure gRPC. Run them on the same host or place them behind a private network boundary.

| Service | Address | Responsibility |
| --- | --- | --- |
| Auth | `localhost:8080` | Obtains and refreshes GeTui tokens and creates APNs JWTs |
| Push | `localhost:8082` | Single, batch, list, app, tag, and task push APIs |
| User | `localhost:8084` | Alias, tag, blacklist, device, and user-management APIs |
| Stats | `localhost:8086` | Push and user reporting APIs, with MongoDB persistence |

## Requirements

- Go 1.16 or later
- A GeTui application with REST API v2 credentials
- MongoDB for the stats service
- An APNs token-authentication key only when APNs delivery is required
- `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc`, and `protoc-go-inject-tag` only when regenerating protobuf code

## Configuration

Every server loads `conf.json` from its current working directory. Keep this file out of version control. The APNs `key` field contains the PKCS#8 PEM content itself, not a path to a key file.

```json
{
  "mongo": {
    "name": "offlinepush",
    "url": "mongodb://127.0.0.1:27017",
    "database": "offlinepush",
    "auth_collection": "auth",
    "stats_collection": "stats",
    "pool_size": 16
  },
  "getui": {
    "agent_id": "getui",
    "bundle_id": "com.example.app",
    "app_id": "GETUI_APP_ID",
    "app_key": "GETUI_APP_KEY",
    "master_secret": "GETUI_MASTER_SECRET"
  },
  "apns": {
    "agent_id": "apns",
    "bundle_id": "com.example.iosapp",
    "env": "Production",
    "key": "-----BEGIN PRIVATE KEY-----\nAPNS_PRIVATE_KEY_CONTENT\n-----END PRIVATE KEY-----",
    "key_id": "APNS_KEY_ID",
    "team_id": "APPLE_TEAM_ID"
  }
}
```

`agent_id` and `bundle_id` identify the push provider selected by each request. Use the same pair in gRPC requests that was configured for the intended provider. For APNs, `bundle_id` is also sent as the `apns-topic` header. `env` must be exactly `Production` for the production gateway; every other value selects the development gateway.

## Run

Start the auth service first. The push, user, and stats services connect to it during initialization and fail after a five-second connection timeout if it is unavailable.

```bash
go run ./server/auth
```

Start the remaining services in separate terminals:

```bash
go run ./server/push
go run ./server/user
go run ./server/stats
```

Build all packages without starting services:

```bash
go build ./...
```

## APIs

The protobuf definitions in [`proto/`](proto) are the gRPC contract and generated Go clients are committed in the same directory.

| Service | Operations |
| --- | --- |
| `Auth` | Get or invalidate provider tokens |
| `Push` | Single and batch delivery, reusable message tasks, list delivery, app and tag broadcast, task stop/remove/query/detail |
| `User` | Alias binding and lookup, tag and blacklist management, user/device status, user info, badges, counts, and CID/device-token binding |
| `Stats` | Task and group reports, push counts, daily push/user reports, and 24-hour online-user data |

The `proxy` package adapts application messages through the `MsgConverter` interface. Its `Unicast`, `Multicast`, and `Broadcast` methods call the push service after conversion.

## Provider Support

### GeTui

The push, user, and stats services implement the corresponding GeTui REST API v2 operations. Batch and list push responses may contain partial-success data even when the top-level GeTui code is non-zero; that response data is preserved for callers.

### APNs

APNs currently supports only `PushToSingle` by CID. The request must contain exactly one CID and an iOS push channel. Alias delivery, batch delivery, task creation, list/app/tag delivery, task management, user APIs, and stats APIs are not implemented for APNs and return an unsupported-operation error.

## Generate Protobuf Code

After changing a `.proto` file, regenerate the Go bindings from the repository root:

```bash
./generatePB.sh
```

This invokes `proto/protoc.sh` and also reapplies the JSON tags expected by the provider payloads. Review generated changes before committing them.

## Development Notes

- `go test ./...` contains integration-style tests and requires valid provider credentials plus dependent services. Do not run it against production credentials.
- `go vet ./...` performs a static check without sending provider requests.
- MongoDB is currently used by the stats service to retain fetched report data.

## License

Licensed under the [MIT License](LICENSE).
