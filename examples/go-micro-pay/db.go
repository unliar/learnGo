package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // msql
)

var DB *gorm.DB

type PayInfo struct {
	gorm.Model
	UID    int64  `gorm:"UNIQUE;NOT NULL" json:"uid"`
	Alipay string `gorm:"DEFAULT:''" json:"alipay"`
	TenPay string `gorm:"DEFAULT:''" json:"ten_pay"`
}

func init() {
	fmt.Println("Pay db init")
	var err error
	DB, err = gorm.Open("mysql", "root:MySQL19930224@tcp(127.0.0.1:3306)/go-server?parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("connect Pay db eror", err.Error())
	}
	DB.LogMode(true)
	if !DB.HasTable(&PayInfo{}) {
		fmt.Println("pay db neet to start initing")
		err := DB.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&PayInfo{}).Error
		if err != nil {
			fmt.Println("init db err")
		}
	}
	fmt.Println("db start success")

}
