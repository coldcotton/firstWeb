package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func GetVotes() []Vote { // 获取所有投票
	ret := make([]Vote, 0) // 创建一个空的切片
	if err := Conn.Table("vote").Find(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return ret
}

func GetVote(id int64) VoteWithOpt { // 获取某个投票项目的详细信息
	var ret Vote
	if err := Conn.Table("vote").Where("id=?", id).First(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}

	opt := make([]VoteOpt, 0)
	if err := Conn.Table("vote_opt").Where("vote_id=?", id).Find(&opt).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	return VoteWithOpt{ // 返回值包含投票和选项
		Vote: ret,
		Opt:  opt,
	}
}

// GORM的事务方法
func DoVote(userId, voteId int64, optIDs []int64) bool { // 投票
	tx := Conn.Begin() // 开启事务
	var ret Vote
	// 是否有这个投票项目
	if err := tx.Table("vote").Where("id = ?", voteId).First(&ret).Error; err != nil {
		fmt.Printf("err:%s", err.Error())
		tx.Rollback() // 回滚
	}

	for _, value := range optIDs {
		// 票数+1
		if err := tx.Table("vote_opt").Where("id = ?", value).Update("count", gorm.Expr("count + ?", 1)).Error; err != nil {
			fmt.Printf("err:%s", err.Error())
			tx.Rollback()
		}
		user := VoteOptUser{
			VoteId:      voteId,
			UserId:      userId,
			VoteOptId:   value,
			CreatedTime: time.Now(),
		}
		if err := tx.Create(&user).Error; err != nil { // 创建用户投票记录，在vote_opt_user表
			fmt.Printf("err:%s", err.Error())
			tx.Rollback()
		}
	}

	tx.Commit() // 提交事务

	return true
}

// DoVoteV1 原生SQL的事务方法
func DoVoteV1(userId, voteId int64, optIDs []int64) bool {
	Conn.Exec("begin").
		Exec("select * from vote where id = ?", voteId).
		Exec("commit")
	return false
}

// DoVoteV2 匿名函数，用的最多的的事务方法
func DoVoteV2(userId, voteId int64, optIDs []int64) bool {
	if err := Conn.Transaction(func(tx *gorm.DB) error { // 开启一个事务
		var ret Vote
		if err := tx.Table("vote").Where("id = ?", voteId).First(&ret).Error; err != nil {
			fmt.Printf("err:%s", err.Error())
			return err
		}

		for _, value := range optIDs {
			if err := tx.Table("vote_opt").Where("id = ?", value).Update("count", gorm.Expr("count + ?", 1)).Error; err != nil {
				fmt.Printf("err:%s", err.Error())
				return err
			}
			user := VoteOptUser{
				VoteId:      voteId,
				UserId:      userId,
				VoteOptId:   value,
				CreatedTime: time.Now(),
			}
			if err := tx.Create(&user).Error; err != nil {
				return err
			}
		}
		return nil // 返回空就提交事务
	}); err != nil {
		fmt.Printf("err:%s", err.Error())
		return false
	}

	return true
}
