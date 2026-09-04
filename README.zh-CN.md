# OfflinePush

[English documentation](README.md)

OfflinePush 是一套基于 Go 和 gRPC 的个推 REST API v2 服务。它提供推送投递、设备与标签管理、推送统计，以及面向业务侧的轻量推送代理。同时支持 APNs Token 鉴权和单设备推送。

## 架构

```text
业务应用 --> proxy 包 --> push 服务 --> 个推或 APNs

user 服务  ---+
stats 服务 ---+--> auth 服务 --> 个推或 APNs 鉴权提供方
push 服务  ---+
```

服务使用固定的本地地址并通过非加密 gRPC 通信。应将它们部署在同一主机，或限制在私有网络边界内。

| 服务 | 地址 | 职责 |
| --- | --- | --- |
| Auth | `localhost:8080` | 获取并刷新个推 Token，生成 APNs JWT |
| Push | `localhost:8082` | 单推、批量推送、列表推送、全量推送、标签推送与任务管理 |
| User | `localhost:8084` | 别名、标签、黑名单、设备和用户管理接口 |
| Stats | `localhost:8086` | 推送与用户统计接口，并将结果写入 MongoDB |

## 环境要求

- Go 1.16 或更高版本
- 已开通 REST API v2 的个推应用
- MongoDB 用于 Stats 服务
- 仅在需要 APNs 推送时配置 APNs Token 鉴权密钥
- 仅在重新生成 protobuf 代码时需要 `protoc`、`protoc-gen-go`、`protoc-gen-go-grpc` 和 `protoc-go-inject-tag`

## 配置

每个服务都会从当前工作目录加载 `conf.json`。该文件包含敏感信息，不应提交到版本库。APNs 的 `key` 字段直接填写 PKCS#8 PEM 内容，而不是密钥文件路径。

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

`agent_id` 和 `bundle_id` 用于选择每次请求对应的推送提供方。gRPC 请求必须携带与目标提供方配置一致的这组值。对于 APNs，`bundle_id` 还会作为 `apns-topic` 请求头发送。只有 `env` 精确等于 `Production` 时才使用生产网关，其他值均使用开发网关。

## 启动

必须先启动 Auth 服务。Push、User 和 Stats 服务会在初始化时连接 Auth 服务，若五秒内不可用会启动失败。

```bash
go run ./server/auth
```

在其他终端启动其余服务：

```bash
go run ./server/push
go run ./server/user
go run ./server/stats
```

仅构建全部包而不启动服务：

```bash
go build ./...
```

## API

[`proto/`](proto) 下的定义是 gRPC 契约，对应生成的 Go 客户端代码也已提交在该目录。

| 服务 | 接口能力 |
| --- | --- |
| `Auth` | 获取或失效推送提供方 Token |
| `Push` | 单推和批量推送、可复用消息任务、列表推送、全量和标签推送、停止删除查询任务与查看详情 |
| `User` | 别名绑定和查询、标签和黑名单管理、用户设备状态、用户信息、角标、人数统计和 CID/设备 Token 绑定 |
| `Stats` | 任务和任务组报表、推送数量、每日推送和用户报表，以及 24 小时在线用户数据 |

`proxy` 包通过 `MsgConverter` 接口适配业务消息。其 `Unicast`、`Multicast` 和 `Broadcast` 方法在完成转换后调用 Push 服务。

## 提供方支持范围

### 个推

Push、User 和 Stats 服务实现对应的个推 REST API v2 接口。批量和列表推送在顶层个推状态码非零时仍可能包含部分成功数据，服务会保留并返回这些逐项结果。

### APNs

APNs 当前仅支持通过 CID 调用 `PushToSingle`。请求必须恰好包含一个 CID，并提供 iOS 推送通道。别名推送、批量推送、任务创建、列表/全量/标签推送、任务管理、用户接口和统计接口尚未实现，调用时会返回不支持错误。

## 生成 Protobuf 代码

修改 `.proto` 文件后，在仓库根目录执行：

```bash
./generatePB.sh
```

该命令会调用 `proto/protoc.sh`，并重新注入推送提供方所需的 JSON 标签。提交前应检查生成文件的变更。

## 开发说明

- `go test ./...` 包含集成性质的测试，需要有效的提供方凭据和依赖服务。不要使用生产凭据执行它。
- `go vet ./...` 只进行静态检查，不会向提供方发送请求。
- MongoDB 当前仅由 Stats 服务用于保存获取到的报表数据。

## 许可证

使用 [MIT License](LICENSE)。
