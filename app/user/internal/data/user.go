package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "kratos-crsdemo/api/user/v1"
	"kratos-crsdemo/app/user/internal/biz"
	"kratos-crsdemo/app/user/internal/entity"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
func (r *userRepo) FindByID(ctx context.Context, id int64) (*pb.GetUserReply, error) {
	var user entity.User
	r.data.gormDB.Where("id = ?", id).First(&user)
	r.log.WithContext(ctx).Info("gormDB: FindByID, id: ", id)
	return &pb.GetUserReply{
		Name: user.Name,
	}, nil
}
