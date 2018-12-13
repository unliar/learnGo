package service

import (
	"fmt"
	"learnGo/examples/go-micro-pay/config"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // msql
)

// DB 是当前数据的实例
var DB *gorm.DB

// UserInfo 用户信息model
type UserInfo struct {
	gorm.Model
	LoginName  string `gorm:"UNIQUE;NOT NULL;"`
	Nickname   string `gorm:"UNIQUE;NOT NULL"`
	Age        int64  `gorm:"DEFAULT:18"`
	Gender     int64  `gorm:"DEFAULT:0"`
	Avatar     string `gorm:"DEFAULT:''"`
	Location   string `gorm:"DEFAULT:'shenzhen'"`
	Profession string `gorm:"DEFAULT:''"`
	Status     int64  `gorm:"DEFAULT:1"`
	Phone      string `gorm:"DEFAULT:''"`
	Email      string `gorm:"DEFAULT:''"`
	WeChatId   string `gorm:"DEFAULT:''"`
	QQId       string `gorm:"DEFAULT:''"`
	Brief      string `gorm:"DEFAULT:''"`
	NationCode string `gorm:"DEFAULT:'86'"`
}

// UserPass 用户密码
type UserPass struct {
	gorm.Model
	UID      int64  `gorm:"UNIQUE;NOT NULL;"`
	Password string `gorm:"NOT NULL"`
}

func init() {
	fmt.Println("db start init")
	var err error
	DB, err = gorm.Open("mysql", config.Config.MySQL)
	DB.LogMode(true)
	if !DB.HasTable(&UserInfo{}) {
		fmt.Println("db UserBase need to create")
		err = DB.Set("gorm:table_options", "ENGINE=InnoDB").
			CreateTable(&UserInfo{}).Error
	}
	if !DB.HasTable(&UserPass{}) {
		fmt.Println("db UserContact need to create")
		err = DB.Set("gorm:table_options", "ENGINE=InnoDB").
			CreateTable(&UserPass{}).Error
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("db init success")
}
