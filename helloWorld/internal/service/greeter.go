package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "helloWorld/api/helloworld/v1"
	"helloWorld/internal/biz"
	"helloWorld/internal/dao"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper("service/greeter", logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.Infof("SayHello Received: %v", in.GetName())
	if in.GetName() == "error" {
		return nil, errors.NotFound("ErrorReason", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *GreeterService) GetUserList(ctx context.Context, in *v1.GetUserListRequest) (*v1.GetUserListReply, error) {
	s.log.Infof("GetUserList Received: %v", in.GetLevel())
	dbUserList, count, err := dao.D.QueryUserList(in.GetLevel())
	if err != nil {
		return nil, err
	}

	var list []*v1.User
	for _, dbUser := range *dbUserList {
		list = append(list, &v1.User{Name: dbUser.Name, Class: dbUser.Class})
	}

	return &v1.GetUserListReply{Count: count, List: list}, nil
}
