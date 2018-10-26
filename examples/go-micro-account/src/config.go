package main

import (
	"fmt"
	"time"

	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

type ConfigModel struct {
	MySQL            string
	JWTTokenKey      string
	MD5Key           string
	Env              string
	ServiceName      string
	ServiceVersion   string
	RegisterInterval time.Duration
	RegisterTTL      time.Duration
}

var Config ConfigModel

func InitConfig() {

	config.Load(
		file.NewSource(
			file.WithPath("./env.json"),
		),
	)
	config.Scan(&Config)
	fmt.Println("InitConfig start====>", Config)
}
