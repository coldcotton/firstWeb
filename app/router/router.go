package router

import (
	"fmt"
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
	// index.Use(checkUser) // 应用中间件到index路由组上
	{

		// vote，需要登录态
		index.GET("/index", logic.Index) // 静态页面

		index.POST("/vote/add", logic.AddVote)
		index.POST("/vote/update", logic.UpdateVote)
		index.POST("/vote/del", logic.DelVote)

		index.GET("/result", logic.ResultInfo)
		index.GET("/result/info", logic.ResultVote)
	}

	// RESTful风格接口
	{
		index.GET("/votes", logic.GetVotes)
		index.GET("/vote", logic.GetVoteInfo)

		index.POST("/vote", logic.AddVote)
		index.PUT("/vote", logic.UpdateVote)
		index.DELETE("/vote", logic.DelVote)

		index.GET("/vote/result", logic.ResultVote)

		index.POST("/do_vote", logic.DoVote)
	}

	{
		// login，不需要登录态
		g.GET("/login", logic.GetLogin)
		g.POST("/login", logic.DoLogin)
		g.GET("/logout", logic.Logout)

		// user
		g.POST("/user/create", logic.CreateUser)
	}

	//验证码
	{
		g.GET("/captcha", func(context *gin.Context) {
			captcha, err := tools.CaptchaGenerate()
			if err != nil {
				context.JSON(http.StatusOK, tools.ECode{
					Code:    10005,
					Message: err.Error(),
				})
				return
			}

			context.JSON(http.StatusOK, tools.ECode{
				Data: captcha,
			})
		})

		g.POST("/captcha/verify", func(context *gin.Context) {
			var param tools.CaptchaData
			if err := context.ShouldBind(&param); err != nil {
				context.JSON(http.StatusOK, tools.ParamErr)
				return
			}

			fmt.Printf("参数为：%+v", param)
			if !tools.CaptchaVerify(param) {
				context.JSON(http.StatusOK, tools.ECode{
					Code:    10008,
					Message: "验证失败",
				})
				return
			}
			context.JSON(http.StatusOK, tools.OK)
		})
	}

	if err := g.Run("0.0.0.0:8080"); err != nil {
		panic("gin 启动失败！！")
	}
}

// cookie
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

// session
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
		context.JSON(http.StatusUnauthorized, tools.NotLogin)
		// context.Abort()
	}

	context.Next()
}
