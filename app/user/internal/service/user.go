package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-crsdemo/api/user/v1"
	biz "kratos-crsdemo/app/user/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	user *biz.UserUsecase
	log  *log.Helper
}

// 初始化
func NewUserService(user *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		user: user,
		log:  log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.user.FindByID(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Name: user.Name,
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}
