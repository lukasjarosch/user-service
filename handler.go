package main

import (
	proto "github.com/lukasjarosch/user-service/proto/auth"
	"golang.org/x/net/context"
)

type service struct {
	repo Repository
}

func (srv *service) Get(ctx context.Context, req *proto.User, res *proto.Response) error {
	user, err := srv.repo.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (srv *service) GetAll(ctx context.Context, req *proto.Request, res *proto.Response) error {
	users, err := srv.repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (srv *service) Auth(ctx context.Context, req *proto.User, res *proto.Token) error {
	_, err := srv.repo.GetByEmailAndPassword(req)
	if err != nil {
		return err
	}
	res.Token = "testingAbc"
	return nil
}

func (srv *service) Create(ctx context.Context, req *proto.User, res *proto.Response) error {
	if err := srv.repo.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (srv *service) ValidateToken(ctx context.Context, req *proto.Token, res *proto.Token) error {
	return nil
}
