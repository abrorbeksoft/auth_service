package main

import (
	"fmt"
	"github.com/abrorbeksoft/auth_service/config"
	auth "github.com/abrorbeksoft/auth_service/genproto/auth_service"
	"github.com/abrorbeksoft/auth_service/services"
	"google.golang.org/grpc"
	"net"
)

func main()  {
	cnf:=config.Load();

	lis,err:=net.Listen("tcp", fmt.Sprintf("%s:%s",cnf.AuthServiceHost,cnf.AuthServicePort))
	if err != nil {
		fmt.Println(err.Error())
	}

	s:=grpc.NewServer()
	authService:=services.NewAuthService();
	auth.RegisterAuthServiceServer(s,authService)

	if err:=s.Serve(lis); err!=nil {
		fmt.Println(err.Error())
	}
}