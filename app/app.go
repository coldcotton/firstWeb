package app

import (
	"github.com/coldcotton/firstWeb/app/model"
	"github.com/coldcotton/firstWeb/app/router"
)

// 启动器方法
func Start() {
	model.NewMysql() // 创建数据库
	defer func() {   // 最后执行
		model.Close()
	}()

	router.NewRouter()
}
