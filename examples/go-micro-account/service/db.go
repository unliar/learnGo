package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type UserBase struct {
	gorm.Model
	LoginName  string `gorm:"unique";NOT NULL;"`
	IDC        string `gorm:"DEFAULT:'';"`
	Nickname   string `gorm:"UNIQUE";NOT NULL"`
	Age        int64  `gorm:"DEFAULT:0"`
	Male       string `gorm:"size:6"`
	Avatar     string `gorm:"size:255"`
	Location   string `gorm:"DEFAULT:'shenzhen'"`
	Profession string `gorm:"DEFAULT:''"`
	Status     int64  `gorm:"DEFAULT:1"`
}

type UserContact struct {
	gorm.Model
	UID      int64  `gorm:"unique";`
	Email    string `gorm:"DEFAULT:''"`
	Phone    string `gorm:"DEFAULT:''"`
	WeChatId string `gorm:"DEFAULT:''"`
	WeiBoId  string `gorm:"DEFAULT:''"`
	QQId     string `gorm:"DEFAULT:''"`
}

func init() {
	fmt.Println("db start init")
	var err error
	DB, err = gorm.Open("mysql", "root:MySQL19930224@tcp(127.0.0.1:3306)/go-server?parseTime=true&loc=Local")
	DB.LogMode(true)
	if !DB.HasTable(&UserBase{}) {
		fmt.Println("db UserBase need to create")
		err = DB.Set("gorm:table_options", "ENGINE=InnoDB").
			CreateTable(&UserBase{}).Error
	}
	if !DB.HasTable(&UserContact{}) {
		fmt.Println("db UserContact need to create")
		err = DB.Set("gorm:table_options", "ENGINE=InnoDB").
			CreateTable(&UserContact{}).Error
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("db init success")
}
