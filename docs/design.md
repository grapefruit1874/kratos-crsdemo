# CRS

## CLI

```markdown
创建xxx微服务模块：kratos new app/xxx --nomod -r https://gitee.com/go-kratos/kratos-layout.git
添加proto文件：kratos proto add api/xxx/v1/xxx.proto
生成proto代码：kratos proto client api/xxx/v1/xxx.proto
生成server代码：kratos proto server api/xxx/v1/xxx.proto -t app/xxx/internal/service
```

```markdown
# 依赖注入
app/xxx/internal/service/service.go 依赖注入
 biz新建xxx.go
app/xxx/internal/biz/biz.go 依赖注入
 data新建xxx.go
app/xxx/internal/data/data.go 依赖注入+orm初始化

# 实例化http和gRPC
app/xxx/internal/server，修改http.go,grpc.go,server.go
v1.RegisterXXXHTTPServer(srv, student)
v1.RegisterXXXServer(srv, student)
```

## user:8000

## member:8001

## order:8003
