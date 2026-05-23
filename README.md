# goZeroDemo

基于 go-zero 框架构建的社交平台后端服务，采用微服务架构，支持用户注册、登录、个人信息管理等功能。

## 技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| Go | 1.26+ | 编程语言 |
| go-zero | 1.10.1 | 微服务框架 |
| gRPC | 1.81+ | 服务间通信 |
| Etcd | 3.5+ | 服务注册与发现 |
| MySQL | 8.0+ | 数据库 |
| JWT | - | 身份认证 |

## 项目结构

```
goZeroDemo/
├── api/                    # API 网关服务
│   ├── desc/               # API 描述文件
│   ├── etc/                # 配置文件
│   ├── internal/           # 内部代码
│   │   ├── config/         # 配置定义
│   │   ├── handler/        # HTTP 处理器
│   │   ├── logic/          # 业务逻辑
│   │   ├── svc/            # 服务上下文
│   │   └── pkg/            # 工具函数
│   └── gozeroapi.go        # API 入口文件
├── rpc/                    # gRPC 服务
│   └── user/               # 用户服务
│       ├── etc/            # 配置文件
│       ├── internal/       # 内部代码
│       ├── user-grpc/      # gRPC 生成代码
│       └── user.go         # RPC 入口文件
└── README.md               # 项目说明
```

## 环境要求

- Go 1.26+
- MySQL 8.0+
- Etcd 3.5+

## 快速开始

### 1. 启动 Etcd

```bash
# 使用 Docker（推荐）
docker run -d --name etcd -p 2379:2379 -p 2380:2380 quay.io/coreos/etcd:v3.5.2

# 或本地启动
etcd
```

### 2. 配置数据库

确保 MySQL 已运行，并创建数据库：

```sql
CREATE DATABASE IF NOT EXISTS zero CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 3. 启动 gRPC 服务

```bash
cd rpc/user
go run user.go
```

### 4. 启动 API 服务

```bash
cd api
go run gozeroapi.go
```

## 配置说明

### API 配置 (`api/etc/gozeroapi-api.yaml`)

```yaml
Name: goZeroApi
Host: 0.0.0.0
Port: 8888

# gRPC 服务配置
UserRpc:
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: user.rpc

# JWT 认证配置
Auth:
  AccessSecret: "your-secret-key"
  AccessExpire: 86400
```

### RPC 配置 (`rpc/user/etc/user.yaml`)

```yaml
Name: user.rpc
ListenOn: 0.0.0.0:8080

# Etcd 注册配置
Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: user.rpc

# MySQL 配置
MySQLConf:
  Enable: true
  User: root
  Password: your-password
  Host: 127.0.0.1
  Port: 3306
  Database: zero
```

## API 接口

### 用户模块

| 接口 | 方法 | 路径 | 认证 |
|------|------|------|------|
| 注册 | POST | /api/user/register | 否 |
| 登录 | POST | /api/user/login | 否 |
| 获取个人信息 | POST | /api/user/my/detail | 是 |

### 文章模块

| 接口 | 方法 | 路径 | 认证 |
|------|------|------|------|
| 获取文章详情 | GET | /api/article/:id | 否 |

## 代码生成

### 生成 API 代码

```bash
cd api
goctl api go -api desc/all.api -dir .
```

### 生成 gRPC 代码

```bash
cd rpc/user
goctl rpc protoc user-grpc.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

### 生成 Swagger 文档

```bash
cd api
goctl api swagger -api desc/all.api -dir ./doc
```

## 服务注册与发现

本项目使用 Etcd 实现服务注册与发现：

1. **服务注册**：gRPC 服务启动时自动注册到 Etcd
2. **服务发现**：API 服务通过 Etcd 获取 gRPC 服务地址
3. **负载均衡**：go-zero 内置负载均衡支持

## 开发规范

- 代码风格遵循 Go 官方标准
- 使用 goctl 生成代码骨架
- 错误处理使用 status.Error (gRPC) / httpx.Error (HTTP)
- 日志使用 go-zero 的 logx 包

## 许可证

MIT License
