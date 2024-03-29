package model

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// 创建一个新的cookie存储会话实例，参数是密钥secretKey
var store = sessions.NewCookieStore([]byte("secretKey!!!!"))
var sessionName = "session-name"

// 获取
func GetSession(c *gin.Context) map[interface{}]interface{} {
	session, _ := store.Get(c.Request, sessionName) // 获取session
	fmt.Printf("session:%+v\n", session.Values)
	return session.Values
}

// 设置
func SetSession(c *gin.Context, name string, id int64) error {
	session, _ := store.Get(c.Request, sessionName)
	session.Values["name"] = name
	session.Values["id"] = id
	return session.Save(c.Request, c.Writer)
}

// 清空
func FlushSession(c *gin.Context) error {
	session, _ := store.Get(c.Request, sessionName)
	fmt.Printf("session : %+v\n", session.Values)
	session.Values["name"] = ""
	session.Values["id"] = int64(0)
	return session.Save(c.Request, c.Writer)
}
