package logic

import (
	"net/http"

	"github.com/coldcotton/firstWeb/app/model"
	"github.com/gin-gonic/gin"
)

func GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.tmpl", nil)
}

func DoLogin(context *gin.Context) {
	var user model.User
	ret := make(map[string]any)
	_ = context.ShouldBind(&user)

	ret = model.GetUser(&user)

	context.JSON(http.StatusOK, ret)

}
