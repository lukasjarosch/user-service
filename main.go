package main

import (
	"fmt"
	proto "github.com/lukasjarosch/user-service/proto/auth"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&proto.User{})

	repo := &UserRepository{db}

	srv := micro.NewService(
		micro.Name("go.micro.srv.auth"),
		micro.Version("latest"),
	)
	srv.Init()

	proto.RegisterAuthServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
