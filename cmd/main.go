package main

import (
	"database/sql"
	"fmt"
	"github.com/abrorbeksoft/auth_service/config"
	auth "github.com/abrorbeksoft/auth_service/genproto/auth_service"
	"github.com/abrorbeksoft/auth_service/services"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
)

func main()  {
	cnf:=config.Load();
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", cnf.PostgresUser,cnf.PostgresPassword,cnf.PostgresDBName)

	session, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	lis,err:=net.Listen("tcp", fmt.Sprintf("%s:%s",cnf.AuthServiceHost,cnf.AuthServicePort))
	if err != nil {
		fmt.Println(err.Error())
	}

	s:=grpc.NewServer()
	authService:=services.NewAuthService(session);
	auth.RegisterAuthServiceServer(s,authService)

	if err:=s.Serve(lis); err!=nil {
		fmt.Println(err.Error())
	}
}