# Go-Zero 自定义模板

## 目录结构

```
deploy/goctl/
└── 1.10.1/                    # goctl 版本号
    ├── api/                   # API 模板
    ├── docker/               # Docker 模板
    ├── gateway/              # Gateway 模板
    ├── kube/                 # Kubernetes 模板
    ├── model/                # 数据库模型模板
    ├── mongo/                # MongoDB 模板
    ├── newapi/               # 新 API 模板
    └── rpc/                  # RPC 模板
```

## 使用方式

### 1. 使用自定义模板生成代码

```bash
# 生成 API
goctl api go -api user.api -dir . --home deploy/goctl/1.10.1

# 生成 Model
goctl model mysql ddl -src user.sql -dir . --home deploy/goctl/1.10.1

# 生成 RPC
goctl rpc protoc user.proto --go_out=./ --go-grpc_out=./ --zrpc_out=./ --home deploy/goctl/1.10.1
```

### 2. 配置环境变量（推荐）

设置 `GOCTL_HOME` 环境变量，这样每次使用时就不需要指定 `--home`：

**Windows PowerShell:**
```powershell
# 临时设置
$env:GOCTL_HOME="deploy/goctl/1.10.1"

# 永久设置（添加到用户环境变量）
[System.Environment]::SetEnvironmentVariable('GOCTL_HOME', 'C:\Users\林江涛\Desktop\goZeroDemo\deploy\goctl\1.10.1', 'User')
```

**Linux/Mac:**
```bash
export GOCTL_HOME=deploy/goctl/1.10.1
```

设置环境变量后，可以直接使用：
```bash
goctl api go -api user.api -dir .
```

## 自定义模板示例

### 示例1：自定义 Handler 返回格式

编辑 `api/handler.tpl`，修改错误返回格式：

```go
func {{.HandlerName}}Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			{{.PacketName}}.Fail(w, http.StatusBadRequest, "请求参数错误")
			return
		}

		l := {{.PacketName}}.New{{.HandlerName}}Logic(r.Context(), svcCtx)
		resp, err := l.{{.HandlerName}}(&req)

		if err != nil {
			{{.PacketName}}.Fail(w, http.StatusInternalServerError, err.Error())
		} else {
			{{.PacketName}}.Success(w, resp)
		}
	}
}
```

### 示例2：自定义 Logic 模板

编辑 `api/logic.tpl`，添加默认日志：

```go
type {{.LogicType}} struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func New{{.LogicType}}(ctx context.Context, svcCtx *svc.ServiceContext) *{{.LogicType}} {
	return &{{.LogicType}}{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *{{.LogicType}}) {{.MethodName}}(req *types.{{.RequestType}}) (*types.{{.ResponseType}}, error) {
	l.Infof("{{.MethodName}} request: %+v", req)
	
	// TODO: add your logic here
	
	l.Infof("{{.MethodName}} response success")
	return &types.{{.ResponseType}}{}, nil
}
```

## 常用命令

```bash
# 查看当前 goctl 版本
goctl --version

# 初始化默认模板
goctl template init --home deploy/goctl/1.10.1

# 查看模板帮助
goctl template -h

# 查看当前使用的模板路径
goctl template info
```

## 模板变量说明

### API 模板可用变量

- `{{.Name}}` - 服务名
- `{{.Prefix}}` - 路由前缀
- `{{.Groups}}` - 分组信息
- `{{.Import}}` - 导入路径
- `{{.HandlerName}}` - Handler 名称
- `{{.LogicType}}` - Logic 类型
- `{{.MethodName}}` - 方法名
- `{{.RequestType}}` - 请求类型
- `{{.ResponseType}}` - 响应类型
- `{{.PacketName}}` - 包名

### Model 模板可用变量

- `{{.Table}}` - 表名
- `{{.TableSnake}}` - 下划线格式表名
- `{{.PrimaryKey}}` - 主键信息
- `{{.PrimaryKeyPrefix}}` - 主键前缀
- `{{.Fields}}` - 字段信息
- `{{.Import}}` - 导入路径

## 最佳实践

1. **先备份**：修改模板前先备份原始模板
2. **小步迭代**：逐个修改模板，及时验证
3. **团队共享**：将模板目录加入 git 版本控制，团队共享
4. **版本控制**：每个 goctl 版本使用独立目录

## 示例模板修改

参考项目中的 `sendemailcodehandler.go`，我们可以修改 `api/handler.tpl` 来统一错误返回格式。
