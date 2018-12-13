package config

import (
	"fmt"
	"os"
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
	dir, _ := os.Getwd()

	config.Load(
		file.NewSource(
			file.WithPath(dir + "/config/env.json"),
		),
	)
	config.Scan(&Config)
	fmt.Println("InitConfig start====>", Config)
}
