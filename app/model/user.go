package model

import (
	"fmt"

	"github.com/coldcotton/firstWeb/app/tools"
)

func GetUser(name string) *User {
	var ret User
	// 查找用户名为name的用户，结果返回到ret中
	if err := Conn.Table("user").Where("name=?", name).Find(&ret).Error; err != nil {
		fmt.Println("err:", err.Error())
	}
	return &ret
}

// 原生sql
func GetUserV1(name string) *User {
	var ret User
	err := Conn.Raw("select * from user where name = ? limit 1", name).Scan(&ret).Error
	if err != nil {
		tools.Logger.Errorf("[GetUserV1]err:%s", err.Error())
	}
	return &ret
}

func CreateUser(user *User) error {
	if err := Conn.Create(user).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		return err
	}
	return nil
}
