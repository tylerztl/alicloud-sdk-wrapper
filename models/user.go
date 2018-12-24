package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Username string
	Password string
}

func AddUser(user User) string {
	mapper := orm.NewOrm()
	id, err := mapper.Insert(&user)
	if err != nil {
		return err.Error()
	} else {
		return string(id)
	}
}

func GetAllUsers() map[string]*User {
	var maps []orm.Params
	mapper := orm.NewOrm()
	num, err := mapper.Raw("SELECT * FROM user").Values(&maps)
	if err == nil {
		fmt.Println(num)
		for _, user := range maps {
			fmt.Println(user["id"], ":", user["name"])
		}
	}
	return nil
}

func GetUser(id int) (user *User, err error) {
	mapper := orm.NewOrm()
	u := &User{Id: id}
	err = mapper.Read(u)
	return u, err
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	return nil, nil
}

func Login(username, password string) bool {
	return false
}

func DeleteUser(uid string) {

}
