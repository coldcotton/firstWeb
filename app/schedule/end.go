package schedule

import (
	"fmt"
	"time"

	"github.com/coldcotton/firstWeb/app/model"
)

// 定时器
func Start() {
	go EndVote()
}

func EndVote() {
	t := time.NewTicker(10 * time.Second) // 通道，每隔一定时间传一个信号
	defer func() {
		t.Stop()
	}()

	for {
		select {
		case <-t.C:
			fmt.Println("EndVote 启动")
			// 函数
			model.EndVote()
			fmt.Println("EndVote 运行完毕")
		}
	}
}
