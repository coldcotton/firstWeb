package model

import "fmt"

type User struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func GetUser(user *User) map[string]any {
	ret := make(map[string]any)
	// 查找用户名为user.name的用户
	if err := Conn.Table("user").Where("name=?", user.Name).Find(&ret).Error; err != nil {
		fmt.Println("err:", err.Error())
	}
	return ret
}
