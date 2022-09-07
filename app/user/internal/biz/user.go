package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-crsdemo/api/user/v1"
)

// interface
type UserRepo interface {
	FindByID(context.Context, int64) (*pb.GetUserReply, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper //日志
}

// 初始化
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) FindByID(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	uc.log.WithContext(ctx).Infof("CreateStudent: %v", req)
	return uc.repo.FindByID(ctx, req.Id)
}
