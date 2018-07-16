package conf

import (
	"flag"
	"fmt"
)

type Conf struct {
	MySQL   string `json:"my_sql"`
	Redis   string `json:"redis"`
	Mongodb string `json:"mongodb"`
	Port    string `json:"port"`
	Mode    string `json:"mode"`
}

var Dev = Conf{
	"root:MySQL19930224@tcp(127.0.0.1:3306)/go-server",
	"127.0.0.1:3306",
	"127.0.0.1:27017",
	":8088",
	"debug",
}

var Pro = Conf{
	"root:MySQL19930224@tcp(127.0.0.1:3306)/go-server",
	"127.0.0.1:3306",
	"127.0.0.1:27017",
	":8088",
	"release",
}

var CurrentConf Conf

func init() {

	e := flag.String("env", "Dev", "当前的运行环境")
	flag.Parse()
	fmt.Println("当前运行环境", *e)
	switch *e {
	case "Dev":
		CurrentConf = Dev

	case "Pro":
		CurrentConf = Pro
	default:
		panic("env error ---no env")

	}

}
