# biz template
## 1.定义struct Demo:
```go
type Demo struct {
ID      string
Name    string
}
```
## 2.定义对Demo的操作接口：
```go
type DemoRepo interface {
Save(context.Context, *Demo) (*Demo, error)
}
```
## 3.对 Demo 的操作加上日志：
```go
type DemoUsercase struct {
repo DemoRepo
data *Data   // 这里 *Data 是连接数据库客户端
log  *log.Helper
}
```
## 4.初始化 DemoUsercase
```go
func NewDemoUsercase(data *Data,repo DemoRepo, logger log.Logger) *DemoUsercase {
return &DemoUsercase{repo: repo, log: log.NewHelper(logger)}
}
```
## 5.编写 CreateDemo 方法，也就是一些业务逻辑编写
```go
func (uc *DemoUsercase) CreateStudent(ctx contxt.Context, stu *Demo) (*Demo, error) {
uc.log.WithContext(ctx).Infof("CreateStudent: %v", stu.ID)
return uc.repo.Save(ctx, stu)
}
```
6.向 wire 注入Demo
internal/biz/biz.go:
```go
var ProviderSet = wire.NewSet(NewGreeterUsecase, NewDemoUsercase)
```