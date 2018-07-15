package main

import (
	"gp-curl/router"
)

func main() {
	r := router.InitRoute()

	r.Run(":8088")
}
