package app

import (
	"github.com/coldcotton/firstWeb/app/model"
	"github.com/coldcotton/firstWeb/app/router"
	"github.com/coldcotton/firstWeb/app/tools"
)

// 启动器方法
func Start() {
	model.NewMysql() // 创建数据库
	defer func() {   // 最后执行
		model.Close()
	}()

	// schedule.Start() // 启动定时器

	tools.NewLogger() // 日志

	router.NewRouter() // 启动路由
}
