package logic

import (
	"net/http"
	"strconv"

	"github.com/coldcotton/firstWeb/app/model"
	"github.com/coldcotton/firstWeb/app/tools"
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	ret := model.GetVotes()
	// 第三个参数是一个`gin.H`类型的map，用于向模板传递数据。这里，传递了一个名为`vote`的键，其值为ret
	context.HTML(http.StatusOK, "index.tmpl", gin.H{"vote": ret})
}

func GetVoteInfo(context *gin.Context) {
	var id int64
	idStr := context.Query("id")            // 从前端请求的查询参数中获取名为id的参数值
	id, _ = strconv.ParseInt(idStr, 10, 64) // str转int，十进制64位
	ret := model.GetVote(id)
	context.HTML(http.StatusOK, "vote.tmpl", gin.H{"vote": ret})

}

func DoVote(context *gin.Context) { //  投票
	userIdstr, _ := context.Cookie("id")

	voteIdstr, _ := context.GetPostForm("vote_id") // 从POST请求的表单数据中获取名为"vote_id"的值
	optstr, _ := context.GetPostFormArray("opt[]")

	userId, _ := strconv.ParseInt(userIdstr, 10, 64) // 用户id
	voteId, _ := strconv.ParseInt(voteIdstr, 10, 64) // 投票id

	opt := make([]int64, 0) // 选项id
	for _, v := range optstr {
		optId, _ := strconv.ParseInt(v, 10, 64)
		opt = append(opt, optId)
	}

	model.DoVote(userId, voteId, opt)
	context.JSON(http.StatusOK, tools.Ecode{
		Message: "投票成功！！！",
	})

}
