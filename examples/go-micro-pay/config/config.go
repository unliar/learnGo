package config

import (
	"errors"
	"fmt"
	"os"
	"time"

	config "github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

// EnvConfig 配置文件类型
type EnvConfig struct {
	MySQL            string
	ServiceName      string
	Env              string
	ServiceVersion   string
	RegisterInterval time.Duration
	RegisterTTL      time.Duration
}

// Config 配置实例
var Config EnvConfig

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(errors.New("get pwd error"))
	}
	ConfigPath := pwd + "/config/env.json"

	config.Load(
		file.NewSource(
			file.WithPath(ConfigPath)))
	config.Scan(&Config)
	fmt.Println("init config====>", Config, ConfigPath)
}
