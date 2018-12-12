package main

import (
	"fmt"
	"learnGo/examples/go-micro-pay/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // msql
)

var DB *gorm.DB

type PayInfo struct {
	gorm.Model
	UID    int64  `gorm:"UNIQUE;NOT NULL" json:"uid"`
	Alipay string `gorm:"DEFAULT:''" json:"alipay"`
	TenPay string `gorm:"DEFAULT:''" json:"tenpay"`
	Intro  string `gorm:"DEFAULT:''" json:"intro"`
}

func init() {
	fmt.Println("Pay db init")
	var err error
	DB, err = gorm.Open("mysql", config.Config.MySQL)
	if err != nil {
		fmt.Println("connect Pay db eror", err.Error())
	}
	DB.LogMode(true)
	if !DB.HasTable(&PayInfo{}) {
		fmt.Println("pay db neet to start initing")
		err := DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&PayInfo{}).Error
		if err != nil {
			fmt.Println("init db err")
			panic(err)
		}
	}
	fmt.Println("db start success")

}
