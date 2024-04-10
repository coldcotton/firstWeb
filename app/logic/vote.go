package logic

import (
	"net/http"
	"strconv"
	"time"

	"github.com/coldcotton/firstWeb/app/model"
	"github.com/coldcotton/firstWeb/app/tools"
	"github.com/gin-gonic/gin"
)

func AddVote(context *gin.Context) { // 增加投票项目
	idStr := context.Query("title")
	optStr, _ := context.GetPostFormArray("opt_name[]")
	//构建结构体
	vote := model.Vote{
		Title:       idStr,
		Type:        0,
		Status:      0,
		CreatedTime: time.Now(),
	}

	opt := make([]model.VoteOpt, 0)
	for _, v := range optStr {
		opt = append(opt, model.VoteOpt{
			Name:        v,
			CreatedTime: time.Now(),
		})
	}

	if err := model.AddVote(vote, opt); err != nil {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, tools.OK)
}

func UpdateVote(context *gin.Context) {

}

func DelVote(context *gin.Context) { // 删除一个投票
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)

	// 实现删除操作的幂等性
	vote := model.GetVote(id)
	if vote.Vote.Id <= 0 {
		context.JSON(http.StatusNoContent, tools.OK)
		return
	}

	if err := model.DelVote(id); !err {
		context.JSON(http.StatusOK, tools.ECode{
			Code:    10006,
			Message: "删除失败！！",
		})
		return
	}

	context.JSON(http.StatusNoContent, tools.OK)
}

func ResultInfo(context *gin.Context) {
	context.HTML(http.StatusOK, "result.tmpl", nil)
}

type ResultData struct {
	Title string
	Count int64
	Opt   []*ResultVoteOpt
}

type ResultVoteOpt struct {
	Name  string
	Count int64
}

// 返回投票结果
func ResultVote(context *gin.Context) {
	var id int64
	idStr := context.Query("id")
	id, _ = strconv.ParseInt(idStr, 10, 64)
	ret := model.GetVote(id)
	data := ResultData{
		Title: ret.Vote.Title,
	}

	for _, v := range ret.Opt {
		data.Count = data.Count + v.Count
		tmp := ResultVoteOpt{
			Name:  v.Name,
			Count: v.Count,
		}
		data.Opt = append(data.Opt, &tmp)
	}

	context.JSON(http.StatusOK, tools.ECode{
		Data: data,
	})
}
