package model

import (
	"fmt"
	"testing"
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
