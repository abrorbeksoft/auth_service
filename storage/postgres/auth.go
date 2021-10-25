package postgres

import (
	"database/sql"
	"fmt"
	auth "github.com/abrorbeksoft/auth_service/genproto/auth_service"
	"github.com/abrorbeksoft/auth_service/storage/repos"
)

type userStorage struct {
	session *sql.DB
}

func NewAuth(session *sql.DB) repos.UserStorageI  {
	return &userStorage{
		session: session,
	}
}

func (storage *userStorage) Create(user *auth.RegisterRequest) string {
	var id string
	fmt.Println(user)
	fmt.Println(user.Password)
	err:=storage.session.QueryRow("INSERT INTO users (id,name,surname,login,password,created_at,updated_at) VALUES (gen_random_uuid(),$1,$2,$3,$4,now(),now()) RETURNING id",user.Name,user.Surname,user.Login,user.Password).Scan(&id)
	if err != nil {
		fmt.Println("Hello world")
		fmt.Println(err.Error())
	}
	return id
}

func (storage *userStorage) GetById(id string) *auth.FindOneResponse {
	var user *auth.FindOneResponse;
	err:=storage.session.QueryRow("SELECT id,name,surname,login,created_at,updated_at FROM users where id=$1",id).Scan(&user.Id,&user.Name,&user.Surname,&user.Login,&user.CreatedAt,&user.UpdatedAt)
	if err!= nil {
		fmt.Println(err.Error())
	}
	return user
}