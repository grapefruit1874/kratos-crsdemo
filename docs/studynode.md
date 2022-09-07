# Kratos

## 步骤

* 安装go
* 安装protoc
* 安装protoc-gen-go：go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
* 安装kratos cli：go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
  > 查看帮助：-h
  > 查看版本：-v
  > 指定分支：-b
  > 指定源头layout：-r https://gitee.com/go-kratos/kratos-layout.git //国内拉取失败可使用gitee源
  > 创建项目：kratos new xxx
  > 共用go.mod,大仓模式：kratos new app/xxx  --nomod
  > 代码生成：go generate ./...
  > 代码运行：kratos run
  >

### 顺序：api->server->service->biz->data

> biz：文档里说了类似 DDD 的 domain 层，也就是 DDD 架构中的领域层。这里还定义了对业务操作的接口。业务逻辑组装。
> data：对数据库 db，缓存 cache 的封装，并且实现 biz 中定义的接口。它将领域对象重新拿出来，这里去掉了 DDD 的基础层。
> service：实现 api 定义的服务层，类似 DDD 的应用层。处理数据传输对象到 biz(领域实体)的转换。同时协同各类 biz 交互，不应处理复杂逻辑。
> server：http 和 grpc 实例的创建和配置，以及注册对应的 service。

* #### api(proto cli)

> 目录：api/xxx/v1
>
> 添加proto文件：kratos proto add api/xxx/v1/xxx.proto
> 生成proto代码：kratos proto client api/xxx/v1/xxx.proto

* #### server(实例化http和grpc)

> 目录：app/xxx/internal/server
> 修改http.go,grpc.go,server.go
> v1.RegisterXXXHTTPServer(srv, student)
> v1.RegisterXXXServer(srv, student)

* #### service(cli)

> 目录：app/xxx/internal/service
>
> 生成service：kratos proto server api/helloworld/demo.proto -t internal/service
>
> service.go依赖注入
>
> 修改xxx.go串通顺逻辑

* #### biz

> 目录：app/xxx/internal/biz
>
> 新建xxx.go
>
> 代码逻辑参考biz_template.md
>
> biz.go依赖注入

* #### data

> 目录：app/xxx/internal/data
>
> gorm:go get -u gorm.io/gorm
> go get -u gorm.io/driver/mysql
> Ent:go get entgo.io/ent/cmd/ent
>
> 新建xxx.go
>
> data.go 依赖注入+orm初始化
>
> 修改data配置文件:user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local

* #### config

> 如果修改了conf.proto
>
> 执行此命令更新conf.pb.go：make config

* #### wire 依赖注入

> 安装命令：go install github.com/google/wire/cmd/wire
> 导入项目：go get -u github.com/google/wire
> 进入到 cmd/xxx 目录，wire命令重新生成wire_gen.go文件

## 数据库
* ### Mysql
* docker安装：
> docker pull mysql
* docker启动：
> docker run --name mysql-local -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -v /root/data:/var/lib/mysql mysql

## 服务注册与发现

* ### Etcd
* docker安装：
> docker pull bitnami/etcd:latest
* 创建docker网络：
> docker network create app-tier --driver bridge
* 启动etcd服务器实例：–network app-tier 指定启动实例的网络配置 
> docker run -d 
> --name etcd-local-server 
> --network app-tier 
> --publish 2379:2379 
> --publish 2380:2380 
> --env ALLOW_NONE_AUTHENTICATION=yes 
> --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-local-server:2379 bitnami/etcd:latest
* 启动etcd客户端实例，并连接上步骤服务
>docker run -it --rm
--network app-tier
--env ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd:latest 
> etcdctl --endpoints http://etcd-local-server:2379 put /message Hello
* 检查
> curl -L http://127.0.0.1:2379/version

## 分布式链路追踪
* ### jaeger localhost:16686
* docker安装：
> docker run -d --name jaeger \
-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
-p 5775:5775/udp \
-p 6831:6831/udp \
-p 6832:6832/udp \
-p 5778:5778 \
-p 16686:16686 \
-p 14268:14268 \
-p 9411:9411 \
jaegertracing/all-in-one:latest