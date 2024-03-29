package model

import (
	"fmt"
)

func GetUser(name string) *User {
	var ret User
	// 查找用户名为name的用户，结果返回到ret中
	if err := Conn.Table("user").Where("name=?", name).Find(&ret).Error; err != nil {
		fmt.Println("err:", err.Error())
	}
	return &ret
}
