package model

import (
	"fmt"
	"testing"
	"time"
)

func TestGetVotes(t *testing.T) {
	NewMysql()
	r := GetVotes()
	fmt.Printf("ret:%+v", r) // %+v以详细的方式显示该值的内容，包括结构体字段的名称
	Close()
}

func TestGetVote(t *testing.T) {
	NewMysql()
	r := GetVote(1) // vote_id
	fmt.Printf("ret:%+v", r)
	Close()
}

func TestDoVote(t *testing.T) {
	NewMysql()
	r := DoVote(1, 1, []int64{1, 2}) // vote_id
	fmt.Printf("ret:%+v", r)
	Close()
}

func TestAddVote(t *testing.T) {
	NewMysql()

	vote := Vote{
		Title:       "测试用例",
		Type:        0,
		Status:      0,
		Time:        0,
		UserId:      0,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	opt := make([]VoteOpt, 0)
	opt = append(opt, VoteOpt{
		Name:        "测试选项1",
		Count:       0,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})
	opt = append(opt, VoteOpt{
		Name:        "测试选项2",
		Count:       0,
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	})

	r := AddVote(vote, opt)
	fmt.Printf("ret:%+v", r)
	Close()
}

func TestGetUserV1(t *testing.T) {
	NewMysql()
	a := GetUserV1("admin")
	fmt.Printf("a:%+v", a)
}
