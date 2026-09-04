# OfflinePush

[English](README.md)

OfflinePush 是面向 GeTui REST API v2 与基于令牌认证的 APNs 的 gRPC 服务套件 提供认证 推送 设备与标签管理 统计以及面向应用的代理包

## 依赖

- Go 1.16 或更高版本
- GeTui REST API v2 凭据
- 统计服务需要 MongoDB
- APNs 推送需要令牌认证密钥
- 修改协议时需要 `protoc` `protoc-gen-go` `protoc-gen-go-grpc` 与 `protoc-go-inject-tag`

## 配置

在仓库根目录创建 `conf.json` 此文件被 Git 忽略 并应包含 MongoDB GeTui 以及需要时的 APNs 凭据 所有服务从当前工作目录读取该文件

默认本地 gRPC 地址为 Auth `localhost:8080` Push `:8082` User `:8084` 和 Stats `:8086` 服务间使用不安全 gRPC 应仅部署在私有网络中

## 运行

先启动 Auth 再分别启动其他服务

```sh
go run ./server/auth
go run ./server/push
go run ./server/user
go run ./server/stats
```

服务协议位于 [`proto`](proto) APNs 当前仅支持单设备推送 其他 API 使用 GeTui

## 测试

```sh
go vet ./...
go test ./...
```

测试可能调用依赖服务和第三方平台 请使用非生产凭据并确保依赖可用

## 生成协议代码

```sh
./generatePB.sh
```

## 许可证

[MIT](LICENSE)
