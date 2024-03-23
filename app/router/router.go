package router

import (
	"github.com/coldcotton/firstWeb/app/logic"
	"github.com/gin-gonic/gin"
)

func NewRouter() {
	g := gin.Default()           // 创建默认路由
	g.LoadHTMLGlob("app/view/*") // 加载html文件

	g.GET("/", logic.Index)

	// 相关路径
	g.GET("/login", logic.GetLogin)
	g.POST("/login", logic.DoLogin)

	if err := g.Run("0.0.0.0:8080"); err != nil {
		panic("gin 启动失败！！")
	}
}
