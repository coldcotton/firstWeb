package router

import (
	"net/http"

	"github.com/coldcotton/firstWeb/app/logic"
	"github.com/coldcotton/firstWeb/app/model"
	"github.com/coldcotton/firstWeb/app/tools"
	"github.com/gin-gonic/gin"
)

func NewRouter() {
	g := gin.Default()           // 创建默认路由
	g.LoadHTMLGlob("app/view/*") // 加载html文件

	index := g.Group("") // 创建路由组
	index.Use(checkUser) // 应用中间件到index路由组上
	index.GET("/index", logic.Index)

	// 没使用cookie的
	// g.GET("/index", logic.Index)
	g.GET("/vote", logic.GetVoteInfo)
	g.POST("/vote", logic.DoVote)

	// 相关路径
	g.GET("/login", logic.GetLogin)
	g.POST("/login", logic.DoLogin)
	g.GET("/logout", logic.Logout)

	if err := g.Run("0.0.0.0:8080"); err != nil {
		panic("gin 启动失败！！")
	}
}

// func checkUser(context *gin.Context) {
// 	name, err := context.Cookie("name") // 获取get请求中，名为name的cookie的值
// 	if err != nil || name == "" {
// 		// context.JSON(http.StatusOK, map[string]string{
// 		// 	"message": "未登录",
// 		// })
// 		context.Redirect(http.StatusFound, "/login") // 未登录则重定向到login页面
// 		// context.Abort()                           // 不继续执行，直接退出
// 	}
// 	context.Next() // 将控制权传递给下一个中间件或者路由处理函数
// }

func checkUser(context *gin.Context) {
	var name string
	var id int64
	values := model.GetSession(context)

	if v, ok := values["name"]; ok {
		name = v.(string)
	}

	if v, ok := values["id"]; ok {
		id = v.(int64)
	}

	if name == "" || id <= 0 {
		context.JSON(http.StatusOK, tools.NotLogin)
		context.Abort()
	}

	context.Next()
}
