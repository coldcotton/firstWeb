package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 所有数据库操作
var Conn *gorm.DB

func NewMysql() {
	// 用户名，密码，主机，数据库名，问号是分隔基本信息和连接参数
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "kq29k29w", "47.108.59.238:3306", "vote")
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s\n", err)
		panic(err)
	}

	Conn = conn
}

func Close() {
	db, _ := Conn.DB()
	_ = db.Close()
}