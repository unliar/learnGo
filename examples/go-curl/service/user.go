package service

import (
	db2 "gp-curl/db"
	"log"
)

type User struct {
	Id   int    `json:"id" `
	Name string `json:"name"`
	Age  int    `json:"age"`
	Male string `json:"male"`
}

var db = db2.MySQL

//根据id获取指定的用户信息
func (u *User) GetUserById() (*User, error) {
	row := db.QueryRow("select * from user_message where id= ?", u.Id)
	err := row.Scan(&u.Id, &u.Name, &u.Age, &u.Male)
	if err != nil {
		return nil, err
	}

	return u, nil
}

//获取所有同龄人
func (u *User) GetUsersByAge(PageIndex int, pageCount int) ([]*User, error) {
	sm, err := db.Prepare("select * from user_message where age= ? order by id desc limit ? offset ?")
	if err != nil {
		return nil, err
	}
	defer sm.Close()

	offset := (PageIndex - 1) * pageCount
	rows, err := sm.Query(u.Age, pageCount, offset)

	if err != nil {
		return nil, err
	}

	list := make([]*User, 0)
	defer rows.Close()
	for rows.Next() {

		item := new(User)

		err := rows.Scan(&item.Id, &item.Name, &item.Age, &item.Male)

		if err != nil {
			log.Fatal(err.Error())
			return nil, err
		}

		list = append(list, item)

	}

	return list, nil
}
