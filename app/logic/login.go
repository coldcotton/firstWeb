package logic

import (
	"net/http"

	"github.com/coldcotton/firstWeb/app/model"
	"github.com/coldcotton/firstWeb/app/tools"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name" form:"name"` // form的值必须和前端一样
	Password string `json:"password" form:"password"`
}

func GetLogin(context *gin.Context) { // 获得登录页面
	context.HTML(http.StatusOK, "login.tmpl", nil)
}

func DoLogin(context *gin.Context) { // 用户登录
	var user User

	if err := context.ShouldBind(&user); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Message: err.Error(),
		})
	}

	ret := model.GetUser(user.Name)
	if ret.Id < 1 || ret.Password != user.Password {
		context.JSON(http.StatusOK, tools.UserErr)

		return
	}

	// 设置cookie
	// context.SetCookie("name", user.Name, 3600, "/", "", false, false) // 第6个是安全性
	// context.SetCookie("id", fmt.Sprint(ret.Id), 3600, "/", "", false, false)

	_ = model.SetSession(context, user.Name, ret.Id) // 用户名、用户id

	context.JSON(http.StatusOK, tools.ECode{
		Message: "登录成功",
	})

	return
}

func Logout(context *gin.Context) { // 用户退出登录
	// context.SetCookie("name", "", 3600, "/", "", false, false)
	// context.SetCookie("id", "", 3600, "/", "", false, false)

	_ = model.FlushSession(context)
	context.Redirect(http.StatusFound, "/login")

}
