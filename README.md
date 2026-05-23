# goZeroDemo

# 项目启动

社交平台

## rpc

```
cd rpc/user

go run user.go
```

## Api

```
cd api
go run gozeroapi.go
```

## ETCD

```
etcd
```

构建api

```
 goctl api go -api desc/all.api -dir .
```

grpc

```
goctl rpc protoc user-grpc.proto --go_out=. --go-grpc_out=. --zrpc_out=.
```

swagger

```
goctl api swagger -api desc/all.api -dir ./doc
