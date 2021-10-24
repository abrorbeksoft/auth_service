package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"os"
)
type config struct {
	AuthServiceHost string
	AuthServicePort string
}

func Load() *config {
	if err:=godotenv.Load(); err!= nil {
		fmt.Println(err.Error())
	}
	conf:=config{}
	conf.AuthServiceHost = cast.ToString(getOrReturnDefaultValue("AUTH_HOST","localhost"))
	conf.AuthServicePort = cast.ToString(getOrReturnDefaultValue("AUTH_PORT","8000"))
	return &conf
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists:=os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
