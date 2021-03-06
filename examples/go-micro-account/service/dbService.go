package service

import (
	"fmt"
	"learnGo/examples/go-micro-account/config"

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
	Brief      string `gorm:"DEFAULT:'你说的江湖,我曾去过,见过最美的风景,也遇过最难忘的人'"`
}

// UserPass 用户密码
type UserPass struct {
	gorm.Model
	UID      int64  `gorm:"UNIQUE;NOT NULL;"`
	Password string `gorm:"NOT NULL"`
}

// UserSecretInfo
type UserSecretInfo struct {
	gorm.Model
	UID         int64  `gorm:"UNIQUE;NOT NULL;"`
	Phone       string `gorm:"UNIQUE;DEFAULT:''"`
	PhoneCodeID int64  `gorm:"DEFAULT:'1'"`
	Email       string `gorm:"UNIQUE;DEFAULT:''"`
	WeChatId    string `gorm:"DEFAULT:''"`
	QQId        string `gorm:"DEFAULT:''"`
	RealName    string `gorm:"DEFAULT:''"`
}

// UserPhoneCode 用户手机号
type UserPhoneCode struct {
	gorm.Model
	NationCode string `gorm:"UNIQUE;DEFAULT:''"`
}

func init() {
	fmt.Println("db start init")
	var err error
	DB, err = gorm.Open("mysql", config.Config.MySQL)
	if err != nil {
		fmt.Println("db connecting  error")
		panic(err)
	}
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
	if !DB.HasTable(&UserSecretInfo{}) {
		fmt.Println("db UserSecretInfo need to create")
		err = DB.Set("gorm:table_options", "ENGINE=InnoDB").
			CreateTable(&UserSecretInfo{}).Error
	}
	if err != nil {
		fmt.Println("db table check error")
		panic(err)
	}
	fmt.Println("db init success")
}
